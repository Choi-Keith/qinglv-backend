package upload

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveImagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveImagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveImagesLogic {
	return &RemoveImagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveImagesLogic) RemoveImages(req *types.RemoveImagesReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	roleId, _ := l.ctx.Value("roleId").(json.Number).Int64()
	if err != nil {
		return err
	}
	if len(req.Images) != 0 {
		for _, image := range req.Images {
			mediaResp, err := l.svcCtx.MediaFileRpc.GetMediaFileByContent(l.ctx, &content.GetMediaFileByContentReq{
				Content: image,
			})
			if err != nil {
				return err
			}
			if userId != int64(mediaResp.File.CreatorId) && roleId > 2 {
				return errors.New("没有权限删除")
			}
			_, err = l.svcCtx.MediaFileRpc.DeleteMediaFile(l.ctx, &content.DeleteMediaFileReq{
				Id: mediaResp.File.Id,
			})
			if err != nil {
				return nil
			}
			after, found := strings.CutPrefix(image, l.svcCtx.Config.Cos.Endpoint)
			if !found {
				return errors.New("没有找到相应的目录")
			}
			if _, err := l.svcCtx.CosClient.Object.Delete(context.Background(), after); err != nil {
				return err
			}

		}
	}
	return nil
}

// 1评论，2帖子，3文章，4个人图库
func getPathByType(svcCtx *svc.ServiceContext, bizType int) string {
	filePath := ""
	switch bizType {
	case 1:
		filePath = svcCtx.Config.Cos.CommentImagePath
	case 2:
		filePath = svcCtx.Config.Cos.PostImagePath
	case 3:
		filePath = svcCtx.Config.Cos.ArticleImagePath
	case 4:
		filePath = svcCtx.Config.Cos.GalleryImagePath
	default:
		filePath = svcCtx.Config.Cos.PostImagePath
	}
	return filePath
}
