package articleclasslogic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleLogic {
	return &UpdateArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateArticleLogic) UpdateArticle(in *content.UpdateArticleReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line
	articleItem, err := l.svcCtx.ArticleModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	status := articleItem.Status
	if in.Status != 0 {
		status = articleItem.Status
	}
	visibility := articleItem.Visibility
	if in.Visibility != 0 {
		visibility = uint64(in.Visibility)
	}
	isTop := articleItem.IsTop
	if in.IsTop != 0 {
		isTop = uint64(in.IsTop)
	}
	commentCount := articleItem.CommentCount
	if in.CommentCount != 0 {
		commentCount = in.CommentCount
	}
	collectionCount := articleItem.CollectionCount
	if in.CollectionCount != 0 {
		collectionCount = in.CollectionCount
	}
	likeCount := articleItem.LikeCount
	if in.LikeCount != 0 {
		likeCount = in.LikeCount
	}
	dislikeCount := articleItem.DislikeCount
	if in.DislikeCount != 0 {
		dislikeCount = in.DislikeCount
	}
	shareCount := articleItem.ShareCount
	if in.ShareCount != 0 {
		shareCount = in.ShareCount
	}
	score := articleItem.Score
	if in.Score != 0 {
		score = float64(in.Score)
	}
	articleItem.Status = status
	articleItem.Visibility = visibility
	articleItem.IsTop = isTop
	articleItem.CommentCount = commentCount
	articleItem.CollectionCount = collectionCount
	articleItem.LikeCount = likeCount
	articleItem.DislikeCount = dislikeCount
	articleItem.ShareCount = shareCount
	articleItem.Score = score
	err = l.svcCtx.ArticleModel.UpdateWithVersion(l.ctx, nil, articleItem)
	if err != nil {
		return nil, err
	}
	return &content.OkResp{}, nil
}
