package mediafileclasslogic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMediaFileListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMediaFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMediaFileListLogic {
	return &GetMediaFileListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: mediaFile
func (l *GetMediaFileListLogic) GetMediaFileList(in *content.GetMediaFileListReq) (*content.GetMediaFileListResp, error) {
	// todo: add your logic here and delete this line

	return &content.GetMediaFileListResp{}, nil
}
