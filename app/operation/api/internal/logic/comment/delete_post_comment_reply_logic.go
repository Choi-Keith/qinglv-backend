package comment

import (
	"context"
	"encoding/json"
	"errors"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostCommentReplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePostCommentReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostCommentReplyLogic {
	return &DeletePostCommentReplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePostCommentReplyLogic) DeletePostCommentReply(req *types.DeletePostCommentReplyReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	roleId, _ := l.ctx.Value("roleId").(json.Number).Int64()
	commentResp, err := l.svcCtx.CommentRpc.GetCommentReplyById(l.ctx, &operation.GetCommentReplyByIdReq{
		Id:   req.Id,
		Type: 1,
	})
	if err != nil {
		return err
	}
	if commentResp.PostCommentReply.CreatorId != uint64(userId) && roleId > 2 {
		return errors.New("没有权限删除")
	}
	if _, err = l.svcCtx.CommentRpc.DeleteCommentReply(l.ctx, &operation.DeleteCommentReplyReq{
		Id:   req.Id,
		Type: 1,
	}); err != nil {
		return err
	}
	return nil
}
