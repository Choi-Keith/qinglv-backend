package articleclasslogic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type DeleteArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleLogic {
	return &DeleteArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteArticleLogic) DeleteArticle(in *content.DeleteArticleReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.ArticleModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		articleResp, err := l.svcCtx.ArticleModel.FindOne(ctx, in.Id)
		if err != nil {
			return err
		}
		if err = l.svcCtx.ArticleModel.DeleteSoft(ctx, session, articleResp); err != nil {
			return err
		}
		articleContentResp, err := l.svcCtx.ArticleContentModel.FindOneByArticleId(ctx, articleResp.Id)
		if err != nil {
			return err
		}
		if err = l.svcCtx.ArticleContentModel.DeleteSoft(ctx, session, articleContentResp); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &content.OkResp{}, nil
}
