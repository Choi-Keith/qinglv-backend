package collection

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"
	"qinglv-backend/app/operation/rpc/operation"
	"qinglv-backend/app/user/rpc/user_client"
	"qinglv-backend/pkg/jwtx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type GetCollectionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetCollectionListLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetCollectionListLogic {
	return &GetCollectionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *GetCollectionListLogic) GetCollectionList(req *types.GetCollectionListReq) (resp *types.GetCollectionListResp, err error) {
	// todo: add your logic here and delete this line
	collectionListResp, err := l.svcCtx.CollectionRpc.GetCollectionList(l.ctx, &operation.GetCollectionListReq{
		GroupId:  req.GroupId,
		PageNum:  uint64(req.PageNum),
		PageSize: uint64(req.PageNum),
	})
	if err != nil {
		return nil, err
	}
	isEnd := false
	total := (req.PageNum-1)*req.PageSize + req.PageSize
	if collectionListResp.Total <= uint64(total) {
		isEnd = true
	}
	if req.BizType == 1 {
		postList, err := l.HanldePost(collectionListResp)
		if err != nil {
			return nil, err
		}
		return &types.GetCollectionListResp{
			Post: types.PostResp{
				List:  postList,
				Total: collectionListResp.Total,
				IsEnd: isEnd,
			},
			Article: types.ArticleResp{},
		}, nil
	}
	if req.BizType == 2 {
		// TODO: 文章收藏
	}

	return &types.GetCollectionListResp{
		Post:    types.PostResp{},
		Article: types.ArticleResp{},
	}, nil
}

func (l *GetCollectionListLogic) HanldePost(collectionListResp *operation.GetCollectionListResp) ([]types.PostItem, error) {
	postList, err := mr.MapReduce(func(source chan<- operation.CollectionItem) {
		for _, collectionItem := range collectionListResp.Data {
			source <- *collectionItem
		}
	}, func(item operation.CollectionItem, writer mr.Writer[types.PostItem], cancel func(error)) {
		postResp, err := l.svcCtx.PostRpc.GetPostDetail(l.ctx, &content.GetPostDetailReq{
			Id: item.TargetId,
		})
		if err != nil {
			cancel(err)
			return
		}
		postContentResp, err := l.svcCtx.PostRpc.GetPostContentByPostId(l.ctx, &content.GetPostContentDetailReq{
			Id: postResp.Post.Id,
		})
		if err != nil {
			cancel(err)
			return
		}
		userResp, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user_client.GetUserByIdReq{
			UserId: postResp.Post.CreatorId,
		})
		if err != nil {
			cancel(err)
			return
		}
		var postItem types.PostItem
		_ = copier.Copy(&postItem, postResp.Post)
		_ = copier.Copy(&postItem.Creator, userResp.User)

		m, err := jwtx.ParseToken(l.r, l.svcCtx.Config.JWTAuth.AccessSecret)
		if err == nil {
			fmt.Printf("m: %+v\n", m)
			userId, ok := m["userId"]
			if ok && userId != 0 {
				if err != nil {
					cancel(err)
					return
				}
				followingResp, err := l.svcCtx.UserRpc.GetFollowingList(l.ctx, &user_client.GetFollowingListReq{
					UserId:      uint64(userId),
					FollowingId: userResp.User.Id,
				})
				if err != nil {
					cancel(err)
					return
				}
				if len(followingResp.Data) != 0 {
					postItem.Creator.IsFollowing = true
				}
			}
		}
		var categoryItem types.CollectionCategory
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
		m := make(map[uint64]types.PostItem, len(collectionListResp.Data))
		for p := range pipe {
			m[p.Id] = p
		}
		for _, postItem := range collectionListResp.Data {
			r = append(r, m[postItem.TargetId])
		}
		writer.Write(r)
	})
	if err != nil {
		return nil, err
	}
	return postList, nil
}
