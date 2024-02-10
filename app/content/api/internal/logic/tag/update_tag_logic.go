package tag

import (
	"context"
	"encoding/json"
	"errors"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTagLogic {
	return &UpdateTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTagLogic) UpdateTag(req *types.UpdateTagReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	roleId, _ := l.ctx.Value("roleId").(json.Number).Int64()
	tagResp, err := l.svcCtx.TagRpc.GetTagById(l.ctx, &content.GetTagByIdReq{
		Id: req.Id,
	})
	if userId != int64(tagResp.Tag.GetCreatorId()) && roleId > 2 {
		return errors.New("没有权限删除")
	}
	if _, err := l.svcCtx.TagRpc.DeleteTag(l.ctx, &content.DeleteTagReq{
		Id: req.Id,
	}); err != nil {
		return err
	}
	return nil
}
