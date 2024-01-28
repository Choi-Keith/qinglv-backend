package logic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type DeletePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: post
func (l *DeletePostLogic) DeletePost(in *content.DeletePostReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.PostModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		if err := l.svcCtx.PostContentModel.Delete(ctx, session, in.PostContentId); err != nil {
			return err
		}
		if err := l.svcCtx.PostModel.Delete(ctx, session, in.Id); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &content.OkResp{}, nil
}
