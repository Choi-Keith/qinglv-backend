package blacklistclasslogic

import (
	"context"
	"errors"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckBlackItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckBlackItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckBlackItemLogic {
	return &CheckBlackItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: blacklist
func (l *CheckBlackItemLogic) CheckBlackItem(in *user.CheckBlackItemReq) (*user.CheckBlackItemResp, error) {
	// todo: add your logic here and delete this line
	if in.UserId == 0 || in.BlackItemId == 0 {
		return nil, errors.New("blackItemId和userId不能为空")
	}
	whereBuilder := l.svcCtx.BlacklistModel.SelectBuilder()
	blacklist, err := l.svcCtx.BlacklistModel.FindAll(l.ctx, whereBuilder, "")
	if err != nil {
		return nil, err
	}
	return &user.CheckBlackItemResp{
		IsBlackItem: len(blacklist) > 0,
	}, nil
}
