package article

import (
	"context"
	"net/http"
	"strings"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/user/rpc/user"
	"qinglv-backend/pkg/jwtx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetArticleByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetArticleByIdLogic {
	return &GetArticleByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *GetArticleByIdLogic) GetArticleById(req *types.GetArticleByIdReq) (resp *types.GetArticleByIdResp, err error) {
	// todo: add your logic here and delete this line
	articleResp, err := l.svcCtx.ArticleRpc.GetArticleDetail(l.ctx, &content.GetArticleDetailReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	articleContentResp, err := l.svcCtx.ArticleRpc.GetArticleContentByArticleId(l.ctx, &content.GetArticleContentDetailReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	userResp, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.GetUserByIdReq{
		UserId: articleResp.Article.CreatorId,
	})
	if err != nil {
		return nil, err
	}
	var articleItem types.ArticleItem
	_ = copier.Copy(&articleItem, articleResp.Article)
	_ = copier.Copy(&articleItem.Creator, userResp.User)
	m, err := jwtx.ParseToken(l.r, l.svcCtx.Config.JWTAuth.AccessSecret)
	if err == nil {
		userId, ok := m["userId"]
		if ok && userId != 0 {
			followingResp, err := l.svcCtx.FollowingRpc.GetFollowingList(l.ctx, &user.GetFollowingListReq{
				UserId:      uint64(userId),
				FollowingId: articleItem.Creator.Id,
			})
			if err != nil {
				return nil, err
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
			logx.Errorf("[Post] GetCategoryDetail failed: %+v\n", err)
			return nil, err
		}
		_ = copier.Copy(&categoryItem, categoryResp.Category)
	}
	articleItem.Category = categoryItem
	tags := strings.Split(articleContentResp.ArticleContent.Tags, ",")
	articleItem.Tag = tags
	return &types.GetArticleByIdResp{
		Article: articleItem,
	}, nil
}
