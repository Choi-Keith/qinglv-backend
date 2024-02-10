package articleclasslogic

import (
	"context"
	"database/sql"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/model/article"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type AddArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddArticleLogic {
	return &AddArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddArticleLogic) AddArticle(in *content.AddArticleReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.ArticleModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		if _, err := l.svcCtx.ArticleModel.Insert(l.ctx, nil, &article.Article{
			Id:          in.Id,
			CreatorId:   in.CreatorId,
			Visibility:  1,
			Status:      1,
			IsTop:       1,
			Version:     1,
			Score:       10.000,
			CreatorName: in.CreatorName,
		}); err != nil {
			return err
		}

		if _, err := l.svcCtx.ArticleContentModel.Insert(l.ctx, nil, &article.ArticleContent{
			Id:           in.ArticleContentId,
			Title:        in.Title,
			Introduction: in.Introduction,
			CategoryId:   sql.NullInt64{Int64: int64(in.CategoryId), Valid: true},
			Tags:         sql.NullString{String: in.Tags, Valid: true},
			CoverImage:   sql.NullString{String: in.CoverImage, Valid: true},
			Content:      in.Content,
			CreatorId:    in.CreatorId,
			Version:      1,
			CreatorName:  in.CreatorName,
		}); err != nil {
			return err
		}
		return nil

	})
	if err != nil {
		return nil, err
	}

	return &content.OkResp{}, nil
}
