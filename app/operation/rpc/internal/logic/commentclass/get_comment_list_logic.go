package commentclasslogic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: Comment
func (l *GetCommentListLogic) GetCommentList(in *operation.GetCommentListReq) (*operation.GetCommentListResp, error) {
	// todo: add your logic here and delete this line

	return &operation.GetCommentListResp{}, nil
}
