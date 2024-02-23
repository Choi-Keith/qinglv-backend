package shareclasslogic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostShareLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePostShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostShareLogic {
	return &DeletePostShareLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePostShareLogic) DeletePostShare(in *operation.DeletePostShareReq) (*operation.OkResp, error) {
	// todo: add your logic here and delete this line
	if err := l.svcCtx.PostShareModel.Delete(l.ctx, nil, in.Id); err != nil {
		return nil, err
	}
	return &operation.OkResp{}, nil
}
