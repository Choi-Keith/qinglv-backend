package post

import (
	"context"
	"net/http"
	"strings"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/content_client"
	"qinglv-backend/common/globalKey"
	"qinglv-backend/pkg/jwtx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type GetUserPostListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetUserPostListLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetUserPostListLogic {
	return &GetUserPostListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *GetUserPostListLogic) GetUserPostList(req *types.GetUserPostListReq, r *http.Request) (resp *types.GetUserPostListResp, err error) {
	// todo: add your logic here and delete this line
	args := []int32{globalKey.PostVisitPublic}

	m, err := jwtx.ParseToken(r, l.svcCtx.Config.JWTAuth.AccessSecret)
	if err == nil {
		roleId, ok := m["roleId"]
		userId := m["userId"]
		if ok && (roleId <= 2 || userId == req.UserId) {
			// 如果是管理员或者自己，则都可以看见
			args = append(args, globalKey.PostVisitFriend, globalKey.PostVisitPrivate)
		}
		// TODO：如果是朋友，则需要再处理
	}

	postListResp, err := l.svcCtx.ContentRpc.GetPostList(l.ctx, &content_client.GetPostListReq{
		CreatorId:  req.UserId,
		Sort:       req.Sort,
		Visibility: args,
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
		postContentResp, err := l.svcCtx.ContentRpc.GetPostContentByPostId(l.ctx, &content_client.GetPostContentDetailReq{
			Id: item.Id,
		})
		if err != nil {
			cancel(err)
			return
		}
		var postItem types.PostItem
		_ = copier.Copy(&postItem, &item)
		var categoryItem types.PostCategory
		if postContentResp.PostContent.CategoryId != 0 {
			categoryResp, err := l.svcCtx.ContentRpc.GetCategoryDetail(l.ctx, &content_client.GetCategoryDetailReq{
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
	if postListResp.Total < uint64(total) {
		isEnd = true
	}
	return &types.GetUserPostListResp{
		List:  postList,
		Total: postListResp.Total,
		IsEnd: isEnd,
	}, nil
}
