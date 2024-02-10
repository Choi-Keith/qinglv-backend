package articleclasslogic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	articleContentModel "qinglv-backend/app/content/rpc/internal/model/article"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleContentByArticleIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleContentByArticleIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleContentByArticleIdLogic {
	return &GetArticleContentByArticleIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetArticleContentByArticleIdLogic) GetArticleContentByArticleId(in *content.GetArticleContentDetailReq) (*content.GetArticleContentDetailResp, error) {
	// todo: add your logic here and delete this line
	articleContentResp, err := l.svcCtx.ArticleContentModel.FindOneByArticleId(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	articleContentItem := genArticleContentItem(articleContentResp)
	return &content.GetArticleContentDetailResp{
		ArticleContent: articleContentItem,
	}, nil
}

func genArticleContentItem(item *articleContentModel.ArticleContent) *content.ArticleContentItem {
	return &content.ArticleContentItem{
		Id:           item.Id,
		CreatorId:    item.CreatorId,
		Title:        item.Title,
		Introduction: item.Introduction,
		CoverImage:   item.CoverImage.String,
		CategoryId:   uint64(item.CategoryId.Int64),
		Tags:         item.Tags.String,
		Content:      item.Content,
		CreatedAt:    uint64(item.CreatedAt.Unix() * 1000),
		UpdatedAt:    uint64(item.UpdatedAt.Unix() * 1000),
	}
}
