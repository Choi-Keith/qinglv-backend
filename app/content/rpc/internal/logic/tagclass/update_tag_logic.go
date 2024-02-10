package tagclasslogic

import (
	"context"
	"database/sql"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTagLogic {
	return &UpdateTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTagLogic) UpdateTag(in *content.UpdateTagReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line
	tagItem, err := l.svcCtx.TagModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	image := tagItem.Image
	if in.Image != "" {
		image = in.Image
	}
	tagItem.Image = image
	description := tagItem.Description.String
	if in.Description != "" {
		description = in.Description
	}
	tagItem.Description = sql.NullString{String: description, Valid: true}
	quoteCount := tagItem.QuoteCount
	if in.QuoteCount != 0 {
		quoteCount = in.QuoteCount
	}
	tagItem.QuoteCount = quoteCount
	l.svcCtx.TagModel.UpdateWithVersion(l.ctx, nil, tagItem)
	return &content.OkResp{}, nil
}
