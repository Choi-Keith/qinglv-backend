package draftclasslogic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetDraftListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDraftListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDraftListLogic {
	return &GetDraftListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDraftListLogic) GetDraftList(in *content.GetDraftListReq) (*content.GetDraftListResp, error) {
	// todo: add your logic here and delete this line
	whereBuilder := l.svcCtx.DraftModel.SelectBuilder()
	if in.CreatorName != "" {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"creator_name": in.CreatorName,
		})
	}
	if in.Title != "" {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"title": in.Title,
		})
	}
	draftListResp, total, err := l.svcCtx.DraftModel.FindPageListByPageWithTotal(l.ctx, whereBuilder, in.PageNum, in.PageSize, "")
	if err != nil {
		return nil, err
	}
	draftList := make([]*content.DraftItem, len(draftListResp))
	for idx, item := range draftListResp {
		draftList[idx] = genDraftItem(item)
	}
	return &content.GetDraftListResp{
		Data:  draftList,
		Total: total,
	}, nil
}
