package upload

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content_client"
	"qinglv-backend/app/user/rpc/user_client"
	"qinglv-backend/pkg/snowflake"
	"qinglv-backend/pkg/utils"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadImagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadImagesLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UploadImagesLogic {
	return &UploadImagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadImagesLogic) UploadImages(req *types.UploadImagesReq, r *http.Request) (resp *types.UploadImagesResp, err error) {
	// todo: add your logic here and delete this line
	_, fileHeaders, err := utils.GetFormFile(r, "images")
	if err != nil {
		return nil, err
	}
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	var images []string
	for _, fh := range fileHeaders {
		file, err := fh.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()
		key, err := l.handleUpload(file, *fh, req.Type, userId)
		if err != nil {
			return nil, err
		}
		id := snowflake.MustID()
		image := fmt.Sprintf("%s%s", l.svcCtx.Config.Cos.Endpoint, key)
		_, err = l.svcCtx.ContentRpc.AddMediaFile(l.ctx, &content_client.AddMediaFileReq{
			Id:        id,
			Content:   image,
			CreatorId: uint64(userId),
			FileSize:  uint64(fh.Size),
			MediaType: 1,
			BizType:   int32(req.Type),
		})
		if err != nil {
			return nil, err
		}
		images = append(images, image)
	}
	return &types.UploadImagesResp{
		Images: images,
	}, nil
}

func (l *UploadImagesLogic) handleUpload(file multipart.File, fh multipart.FileHeader, bizType int, userId int64) (string, error) {
	fileByte, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	_, filename, _, _ := runtime.Caller(0)
	filePath := fmt.Sprintf("%s/%s", filepath.Dir(filename), fh.Filename)
	err = os.WriteFile(filePath, fileByte, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close()

	userResp, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user_client.GetUserByIdReq{
		UserId: uint64(userId),
	})
	if err != nil {
		return "", err
	}
	timeStr := strconv.Itoa(int(time.Now().Unix()))
	key := fmt.Sprintf("%s%s_%s_%s", getPathByType(l.svcCtx, bizType), timeStr, userResp.User.Nickname, fh.Filename)
	text := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("轻旅 | %s", userResp.User.Nickname)))
	fill := base64.StdEncoding.EncodeToString([]byte("#9E9E9E"))
	rule := fmt.Sprintf("watermark/2/text/%s/font/dGFob21hLnR0Zg/fontsize/72/fill/%s/dissolve/50/shadow/0/gravity/center/dx/0/dy/0", text, fill)
	pic := &cos.PicOperations{
		IsPicInfo: 1,
		Rules: []cos.PicOperationsRules{
			{
				FileId: key,
				Rule:   rule,
			},
		},
	}
	opt := &cos.ObjectPutOptions{
		ACLHeaderOptions: nil,
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			XOptionHeader: &http.Header{},
		},
	}
	opt.XOptionHeader.Add("Pic-Operations", cos.EncodePicOperations(pic))
	logx.Debugf("imagesResp: %+v, %s\n", opt.ObjectPutHeaderOptions.XOptionHeader, key)
	if _, _, err = l.svcCtx.CosClient.CI.PutFromFile(context.Background(), key, filePath, nil); err != nil {
		return "", err
	}

	if err = os.Remove(filePath); err != nil {
		return "", err
	}

	return key, nil
}
