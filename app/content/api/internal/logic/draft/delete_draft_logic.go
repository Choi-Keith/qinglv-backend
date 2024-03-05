package draft

import (
	"context"
	"encoding/json"
	"errors"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDraftLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDraftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDraftLogic {
	return &DeleteDraftLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDraftLogic) DeleteDraft(req *types.DeleteDraftReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	roleId, _ := l.ctx.Value("roleId").(json.Number).Int64()
	draftResp, err := l.svcCtx.DraftRpc.GetDraftById(l.ctx, &content.GetDraftByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return err
	}
	if draftResp.Draft.CreatorId != uint64(userId) && roleId > 2 {
		return errors.New("没有权限删除")
	}
	if _, err = l.svcCtx.DraftRpc.DeleteDraft(l.ctx, &content.DeleteDraftReq{
		Id: req.Id,
	}); err != nil {
		return err
	}
	return nil
}
