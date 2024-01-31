package post

import (
	"context"
	"net/http"
	"strings"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/user/rpc/user_client"
	"qinglv-backend/pkg/jwtx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetPostByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetPostByIdLogic {
	return &GetPostByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *GetPostByIdLogic) GetPostById(req *types.GetPostByIdReq, r *http.Request) (resp *types.GetPostByIdResp, err error) {
	// todo: add your logic here and delete this line
	postResp, err := l.svcCtx.PostRpc.GetPostDetail(l.ctx, &content.GetPostDetailReq{
		Id: req.Id,
	})
	if err != nil {
		logx.Errorf("[Post] GetPostDetail failed: %+v\n", err)
		return nil, err
	}
	postContentResp, err := l.svcCtx.PostRpc.GetPostContentByPostId(l.ctx, &content.GetPostContentDetailReq{
		Id: postResp.Post.Id,
	})
	if err != nil {
		logx.Errorf("[Post] GetPostContentByPostId failed: %+v\n", err)
		return nil, err
	}
	userResp, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user_client.GetUserByIdReq{
		UserId: postResp.Post.CreatorId,
	})
	if err != nil {
		logx.Errorf("[Post] GetUserById failed: %+v\n", err)
		return nil, err
	}
	var postItem types.PostItem
	_ = copier.Copy(&postItem, postResp.Post)
	_ = copier.Copy(&postItem.Creator, userResp.User)
	m, err := jwtx.ParseToken(r, l.svcCtx.Config.JWTAuth.AccessSecret)
	if err == nil {
		userId, ok := m["userId"]
		if ok && userId != 0 {
			if err != nil {
				logx.Errorf("[Post] get userId failed: %+v\n", err)
				return nil, err
			}
			followingResp, err := l.svcCtx.UserRpc.GetFollowingList(l.ctx, &user_client.GetFollowingListReq{
				UserId:      uint64(userId),
				FollowingId: postItem.Creator.Id,
			})
			if err != nil {
				logx.Errorf("[Post] GetFollowingList failed: %+v\n", err)
				return nil, err
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
			logx.Errorf("[Post] GetCategoryDetail failed: %+v\n", err)
			return nil, err
		}
		_ = copier.Copy(&categoryItem, categoryResp.Category)
	}
	postItem.Category = categoryItem
	topics := strings.Split(postContentResp.PostContent.Topics, ",")
	postItem.Topic = topics
	postItem.Content = postContentResp.PostContent.Content
	postItem.Images = postContentResp.PostContent.Images

	return &types.GetPostByIdResp{
		Post: postItem,
	}, nil
}
