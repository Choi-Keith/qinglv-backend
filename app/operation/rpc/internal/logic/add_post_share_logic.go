package logic

import (
	"context"
	"time"

	"qinglv-backend/app/operation/rpc/internal/model/share"
	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPostShareLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPostShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPostShareLogic {
	return &AddPostShareLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: Share
func (l *AddPostShareLogic) AddPostShare(in *operation.AddPostShareReq) (*operation.OkResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.PostShareModel.Insert(l.ctx, nil, &share.PostShare{
		Id:        in.Id,
		CreatorId: in.CreatorId,
		PostId:    in.PostId,
		Type:      uint64(in.Type),
		DeletedAt: time.Now(),
		Version:   1,
	})
	if err != nil {
		return nil, err
	}
	return &operation.OkResp{}, nil
}
