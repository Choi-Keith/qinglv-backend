package tag

import (
	"context"
	"encoding/json"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/pkg/snowflake"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTagLogic {
	return &AddTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTagLogic) AddTag(req *types.AddTagReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	roleId, _ := l.ctx.Value("roleId").(json.Number).Int64()
	creatorName := l.ctx.Value("nickname").(string)
	createdType := 1
	if roleId > 2 {
		createdType = 2
	}
	id := snowflake.MustID()
	if _, err = l.svcCtx.TagRpc.AddTag(l.ctx, &content.AddTagReq{
		Id:          id,
		CreatorId:   uint64(userId),
		Name:        req.Name,
		Description: req.Description,
		Image:       req.Image,
		Type:        int32(createdType),
		CreatorName: creatorName,
	}); err != nil {
		return err
	}
	return nil
}
