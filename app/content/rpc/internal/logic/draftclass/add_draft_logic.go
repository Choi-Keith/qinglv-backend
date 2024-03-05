package draftclasslogic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/model/draft"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDraftLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddDraftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDraftLogic {
	return &AddDraftLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddDraftLogic) AddDraft(in *content.AddDraftReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line
	if _, err := l.svcCtx.DraftModel.Insert(l.ctx, nil, &draft.ArticleDraft{
		Id:          in.Id,
		Content:     in.Content,
		Title:       in.Title,
		CreatorId:   in.CreatorId,
		CreatorName: in.CreatorName,
		IsDel:       0,
		Version:     1,
	}); err != nil {
		return nil, err
	}
	return &content.OkResp{}, nil
}
