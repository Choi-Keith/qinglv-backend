package post

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/user/rpc/user_client"
	"qinglv-backend/common/globalKey"
	"qinglv-backend/pkg/jwtx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type GetPostListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetPostListLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetPostListLogic {
	return &GetPostListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *GetPostListLogic) GetPostList(req *types.GetPostListReq, r *http.Request) (resp *types.GetPostListResp, err error) {
	// todo: add your logic here and delete this line
	var creatorId uint64
	if req.Creator != "" {
		userResp, err := l.svcCtx.UserRpc.CheckNicknameExist(l.ctx, &user_client.CheckNicknameExistReq{
			Nickname: req.Creator,
		})
		if err != nil {
			return nil, err
		}
		if userResp.IsExist {
			creatorId = userResp.User.Id
		}
	}
	args := []int32{globalKey.PostVisitPublic}
	postListResp, err := l.svcCtx.PostRpc.GetPostList(l.ctx, &content.GetPostListReq{
		Status:     int32(req.Status),
		Visibility: args,
		Score:      req.Score,
		Sort:       req.Sort,
		CreatorId:  creatorId,
		PageNum:    uint64(req.PageNum),
		PageSize:   uint64(req.PageSize),
	})
	if err != nil {
		return nil, err
	}
	postList, err := mr.MapReduce(func(source chan<- content.PostItem) {
		for _, postItem := range postListResp.Data {
			source <- *postItem
		}
	}, func(item content.PostItem, writer mr.Writer[types.PostItem], cancel func(error)) {
		postContentResp, err := l.svcCtx.PostRpc.GetPostContentByPostId(l.ctx, &content.GetPostContentDetailReq{
			Id: item.Id,
		})
		if err != nil {
			cancel(err)
			return
		}
		userResp, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user_client.GetUserByIdReq{
			UserId: item.CreatorId,
		})
		if err != nil {
			cancel(err)
			return
		}
		var postItem types.PostItem
		_ = copier.Copy(&postItem, &item)
		_ = copier.Copy(&postItem.Creator, userResp.User)
		m, err := jwtx.ParseToken(r, l.svcCtx.Config.JWTAuth.AccessSecret)
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
					FollowingId: postItem.Creator.Id,
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
		m := make(map[uint64]types.PostItem, len(postListResp.Data))
		for p := range pipe {
			m[p.Id] = p
		}
		for _, postItem := range postListResp.Data {
			r = append(r, m[postItem.Id])
		}
		writer.Write(r)
	})
	if err != nil {
		return nil, err
	}
	isEnd := false
	total := (req.PageNum-1)*req.PageSize + req.PageSize
	if postListResp.Total <= uint64(total) {
		isEnd = true
	}

	return &types.GetPostListResp{
		List:  postList,
		Total: postListResp.Total,
		IsEnd: isEnd,
	}, nil
}
