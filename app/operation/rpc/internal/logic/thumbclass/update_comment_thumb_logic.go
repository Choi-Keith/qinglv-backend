package thumbclasslogic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentThumbLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentThumbLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentThumbLogic {
	return &UpdateCommentThumbLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: PostCommentThumb
func (l *UpdateCommentThumbLogic) UpdateCommentThumb(in *operation.UpdateCommentThumbReq) (*operation.OkResp, error) {
	// todo: add your logic here and delete this line
	if in.Type == 1 {
		postCommentThumbResp, err := l.svcCtx.PostCommentThumbModel.FindOne(l.ctx, in.Id)
		if err != nil {
			return nil, err
		}
		postCommentThumbResp.Like = int64(in.Like)
		postCommentThumbResp.Dislike = int64(in.Dislike)
		if err = l.svcCtx.PostCommentThumbModel.UpdateWithVersion(l.ctx, nil, postCommentThumbResp); err != nil {
			return nil, err
		}

	} else if in.Type == 2 {
		// TODO: 更改文章点赞和点踩功能
	}
	return &operation.OkResp{}, nil
}
