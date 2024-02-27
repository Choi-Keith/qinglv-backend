package articleclasslogic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteArticleFeedByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteArticleFeedByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleFeedByIdsLogic {
	return &DeleteArticleFeedByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteArticleFeedByIdsLogic) DeleteArticleFeedByIds(in *content.DeleteArticleFeedByIdsReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line
	whereBuilder := l.svcCtx.ArticleFeedModel.SelectBuilder()
	if in.UserId != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"user_id": in.UserId,
		})
	}
	if in.AuthorId != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"author_id": in.AuthorId,
		})
	}
	if in.ArticleId != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"article_id": in.ArticleId,
		})
	}
	list, _ := l.svcCtx.ArticleFeedModel.FindAll(l.ctx, whereBuilder, "")
	for _, articleItem := range list {
		l.svcCtx.ArticleFeedModel.Delete(l.ctx, nil, articleItem.Id)
	}
	return &content.OkResp{}, nil
}
