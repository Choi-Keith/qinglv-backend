package draftclasslogic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDraftLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDraftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDraftLogic {
	return &DeleteDraftLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteDraftLogic) DeleteDraft(in *content.DeleteDraftReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line
	draftResp, err := l.svcCtx.DraftModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	if err = l.svcCtx.DraftModel.DeleteSoft(l.ctx, nil, draftResp); err != nil {
		return nil, err
	}
	return &content.OkResp{}, nil
}
