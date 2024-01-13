package black

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBlackListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBlackListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBlackListLogic {
	return &GetBlackListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: blacklist
func (l *GetBlackListLogic) GetBlackList(in *user.GetBlackListReq) (*user.GetBlackListResp, error) {
	// todo: add your logic here and delete this line

	return &user.GetBlackListResp{}, nil
}
