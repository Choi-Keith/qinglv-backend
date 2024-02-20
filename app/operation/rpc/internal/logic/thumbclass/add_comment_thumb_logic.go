package thumbclasslogic

import (
	"context"
	"database/sql"
	"time"

	"qinglv-backend/app/operation/rpc/internal/model/thumb"
	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentThumbLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCommentThumbLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentThumbLogic {
	return &AddCommentThumbLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: CommentThumb
func (l *AddCommentThumbLogic) AddCommentThumb(in *operation.AddCommentThumbReq) (*operation.OkResp, error) {
	// todo: add your logic here and delete this line
	if in.Type == 1 {
		_, err := l.svcCtx.PostCommentThumbModel.Insert(l.ctx, nil, &thumb.PostCommentThumb{
			Id:          in.Id,
			PostId:      in.PostId,
			CommentId:   in.CommentId,
			CreatorId:   in.CreatorId,
			CreatorName: in.CreatorName,
			CommentType: int64(in.CommentType),
			Like:        int64(in.Like),
			Dislike:     int64(in.Dislike),
			DeletedAt:   time.Now(),
			Version:     1,
			ReplyId:     sql.NullInt64{Int64: int64(in.ReplyId), Valid: true},
		})
		if err != nil {
			return nil, err
		}
	}
	if in.Type == 2 {
		_, err := l.svcCtx.ArticleCommentThumbModel.Insert(l.ctx, nil, &thumb.ArticleCommentThumb{
			Id:          in.Id,
			ArticleId:   in.ArticleId,
			CommentId:   in.CommentId,
			CreatorId:   in.CreatorId,
			CreatorName: in.CreatorName,
			CommentType: int64(in.CommentType),
			Like:        int64(in.Like),
			Dislike:     int64(in.Dislike),
			DeletedAt:   time.Now(),
			Version:     1,
			ReplyId:     sql.NullInt64{Int64: int64(in.ReplyId), Valid: true},
		})
		if err != nil {
			return nil, err
		}
	}
	return &operation.OkResp{}, nil
}
