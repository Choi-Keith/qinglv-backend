package commentclasslogic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentLogic {
	return &UpdateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCommentLogic) UpdateComment(in *operation.UpdateCommentReq) (*operation.OkResp, error) {
	// todo: add your logic here and delete this line

	if in.Type == 1 {
		postResp, err := l.svcCtx.PostCommentModel.FindOne(l.ctx, in.Id)
		if err != nil {
			return nil, err
		}
		logx.Debugf("[UpdateComment] in: %+v\n", in)
		postResp.LikeCount = in.LikeCount
		postResp.DislikeCount = in.DislikeCount
		postResp.Score = in.Score
		l.svcCtx.PostCommentModel.UpdateWithVersion(l.ctx, nil, postResp)
	}
	if in.Type == 2 {
		articleResp, err := l.svcCtx.ArticleCommentModel.FindOne(l.ctx, in.Id)
		if err != nil {
			return nil, err
		}
		logx.Debugf("[UpdateComment] in: %+v\n", in)
		articleResp.LikeCount = in.LikeCount
		articleResp.DislikeCount = in.DislikeCount
		articleResp.Score = in.Score
		l.svcCtx.ArticleCommentModel.UpdateWithVersion(l.ctx, nil, articleResp)
	}
	return &operation.OkResp{}, nil
}
