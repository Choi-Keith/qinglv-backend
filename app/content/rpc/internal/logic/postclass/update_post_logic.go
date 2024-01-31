package postclasslogic

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
	status := postItem.Status
	if in.Status != 0 {
		status = postItem.Status
	}
	visibility := postItem.Visibility
	if in.Visibility != 0 {
		visibility = uint64(in.Visibility)
	}
	isTop := postItem.IsTop
	if in.IsTop != 0 {
		isTop = uint64(in.IsTop)
	}
	commentCount := postItem.CommentCount
	if in.CommentCount != 0 {
		commentCount = in.CommentCount
	}
	collectionCount := postItem.CollectionCount
	if in.CollectionCount != 0 {
		collectionCount = in.CollectionCount
	}
	likeCount := postItem.LikeCount
	if in.LikeCount != 0 {
		likeCount = in.LikeCount
	}
	dislikeCount := postItem.DislikeCount
	if in.DislikeCount != 0 {
		dislikeCount = in.DislikeCount
	}
	shareCount := postItem.ShareCount
	if in.ShareCount != 0 {
		shareCount = in.ShareCount
	}
	score := postItem.Score
	if in.Score != 0 {
		score = in.Score
	}
	postItem.Status = status
	postItem.Visibility = visibility
	postItem.IsTop = isTop
	postItem.CommentCount = commentCount
	postItem.CollectionCount = collectionCount
	postItem.LikeCount = likeCount
	postItem.DislikeCount = dislikeCount
	postItem.ShareCount = shareCount
	postItem.Score = score
	err = l.svcCtx.PostModel.UpdateWithVersion(l.ctx, nil, postItem)
	if err != nil {
		return nil, err
	}
	return &content.OkResp{}, nil
}
