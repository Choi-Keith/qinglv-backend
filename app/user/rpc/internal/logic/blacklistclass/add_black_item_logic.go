package blacklistclasslogic

import (
	"context"
	"time"

	"qinglv-backend/app/user/rpc/internal/model/black"
	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddBlackItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddBlackItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBlackItemLogic {
	return &AddBlackItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: blacklist
func (l *AddBlackItemLogic) AddBlackItem(in *user.AddBlackItemReq) (*user.AddBlackItemResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.BlacklistModel.Insert(l.ctx, nil, &black.Blacklist{
		Id:        in.Id,
		UserId:    in.UserId,
		BlackId:   in.BlackItemId,
		DeletedAt: time.Now(),
		Version:   1,
	})
	if err != nil {
		return nil, err
	}
	return &user.AddBlackItemResp{}, nil
}
