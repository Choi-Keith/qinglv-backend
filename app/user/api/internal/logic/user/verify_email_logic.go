package user

import (
	"context"
	"encoding/json"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user"
	"qinglv-backend/common/schema"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyEmailLogic {
	return &VerifyEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyEmailLogic) VerifyEmail(req *types.VerifyEmailReq) (resp *types.VerifyEmailResp, err error) {
	// todo: add your logic here and delete this line
	verifyRegisterCodeResp, err := l.svcCtx.EmailRpc.VerifyRegisterCode(l.ctx, &user.VerifyRegisterCodeReq{
		Code: req.Code,
	})
	if err != nil {
		return nil, err
	}
	codeContent := new(schema.EmailContent)
	if err := json.Unmarshal([]byte(verifyRegisterCodeResp.CodeContent), codeContent); err != nil {
		return nil, err
	}
	_, err = l.svcCtx.UserRpc.UpdateEmailStatus(l.ctx, &user.UpdateEmailStatusReq{
		UserId:     codeContent.UserId,
		MailStatus: 2,
	})
	if err != nil {
		return nil, err
	}
	return &types.VerifyEmailResp{
		VerifyResult: true,
	}, nil
}
