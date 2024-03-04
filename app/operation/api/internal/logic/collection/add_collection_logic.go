package collection

import (
	"context"
	"encoding/json"
	"errors"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"
	"qinglv-backend/app/operation/rpc/operation"
	"qinglv-backend/app/user/rpc/user"
	"qinglv-backend/pkg/snowflake"
	"qinglv-backend/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCollectionLogic {
	return &AddCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCollectionLogic) AddCollection(req *types.AddCollectionReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	groupResp, err := l.svcCtx.CollectionRpc.GetCollectionGroupById(l.ctx, &operation.GetCollectionGroupByIdReq{
		Id: req.GroupId,
	})
	if err != nil {
		return err
	}
	collectionResp, err := l.svcCtx.CollectionRpc.GetCollectionAll(l.ctx, &operation.GetCollectionAllReq{
		CreatorId: uint64(userId),
		TargetId:  req.TargetId,
	})
	if err != nil {
		return err
	}
	if len(collectionResp.Data) > 0 {
		return errors.New("收藏夹已有")
	}
	id := snowflake.MustID()
	if _, err = l.svcCtx.CollectionRpc.AddCollection(l.ctx, &operation.AddCollectionReq{
		Id:        id,
		CreatorId: uint64(userId),
		TargetId:  req.TargetId,
		GroupId:   req.GroupId,
	}); err != nil {
		return err
	}
	var updateUserId uint64
	if groupResp.CollectionGroup.BizType == 1 {
		postResp, err := l.svcCtx.PostRpc.GetPostDetail(l.ctx, &content.GetPostDetailReq{
			Id: req.TargetId,
		})
		if err != nil {
			return err
		}
		postContentResp, err := l.svcCtx.PostRpc.GetPostContentByPostId(l.ctx, &content.GetPostContentDetailReq{
			Id: postResp.Post.Id,
		})
		if err != nil {
			return err
		}
		updateUserId = postResp.Post.CreatorId
		score := utils.HandleScore(postResp.Post.CreatedAt, 2, 1.5)
		if _, err := l.svcCtx.PostRpc.UpdatePost(l.ctx, &content.UpdatePostReq{
			Id:              req.TargetId,
			CollectionCount: postResp.Post.CollectionCount + 1,
			Score:           postResp.Post.Score + score,
		}); err != nil {
			return err
		}
		go l.svcCtx.NotificationRpc.AddNotification(l.ctx, &user.AddNotificationReq{
			Id:             snowflake.MustID(),
			Type:           2,
			ActionType:     2,
			SenderUserId:   uint64(userId),
			ReceiverUserId: postResp.Post.CreatorId,
			BizType:        1,
			TargetId:       postResp.Post.Id,
			TargetTitle:    postContentResp.PostContent.Content,
		})
	} else if groupResp.CollectionGroup.BizType == 2 {
		articleResp, err := l.svcCtx.ArticleRpc.GetArticleDetail(l.ctx, &content.GetArticleDetailReq{
			Id: req.TargetId,
		})
		if err != nil {
			return err
		}
		articleContentResp, err := l.svcCtx.ArticleRpc.GetArticleContentByArticleId(l.ctx, &content.GetArticleContentDetailReq{
			Id: articleResp.Article.Id,
		})
		if err != nil {
			return err
		}
		updateUserId = articleResp.Article.CreatorId
		score := utils.HandleScore(articleResp.Article.CreatedAt, 4, 1.5)
		if _, err := l.svcCtx.ArticleRpc.UpdateArticle(l.ctx, &content.UpdateArticleReq{
			Id:              req.TargetId,
			CollectionCount: articleResp.Article.CollectionCount + 1,
			Score:           articleResp.Article.Score + score,
		}); err != nil {
			return err
		}
		go l.svcCtx.NotificationRpc.AddNotification(l.ctx, &user.AddNotificationReq{
			Id:             snowflake.MustID(),
			Type:           2,
			ActionType:     2,
			SenderUserId:   uint64(userId),
			ReceiverUserId: articleResp.Article.CreatorId,
			BizType:        2,
			TargetId:       articleResp.Article.Id,
			TargetTitle:    articleContentResp.ArticleContent.Title,
		})
	}
	userScore := 1
	if groupResp.CollectionGroup.BizType == 2 {
		userScore = 2
	}
	if _, err := l.svcCtx.UserRpc.UpdateUserScoreLevel(l.ctx, &user.UpdateUserScoreLevelReq{
		Id:    updateUserId,
		Score: int32(userScore),
		Op:    "add",
	}); err != nil {
		return err
	}

	return nil
}
