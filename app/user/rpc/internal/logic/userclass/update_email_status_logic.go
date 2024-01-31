package userclasslogic

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEmailStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateEmailStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEmailStatusLogic {
	return &UpdateEmailStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: user
func (l *UpdateEmailStatusLogic) UpdateEmailStatus(in *user.UpdateEmailStatusReq) (*user.UpdateEmailStatusResp, error) {
	// todo: add your logic here and delete this line
	userItem, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	userItem.MailStatus = int64(in.MailStatus)
	if err := l.svcCtx.UserModel.UpdateWithVersion(l.ctx, nil, userItem); err != nil {
		return nil, err
	}
	return &user.UpdateEmailStatusResp{}, nil
}
