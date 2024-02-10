package article

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/user/rpc/user"
	"qinglv-backend/common/globalKey"
	"qinglv-backend/pkg/jwtx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type GetArticleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetArticleListLogic {
	return &GetArticleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *GetArticleListLogic) GetArticleList(req *types.GetArticleListReq) (resp *types.GetArticleListResp, err error) {
	// todo: add your logic here and delete this line

	args := []int32{globalKey.PostVisitPublic}
	articleListResp, err := l.svcCtx.ArticleRpc.GetArticleList(l.ctx, &content.GetArticleListReq{
		Status:      int32(req.Status),
		Visibility:  args,
		Score:       float64(req.Score),
		Sort:        req.Sort,
		CreatorName: req.Creator,
		PageNum:     uint64(req.PageNum),
		PageSize:    uint64(req.PageSize),
	})
	if err != nil {
		return nil, err
	}
	articleList, err := mr.MapReduce(func(source chan<- content.ArticleItem) {
		for _, articleItem := range articleListResp.Data {
			source <- *articleItem
		}
	}, func(item content.ArticleItem, writer mr.Writer[types.ArticleItem], cancel func(error)) {
		articleContentResp, err := l.svcCtx.ArticleRpc.GetArticleContentByArticleId(l.ctx, &content.GetArticleContentDetailReq{
			Id: item.Id,
		})
		if err != nil {
			cancel(err)
			return
		}
		userResp, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.GetUserByIdReq{
			UserId: item.CreatorId,
		})
		if err != nil {
			cancel(err)
			return
		}
		var articleItem types.ArticleItem
		_ = copier.Copy(&articleItem, &item)
		_ = copier.Copy(&articleItem.Creator, userResp.User)
		m, err := jwtx.ParseToken(l.r, l.svcCtx.Config.JWTAuth.AccessSecret)
		if err == nil {
			fmt.Printf("m: %+v\n", m)
			userId, ok := m["userId"]
			if ok && userId != 0 {
				followingResp, err := l.svcCtx.FollowingRpc.GetFollowingList(l.ctx, &user.GetFollowingListReq{
					UserId:      uint64(userId),
					FollowingId: articleItem.Creator.Id,
				})
				if err != nil {
					cancel(err)
					return
				}
				if len(followingResp.Data) != 0 {
					articleItem.Creator.IsFollowing = true
				}
			}
		}
		var categoryItem types.ArticleCategory
		if articleContentResp.ArticleContent.CategoryId != 0 {
			categoryResp, err := l.svcCtx.CategoryRpc.GetCategoryDetail(l.ctx, &content.GetCategoryDetailReq{
				Id: articleContentResp.ArticleContent.CategoryId,
			})
			if err != nil {
				cancel(err)
				return
			}
			_ = copier.Copy(&categoryItem, categoryResp.Category)
		}
		articleItem.Category = categoryItem
		topics := strings.Split(articleContentResp.ArticleContent.Tags, ",")
		articleItem.Tag = topics
		articleItem.Content = articleContentResp.ArticleContent.Content
		articleItem.CoverImage = articleContentResp.ArticleContent.CoverImage
		writer.Write(articleItem)
	}, func(pipe <-chan types.ArticleItem, writer mr.Writer[[]types.ArticleItem], cancel func(error)) {
		var r []types.ArticleItem
		m := make(map[uint64]types.ArticleItem, len(articleListResp.Data))
		for p := range pipe {
			m[p.Id] = p
		}
		for _, ArticleItem := range articleListResp.Data {
			r = append(r, m[ArticleItem.Id])
		}
		writer.Write(r)
	})
	if err != nil {
		return nil, err
	}
	isEnd := false
	total := (req.PageNum-1)*req.PageSize + req.PageSize
	if articleListResp.Total <= uint64(total) {
		isEnd = true
	}

	return &types.GetArticleListResp{
		List:  articleList,
		Total: articleListResp.Total,
		IsEnd: isEnd,
	}, nil
}
