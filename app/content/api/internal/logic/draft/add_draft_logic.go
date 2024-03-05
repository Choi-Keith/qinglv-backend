package draft

import (
	"context"
	"encoding/json"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/pkg/snowflake"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDraftLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddDraftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDraftLogic {
	return &AddDraftLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddDraftLogic) AddDraft(req *types.AddDraftReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	creatorName := l.ctx.Value("nickname").(string)
	if _, err := l.svcCtx.DraftRpc.AddDraft(l.ctx, &content.AddDraftReq{
		Id:          snowflake.MustID(),
		CreatorId:   uint64(userId),
		CreatorName: creatorName,
		Title:       req.Title,
		Content:     req.Content,
	}); err != nil {
		return err
	}
	return nil
}
