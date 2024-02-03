package commentclasslogic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentReplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentReplyLogic {
	return &UpdateCommentReplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCommentReplyLogic) UpdateCommentReply(in *operation.UpdateCommentReplyReq) (*operation.OkResp, error) {
	// todo: add your logic here and delete this line
	if in.Type == 1 {
		postReplyItem, err := l.svcCtx.PostCommentReplyModel.FindOne(l.ctx, in.Id)
		if err != nil {
			return nil, err
		}
		postReplyItem.LikeCount = in.LikeCount
		postReplyItem.DislikeCount = in.DislikeCount
		postReplyItem.Score = in.Score
		l.svcCtx.PostCommentReplyModel.UpdateWithVersion(l.ctx, nil, postReplyItem)
	}
	return &operation.OkResp{}, nil
}
