package logic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostLogic {
	return &UpdatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: post
func (l *UpdatePostLogic) UpdatePost(in *content.UpdatePostReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line
	postItem, err := l.svcCtx.PostModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	postItem.Status = uint64(in.Status)
	postItem.Visibility = uint64(in.Visibility)
	postItem.IsTop = uint64(in.IsTop)
	postItem.CommentCount = in.CommentCount
	postItem.CollectionCount = in.CollectionCount
	postItem.LikeCount = in.LikeCount
	postItem.DislikeCount = in.DislikeCount
	postItem.ShareCount = in.ShareCount
	postItem.Score = in.Score
	err = l.svcCtx.PostModel.UpdateWithVersion(l.ctx, nil, postItem)
	if err != nil {
		return nil, err
	}
	return &content.OkResp{}, nil
}
