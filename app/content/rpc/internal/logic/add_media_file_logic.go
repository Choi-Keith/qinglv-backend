package logic

import (
	"context"
	"time"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/model/media_file"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMediaFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMediaFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMediaFileLogic {
	return &AddMediaFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: mediaFile
func (l *AddMediaFileLogic) AddMediaFile(in *content.AddMediaFileReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.MediaFileModel.Insert(l.ctx, nil, &media_file.MediaFile{
		Id:        in.Id,
		CreatorId: in.CreatorId,
		Content:   in.Content,
		FileSize:  in.FileSize,
		MediaType: uint64(in.MediaType),
		BizType:   uint64(in.BizType),
		DeletedAt: time.Now(),
		Version:   1,
	})
	if err != nil {
		return nil, err
	}
	return &content.OkResp{}, nil
}
