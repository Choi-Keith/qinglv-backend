package commentclasslogic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type DeleteCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: Comment
func (l *DeleteCommentLogic) DeleteComment(in *operation.DeleteCommentReq) (*operation.OkResp, error) {
	// todo: add your logic here and delete this line
	if in.Type == 1 {
		whereBuilder := l.svcCtx.PostCommentReplyModel.SelectBuilder().Where(squirrel.Eq{
			"comment_id": in.Id,
		})
		postCommentReplyResp, err := l.svcCtx.PostCommentReplyModel.FindAll(l.ctx, whereBuilder, "")
		if err != nil {
			return nil, err
		}
		l.svcCtx.PostCommentModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
			for _, item := range postCommentReplyResp {
				if err := l.svcCtx.PostCommentReplyModel.Delete(ctx, session, item.Id); err != nil {
					return err
				}
			}
			if err := l.svcCtx.PostCommentModel.Delete(ctx, session, in.Id); err != nil {
				return err
			}
			for _, postCommentThumbItem := range in.PostCommentThumbList {
				if err := l.svcCtx.PostCommentThumbModel.Delete(l.ctx, session, postCommentThumbItem.Id); err != nil {
					return err
				}
			}
			return nil
		})

	}
	if in.Type == 2 {
		whereBuilder := l.svcCtx.ArticleCommentReplyModel.SelectBuilder().Where(squirrel.Eq{
			"comment_id": in.Id,
		})
		articleCommentReplyResp, err := l.svcCtx.ArticleCommentReplyModel.FindAll(l.ctx, whereBuilder, "")
		if err != nil {
			return nil, err
		}
		l.svcCtx.ArticleCommentModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
			for _, item := range articleCommentReplyResp {
				if err := l.svcCtx.ArticleCommentReplyModel.Delete(ctx, session, item.Id); err != nil {
					return err
				}
			}
			if err := l.svcCtx.ArticleCommentModel.Delete(ctx, session, in.Id); err != nil {
				return err
			}
			for _, articleCommentThumbItem := range in.ArticleCommentThumbList {
				if err := l.svcCtx.ArticleCommentThumbModel.Delete(l.ctx, session, articleCommentThumbItem.Id); err != nil {
					return err
				}
			}
			return nil
		})
	}
	return &operation.OkResp{}, nil
}
