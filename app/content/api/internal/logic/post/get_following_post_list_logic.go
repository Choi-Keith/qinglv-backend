package post

import (
	"context"
	"encoding/json"
	"strings"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/user/rpc/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type GetFollowingPostListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFollowingPostListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowingPostListLogic {
	return &GetFollowingPostListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFollowingPostListLogic) GetFollowingPostList(req *types.GetPostListReq) (resp *types.GetPostListResp, err error) {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	postFeedListResp, err := l.svcCtx.PostRpc.GetPostFeedList(l.ctx, &content.GetPostFeedListReq{
		UserId: uint64(userId),
	})
	if err != nil {
		return nil, err
	}
	postList, err := mr.MapReduce(func(source chan<- content.PostFeedItem) {
		for _, postItem := range postFeedListResp.Data {
			source <- *postItem
		}
	}, func(item content.PostFeedItem, writer mr.Writer[types.PostItem], cancel func(error)) {
		postResp, err := l.svcCtx.PostRpc.GetPostDetail(l.ctx, &content.GetPostDetailReq{
			Id: item.PostId,
		})
		if err != nil {
			cancel(err)
			return
		}
		postContentResp, err := l.svcCtx.PostRpc.GetPostContentByPostId(l.ctx, &content.GetPostContentDetailReq{
			Id: item.PostId,
		})
		if err != nil {
			cancel(err)
			return
		}
		userResp, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.GetUserByIdReq{
			UserId: item.UserId,
		})
		if err != nil {
			cancel(err)
			return
		}
		var postItem types.PostItem
		_ = copier.Copy(&postItem, &postResp)
		_ = copier.Copy(&postItem.Creator, userResp.User)
		postItem.Creator.IsFollowing = true
		var categoryItem types.PostCategory
		if postContentResp.PostContent.CategoryId != 0 {
			categoryResp, err := l.svcCtx.CategoryRpc.GetCategoryDetail(l.ctx, &content.GetCategoryDetailReq{
				Id: postContentResp.PostContent.CategoryId,
			})
			if err != nil {
				cancel(err)
				return
			}
			_ = copier.Copy(&categoryItem, categoryResp.Category)
		}
		postItem.Category = categoryItem
		topics := strings.Split(postContentResp.PostContent.Topics, ",")
		postItem.Topic = topics
		postItem.Content = postContentResp.PostContent.Content
		postItem.Images = postContentResp.PostContent.Images
		writer.Write(postItem)
	}, func(pipe <-chan types.PostItem, writer mr.Writer[[]types.PostItem], cancel func(error)) {
		var r []types.PostItem
		m := make(map[uint64]types.PostItem, len(postFeedListResp.Data))
		for p := range pipe {
			m[p.Id] = p
		}
		for _, postItem := range postFeedListResp.Data {
			r = append(r, m[postItem.Id])
		}
		writer.Write(r)
	})
	if err != nil {
		return nil, err
	}
	isEnd := false
	total := (req.PageNum-1)*req.PageSize + req.PageSize
	if postFeedListResp.Total <= uint64(total) {
		isEnd = true
	}

	return &types.GetPostListResp{
		List:  postList,
		Total: postFeedListResp.Total,
		IsEnd: isEnd,
	}, nil
}
