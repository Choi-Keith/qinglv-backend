package draftclasslogic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	draftModel "qinglv-backend/app/content/rpc/internal/model/draft"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDraftByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDraftByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDraftByIdLogic {
	return &GetDraftByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDraftByIdLogic) GetDraftById(in *content.GetDraftByIdReq) (*content.GetDraftByIdResp, error) {
	// todo: add your logic here and delete this line
	draftResp, err := l.svcCtx.DraftModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	draftItem := genDraftItem(draftResp)
	return &content.GetDraftByIdResp{
		Draft: draftItem,
	}, nil
}

func genDraftItem(item *draftModel.ArticleDraft) *content.DraftItem {
	return &content.DraftItem{
		Id:          item.Id,
		Title:       item.Title,
		Content:     item.Content,
		CreatorId:   item.CreatorId,
		CreatorName: item.CreatorName,
		CreatedAt:   uint64(item.CreatedAt.Unix() * 1000),
		UpdatedAt:   uint64(item.UpdatedAt.Unix() * 1000),
	}
}
