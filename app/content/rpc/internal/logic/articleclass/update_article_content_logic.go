package articleclasslogic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateArticleContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateArticleContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleContentLogic {
	return &UpdateArticleContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateArticleContentLogic) UpdateArticleContent(in *content.UpdateArticleContentReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line
	articleItem, err := l.svcCtx.ArticleContentModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	title := articleItem.Title
	if in.Title != "" {
		title = in.Title
	}
	articleItem.Title = title
	introduction := articleItem.Introduction
	if in.Introduction != "" {
		introduction = in.Introduction
	}
	articleItem.Introduction = introduction
	coverImage := articleItem.CoverImage
	if in.CoverImage != "" {
		coverImage = articleItem.CoverImage
	}
	articleItem.CoverImage = coverImage
	contentStr := articleItem.Content
	if in.Content != "" {
		contentStr = in.Content
	}
	articleItem.Content = contentStr
	if err = l.svcCtx.ArticleContentModel.UpdateWithVersion(l.ctx, nil, articleItem); err != nil {
		return nil, err
	}
	return &content.OkResp{}, nil
}
