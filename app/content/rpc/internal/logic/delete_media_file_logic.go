package logic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMediaFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMediaFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMediaFileLogic {
	return &DeleteMediaFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: mediaFile
func (l *DeleteMediaFileLogic) DeleteMediaFile(in *content.DeleteMediaFileReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.MediaFileModel.Delete(l.ctx, nil, in.Id)
	if err != nil {
		return nil, err
	}
	return &content.OkResp{}, nil
}
