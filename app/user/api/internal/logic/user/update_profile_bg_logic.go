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
	profileBg = "profileBg"
)

type UpdateProfileBgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUpdateProfileBgLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UpdateProfileBgLogic {
	return &UpdateProfileBgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UpdateProfileBgLogic) UpdateProfileBg(req *types.UpdateProfileReq) error {
	// todo: add your logic here and delete this line
	file, header, err := l.r.FormFile(profileBg)
	if err != nil {
		return err
	}
	defer file.Close()
	timeStr := strconv.Itoa(int(time.Now().Unix()))
	key := fmt.Sprintf("%s%s_%s", l.svcCtx.Config.Cos.ProfileBgPath, timeStr, header.Filename)
	_, err = l.svcCtx.CosClient.Object.Put(context.Background(), key, file, nil)
	if err != nil {
		return err
	}
	_, err = l.svcCtx.UserRpc.UpdateProfileBg(l.ctx, &user.UpdateProfileBgReq{
		ProfileBg: fmt.Sprintf("%s%s", l.svcCtx.Config.Cos.Endpoint, key),
		UserId:    req.Id,
	})
	if err != nil {
		return err
	}
	return nil
}
