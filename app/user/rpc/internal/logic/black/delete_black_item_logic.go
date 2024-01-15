package black

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBlackItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteBlackItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBlackItemLogic {
	return &DeleteBlackItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: blacklist
func (l *DeleteBlackItemLogic) DeleteBlackItem(in *user.DeleteBlackItemReq) (*user.DeleteBlackItemResp, error) {
	// todo: add your logic here and delete this line
	blackItem, err := l.svcCtx.BlacklistModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.BlacklistModel.Delete(l.ctx, nil, blackItem.Id)
	if err != nil {
		return nil, err
	}
	return &user.DeleteBlackItemResp{}, nil
}
