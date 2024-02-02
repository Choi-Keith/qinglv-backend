package commentclasslogic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentReplyByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentReplyByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentReplyByIdLogic {
	return &GetCommentReplyByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: Comment
func (l *GetCommentReplyByIdLogic) GetCommentReplyById(in *operation.GetCommentReplyByIdReq) (*operation.GetCommentReplyByIdResp, error) {
	// todo: add your logic here and delete this line
	if in.Type == 1 {
		commentReplyResp, err := l.svcCtx.PostCommentReplyModel.FindOne(l.ctx, in.Id)
		if err != nil {
			return nil, err
		}
		commentReplyItem := genCommentReplyItem(commentReplyResp)
		return &operation.GetCommentReplyByIdResp{
			PostCommentReply: commentReplyItem,
		}, nil
	}
	return &operation.GetCommentReplyByIdResp{}, nil
}
