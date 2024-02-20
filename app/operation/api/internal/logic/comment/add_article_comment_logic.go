package comment

import (
	"context"
	"encoding/json"
	"net/http"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"
	"qinglv-backend/app/operation/rpc/operation"
	"qinglv-backend/app/user/rpc/user"
	"qinglv-backend/pkg/snowflake"
	"qinglv-backend/pkg/utils"

	"github.com/techxmind/ip2location"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddArticleCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewAddArticleCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *AddArticleCommentLogic {
	return &AddArticleCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *AddArticleCommentLogic) AddArticleComment(req *types.AddArticleCommentReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	creatorName := l.ctx.Value("nickname").(string)
	id := snowflake.MustID()
	ip := utils.GetClientIP(l.r)
	loc, err := ip2location.Get(ip)
	if err != nil {
		return err
	}
	articleResp, err := l.svcCtx.ArticleRpc.GetArticleDetail(l.ctx, &content.GetArticleDetailReq{
		Id: req.ArticleId,
	})
	if err != nil {
		return err
	}
	if _, err = l.svcCtx.CommentRpc.AddComment(l.ctx, &operation.AddCommentReq{
		Id:          id,
		ArticleId:   req.ArticleId,
		Content:     req.Content,
		CreatorId:   uint64(userId),
		Location:    loc.Province,
		CreatorName: creatorName,
		Type:        2,
	}); err != nil {
		return err
	}
	score := utils.HandleScore(articleResp.Article.CreatedAt, 2, 1.5)
	if _, err = l.svcCtx.ArticleRpc.UpdateArticle(l.ctx, &content.UpdateArticleReq{
		Id:           req.ArticleId,
		CommentCount: articleResp.Article.CommentCount + 1,
		Score:        articleResp.Article.Score + score,
	}); err != nil {
		return err
	}
	if _, err := l.svcCtx.UserRpc.UpdateUserScoreLevel(l.ctx, &user.UpdateUserScoreLevelReq{
		Id:    articleResp.Article.CreatorId,
		Score: 1,
		Op:    "add",
	}); err != nil {
		return err
	}
	return nil
}
