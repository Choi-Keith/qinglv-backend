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

type DeleteArticleCommentReplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteArticleCommentReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleCommentReplyLogic {
	return &DeleteArticleCommentReplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteArticleCommentReplyLogic) DeleteArticleCommentReply(req *types.DeleteArticleCommentReplyReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	roleId, _ := l.ctx.Value("roleId").(json.Number).Int64()
	commentResp, err := l.svcCtx.CommentRpc.GetCommentReplyById(l.ctx, &operation.GetCommentReplyByIdReq{
		Id:   req.Id,
		Type: 2,
	})
	if err != nil {
		return err
	}
	if commentResp.ArticleCommentReply.CreatorId != uint64(userId) && roleId > 2 {
		return errors.New("没有权限删除")
	}
	articleCommentThumbResp, err := l.svcCtx.ThumbRpc.GetCommentThumbDetail(l.ctx, &operation.GetCommentThumbDetailReq{
		ReplyId: req.Id,
		Type:    2,
	})
	if err != nil {
		return err
	}
	if _, err = l.svcCtx.CommentRpc.DeleteCommentReply(l.ctx, &operation.DeleteCommentReplyReq{
		Id:                      req.Id,
		Type:                    2,
		ArticleCommentThumbList: articleCommentThumbResp.Article,
	}); err != nil {
		return err
	}
	return nil
}
