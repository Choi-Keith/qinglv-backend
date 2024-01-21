package logic

import (
	"context"
	"time"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/model/post"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPostLogic {
	return &AddPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: post
func (l *AddPostLogic) AddPost(in *content.AddPostReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.PostModel.Insert(l.ctx, nil, &post.Post{
		Id:              in.Id,
		CreatorId:       in.CreatorId,
		Status:          uint64(in.Status),
		Visibility:      1,
		IsTop:           2,
		Version:         1,
		Location:        in.Location,
		DeletedAt:       time.Now(),
		LatestRepliedOn: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &content.OkResp{}, nil
}
