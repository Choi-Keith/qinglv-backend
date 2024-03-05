package draft

import (
	"context"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDraftLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDraftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDraftLogic {
	return &UpdateDraftLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDraftLogic) UpdateDraft(req *types.UpdateDraftReq) error {
	// todo: add your logic here and delete this line
	if _, err := l.svcCtx.DraftRpc.UpdateDraft(l.ctx, &content.UpdateDraftReq{
		Id:      req.Id,
		Title:   req.Title,
		Content: req.Content,
	}); err != nil {
		return err
	}
	return nil
}
