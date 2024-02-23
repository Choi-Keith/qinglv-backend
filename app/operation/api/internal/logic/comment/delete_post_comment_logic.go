package comment

import (
	"context"
	"encoding/json"
	"errors"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePostCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostCommentLogic {
	return &DeletePostCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePostCommentLogic) DeletePostComment(req *types.DeletePostCommentReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	roleId, _ := l.ctx.Value("roleId").(json.Number).Int64()
	commentResp, err := l.svcCtx.CommentRpc.GetCommentById(l.ctx, &operation.GetCommentByIdReq{
		Id:   req.Id,
		Type: 1,
	})
	if err != nil {
		return err
	}
	if commentResp.PostComment.CreatorId != uint64(userId) && roleId > 2 {
		return errors.New("没有权限删除")
	}
	commentThumbResp, err := l.svcCtx.ThumbRpc.GetCommentThumbDetail(l.ctx, &operation.GetCommentThumbDetailReq{
		CommentId: req.Id,
		Type:      1,
	})
	if err != nil {
		return err
	}
	if _, err = l.svcCtx.CommentRpc.DeleteComment(l.ctx, &operation.DeleteCommentReq{
		CommentId:            req.Id,
		Type:                 1,
		PostCommentThumbList: commentThumbResp.Post,
	}); err != nil {
		return err
	}
	postResp, err := l.svcCtx.PostRpc.GetPostDetail(l.ctx, &content.GetPostDetailReq{
		Id: commentResp.PostComment.PostId,
	})
	if err != nil {
		return err
	}
	if _, err = l.svcCtx.PostRpc.UpdatePost(l.ctx, &content.UpdatePostReq{
		Id:           postResp.Post.Id,
		CommentCount: postResp.Post.CommentCount - 1,
	}); err != nil {
		return err
	}
	return nil
}
