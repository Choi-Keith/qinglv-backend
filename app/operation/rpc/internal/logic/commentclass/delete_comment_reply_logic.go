package commentclasslogic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type DeleteCommentReplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentReplyLogic {
	return &DeleteCommentReplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: Comment
func (l *DeleteCommentReplyLogic) DeleteCommentReply(in *operation.DeleteCommentReplyReq) (*operation.OkResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.PostCommentReplyModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		err := l.svcCtx.PostCommentReplyModel.Delete(ctx, session, in.Id)
		if err != nil {
			return err
		}
		for _, postCommentThumbItem := range in.PostCommentThumbList {
			err := l.svcCtx.PostCommentThumbModel.Delete(ctx, session, postCommentThumbItem.Id)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &operation.OkResp{}, nil
}
