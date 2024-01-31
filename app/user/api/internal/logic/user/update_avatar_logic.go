package user

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	avatar = "avatar"
)

type UpdateAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUpdateAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UpdateAvatarLogic {
	return &UpdateAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UpdateAvatarLogic) UpdateAvatar(req *types.AvatarReq) error {
	// todo: add your logic here and delete this line
	file, header, err := l.r.FormFile(avatar)
	logx.Debugf("[User] UpdateAvatar file: %+v, header: %+v\n", file, header)
	if err != nil {
		return err
	}
	defer file.Close()
	timeStr := strconv.Itoa(int(time.Now().Unix()))
	key := fmt.Sprintf("%s%s_%s", l.svcCtx.Config.Cos.AvatarPath, timeStr, header.Filename)
	_, err = l.svcCtx.CosClient.Object.Put(context.Background(), key, file, nil)
	if err != nil {
		return err
	}
	_, err = l.svcCtx.UserRpc.UpdateAvatar(l.ctx, &user.UpdateAvatarReq{
		Avatar: fmt.Sprintf("%s%s", l.svcCtx.Config.Cos.Endpoint, key),
		UserId: req.Id,
	})
	if err != nil {
		return err
	}
	return nil
}
