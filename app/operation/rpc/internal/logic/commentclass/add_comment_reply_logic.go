package commentclasslogic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/model/comment"
	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentReplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCommentReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentReplyLogic {
	return &AddCommentReplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: Comment
func (l *AddCommentReplyLogic) AddCommentReply(in *operation.AddCommentReplyReq) (*operation.OkResp, error) {
	// todo: add your logic here and delete this line
	if in.Type == 1 {
		l.svcCtx.PostCommentReplyModel.Insert(l.ctx, nil, &comment.PostCommentReply{
			Id:           in.Id,
			PostId:       in.PostId,
			CommentId:    in.CommentId,
			CreatorId:    in.CreatorId,
			CreatorName:  in.CreatorName,
			AtUserId:     in.AtUserId,
			AtUserName:   in.AtUserName,
			Content:      in.Content,
			Location:     in.Location,
			Version:      1,
			LikeCount:    0,
			DislikeCount: 0,
		})
	}
	if in.Type == 2 {
		l.svcCtx.ArticleCommentReplyModel.Insert(l.ctx, nil, &comment.ArticleCommentReply{
			Id:           in.Id,
			ArticleId:    in.ArticleId,
			CommentId:    in.CommentId,
			CreatorId:    in.CreatorId,
			CreatorName:  in.CreatorName,
			AtUserId:     in.AtUserId,
			AtUserName:   in.AtUserName,
			Content:      in.Content,
			Location:     in.Location,
			Version:      1,
			LikeCount:    0,
			DislikeCount: 0,
		})
	}
	return &operation.OkResp{}, nil
}
