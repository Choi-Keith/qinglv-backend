package commentclasslogic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentAllLogic {
	return &GetCommentAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentAllLogic) GetCommentAll(in *operation.GetCommentAllReq) (*operation.GetCommentAllResp, error) {
	// todo: add your logic here and delete this line
	if in.Type == 1 {
		whereBuilder := l.svcCtx.PostCommentModel.SelectBuilder()
		if in.PostId != 0 {
			whereBuilder = whereBuilder.Where(squirrel.Eq{
				"post_id": in.PostId,
			})
		}
		commentListResp, err := l.svcCtx.PostCommentModel.FindAll(l.ctx, whereBuilder, "")
		if err != nil {
			return nil, err
		}
		commentList := make([]*operation.PostCommentItem, len(commentListResp))
		for idx, commentItem := range commentListResp {
			commentList[idx] = genPostCommentItem(commentItem)
		}
		return &operation.GetCommentAllResp{
			Posts:    commentList,
			Articles: nil,
		}, nil
	}
	if in.Type == 2 {
		whereBuilder := l.svcCtx.ArticleCommentModel.SelectBuilder()
		if in.ArticleId != 0 {
			whereBuilder = whereBuilder.Where(squirrel.Eq{
				"article_id": in.PostId,
			})
		}
		commentListResp, err := l.svcCtx.ArticleCommentModel.FindAll(l.ctx, whereBuilder, "")
		if err != nil {
			return nil, err
		}
		commentList := make([]*operation.ArticleCommentItem, len(commentListResp))
		for idx, commentItem := range commentListResp {
			commentList[idx] = genArticleCommentItem(commentItem)
		}
		return &operation.GetCommentAllResp{
			Posts:    nil,
			Articles: commentList,
		}, nil
	}
	return &operation.GetCommentAllResp{}, nil
}
