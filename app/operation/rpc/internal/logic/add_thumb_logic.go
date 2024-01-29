package logic

import (
	"context"
	"time"

	"qinglv-backend/app/operation/rpc/internal/model/thumb"
	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddThumbLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddThumbLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddThumbLogic {
	return &AddThumbLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: Thumb
func (l *AddThumbLogic) AddThumb(in *operation.AddThumbReq) (*operation.OkResp, error) {
	// todo: add your logic here and delete this line
	if in.Type == 1 {
		_, err := l.svcCtx.PostThumbModel.Insert(l.ctx, nil, &thumb.PostThumb{
			Id:        in.Id,
			CreatorId: in.CreatorId,
			PostId:    in.PostId,
			Like:      uint64(in.Like),
			Dislike:   uint64(in.Dislike),
			DeletedAt: time.Now(),
			Version:   1,
		})
		if err != nil {
			return nil, err
		}
	} else if in.Type == 2 {
		// TODO: 文章点赞和点踩功能
	}
	return &operation.OkResp{}, nil
}
