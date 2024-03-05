package draft

import (
	"context"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetDraftListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDraftListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDraftListLogic {
	return &GetDraftListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDraftListLogic) GetDraftList(req *types.GetDraftListReq) (resp *types.GetDraftListResp, err error) {
	// todo: add your logic here and delete this line
	draftResp, err := l.svcCtx.DraftRpc.GetDraftList(l.ctx, &content.GetDraftListReq{
		Title:       req.Title,
		CreatorName: req.Creator,
		PageNum:     int64(req.PageNum),
		PageSize:    int64(req.PageSize),
	})
	if err != nil {
		return nil, err
	}
	draftList := make([]types.DraftItem, len(draftResp.Data))
	for idx, item := range draftResp.Data {
		_ = copier.Copy(&draftList[idx], item)
	}
	isEnd := false
	total := (req.PageNum-1)*req.PageSize + req.PageSize
	if draftResp.Total <= int64(total) {
		isEnd = true
	}
	return &types.GetDraftListResp{
		List:  draftList,
		Total: uint64(draftResp.Total),
		IsEnd: isEnd,
	}, nil
}
