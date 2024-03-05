package draftclasslogic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDraftLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDraftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDraftLogic {
	return &UpdateDraftLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDraftLogic) UpdateDraft(in *content.UpdateDraftReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line
	draftResp, err := l.svcCtx.DraftModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	title := draftResp.Title
	if in.Title != "" {
		title = in.Title
	}
	newContent := draftResp.Content
	if in.Content != "" {
		newContent = in.Content
	}
	draftResp.Title = title
	draftResp.Content = newContent
	if _, err := l.svcCtx.DraftModel.Update(l.ctx, nil, draftResp); err != nil {
		return nil, err
	}
	return &content.OkResp{}, nil
}
