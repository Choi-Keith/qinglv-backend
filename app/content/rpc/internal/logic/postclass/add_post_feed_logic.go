package postclasslogic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/model/post"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPostFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPostFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPostFeedLogic {
	return &AddPostFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddPostFeedLogic) AddPostFeed(in *content.AddPostFeedReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line
	if _, err := l.svcCtx.PostFeedModel.Insert(l.ctx, nil, &post.PostFeed{
		Id:       in.Id,
		UserId:   in.UserId,
		AuthorId: in.AuthorId,
		PostId:   in.PostId,
		IsDel:    0,
		Version:  1,
	}); err != nil {
		return nil, err
	}
	return &content.OkResp{}, nil
}
