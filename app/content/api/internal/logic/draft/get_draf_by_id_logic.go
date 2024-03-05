package draft

import (
	"context"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetDrafByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDrafByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDrafByIdLogic {
	return &GetDrafByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDrafByIdLogic) GetDrafById(req *types.GetDraftByIdReq) (resp *types.GetDraftByIdResp, err error) {
	// todo: add your logic here and delete this line
	draftResp, err := l.svcCtx.DraftRpc.GetDraftById(l.ctx, &content.GetDraftByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	var draftItem types.DraftItem
	_ = copier.Copy(&draftItem, draftResp.Draft)

	return &types.GetDraftByIdResp{
		Draft: draftItem,
	}, nil
}
