package commentclasslogic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentReplyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentReplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentReplyListLogic {
	return &GetCommentReplyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: Comment
func (l *GetCommentReplyListLogic) GetCommentReplyList(in *operation.GetCommentReplyListReq) (*operation.GetCommentReplyListResp, error) {
	// todo: add your logic here and delete this line

	return &operation.GetCommentReplyListResp{}, nil
}
