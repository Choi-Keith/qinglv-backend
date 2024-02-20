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

type DeleteArticleCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteArticleCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleCommentLogic {
	return &DeleteArticleCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteArticleCommentLogic) DeleteArticleComment(req *types.DeleteArticleCommentReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	roleId, _ := l.ctx.Value("roleId").(json.Number).Int64()
	commentResp, err := l.svcCtx.CommentRpc.GetCommentById(l.ctx, &operation.GetCommentByIdReq{
		Id:   req.Id,
		Type: 2,
	})
	if err != nil {
		return err
	}
	if commentResp.ArticleComment.CreatorId != uint64(userId) && roleId > 2 {
		return errors.New("没有权限删除")
	}
	commentThumbResp, err := l.svcCtx.ThumbRpc.GetCommentThumbDetail(l.ctx, &operation.GetCommentThumbDetailReq{
		CommentId: req.Id,
		Type:      2,
	})
	if err != nil {
		return err
	}
	if _, err = l.svcCtx.CommentRpc.DeleteComment(l.ctx, &operation.DeleteCommentReq{
		Id:                      req.Id,
		Type:                    2,
		ArticleCommentThumbList: commentThumbResp.Article,
	}); err != nil {
		return err
	}
	articleResp, err := l.svcCtx.ArticleRpc.GetArticleDetail(l.ctx, &content.GetArticleDetailReq{
		Id: commentResp.ArticleComment.ArticleId,
	})
	if err != nil {
		return err
	}
	if _, err = l.svcCtx.PostRpc.UpdatePost(l.ctx, &content.UpdatePostReq{
		Id:           articleResp.Article.Id,
		CommentCount: articleResp.Article.CommentCount - 1,
	}); err != nil {
		return err
	}
	return nil
}
