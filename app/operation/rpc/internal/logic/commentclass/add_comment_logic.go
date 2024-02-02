package commentclasslogic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/model/comment"
	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: Comment
func (l *AddCommentLogic) AddComment(in *operation.AddCommentReq) (*operation.OkResp, error) {
	// todo: add your logic here and delete this line
	if in.Type == 1 {
		_, err := l.svcCtx.PostCommentModel.Insert(l.ctx, nil, &comment.PostComment{
			Id:           in.Id,
			CreatorId:    in.CreatorId,
			PostId:       in.PostId,
			CreatorName:  in.CreatorName,
			Location:     in.Location,
			LikeCount:    0,
			Content:      in.Content,
			DislikeCount: 0,
			Score:        0,
			Version:      1,
		})
		if err != nil {
			return nil, err
		}
	}
	return &operation.OkResp{}, nil
}
