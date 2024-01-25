package logic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMediaFileByContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMediaFileByContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMediaFileByContentLogic {
	return &GetMediaFileByContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: mediaFile
func (l *GetMediaFileByContentLogic) GetMediaFileByContent(in *content.GetMediaFileByContentReq) (*content.GetMediaFileByContentResp, error) {
	// todo: add your logic here and delete this line
	mediaFileResp, err := l.svcCtx.MediaFileModel.FindOneByContent(l.ctx, in.Content)
	if err != nil {
		return nil, err
	}

	return &content.GetMediaFileByContentResp{
		File: &content.MediaFile{
			Id:        mediaFileResp.Id,
			CreatorId: mediaFileResp.CreatorId,
			Content:   mediaFileResp.Content,
			MediaType: int32(mediaFileResp.MediaType),
			FileSize:  mediaFileResp.FileSize,
			BizType:   int32(mediaFileResp.BizType),
			CreatedAt: uint64(mediaFileResp.CreatedAt.Unix() * 1000),
			UpdatedAt: uint64(mediaFileResp.UpdatedAt.Unix() * 1000),
		},
	}, nil
}
