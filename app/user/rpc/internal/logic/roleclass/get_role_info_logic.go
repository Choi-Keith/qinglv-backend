package roleclasslogic

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleInfoLogic {
	return &GetRoleInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: role
func (l *GetRoleInfoLogic) GetRoleInfo(in *user.GetRoleInfoReq) (*user.GetRoleInfoResp, error) {
	// todo: add your logic here and delete this line
	roleItem, err := l.svcCtx.RoleModel.FindOne(l.ctx, in.Id)
	if err != nil {
		logx.Errorf("[Role RPC] GetRoleInfo failed: %v\n", err)
		return nil, err
	}
	item := &user.RoleItem{
		Id:        roleItem.Id,
		Name:      roleItem.Name,
		CreatedAt: uint64(roleItem.CreatedAt.Unix() * 1000),
		UpdatedAt: uint64(roleItem.UpdatedAt.Unix() * 1000),
	}
	logx.Debugf("[Role RPC] GetRoleInfo item: %+v, CresatedAt: %v\n", item, roleItem)
	return &user.GetRoleInfoResp{
		Role: item,
	}, nil
}
