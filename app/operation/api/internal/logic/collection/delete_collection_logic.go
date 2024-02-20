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

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCollectionLogic {
	return &DeleteCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCollectionLogic) DeleteCollection(req *types.DeleteCollectionReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	roleId, _ := l.ctx.Value("roleId").(json.Number).Int64()

	collectionResp, err := l.svcCtx.CollectionRpc.GetCollectionById(l.ctx, &operation.GetCollectionByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return err
	}
	if collectionResp.Collection.CreatorId != uint64(userId) && roleId > 2 {
		return errors.New("没有权限删除")
	}
	if _, err = l.svcCtx.CollectionRpc.DeleteCollection(l.ctx, &operation.DeleteCollectionReq{
		Id: req.Id,
	}); err != nil {
		return err
	}
	groupResp, err := l.svcCtx.CollectionRpc.GetCollectionGroupById(l.ctx, &operation.GetCollectionGroupByIdReq{
		Id: collectionResp.Collection.GroupId,
	})
	if err != nil {
		return err
	}
	var updateUserId uint64
	if groupResp.CollectionGroup.BizType == 1 {
		postResp, err := l.svcCtx.PostRpc.GetPostDetail(l.ctx, &content.GetPostDetailReq{
			Id: collectionResp.Collection.TargetId,
		})
		if err != nil {
			return err
		}
		updateUserId = postResp.Post.CreatorId
		if _, err = l.svcCtx.PostRpc.UpdatePost(l.ctx, &content.UpdatePostReq{
			Id:              collectionResp.Collection.TargetId,
			CollectionCount: postResp.Post.CollectionCount - 1,
		}); err != nil {
			return err
		}
	} else if groupResp.CollectionGroup.BizType == 2 {
		articleResp, err := l.svcCtx.ArticleRpc.GetArticleDetail(l.ctx, &content.GetArticleDetailReq{
			Id: collectionResp.Collection.TargetId,
		})
		if err != nil {
			return err
		}
		updateUserId = articleResp.Article.CreatorId
		if _, err := l.svcCtx.ArticleRpc.UpdateArticle(l.ctx, &content.UpdateArticleReq{
			Id:              collectionResp.Collection.TargetId,
			CollectionCount: articleResp.Article.CollectionCount - 1,
		}); err != nil {
			return err
		}
	}
	userScore := 1
	if groupResp.CollectionGroup.BizType == 2 {
		userScore = 2
	}
	if _, err := l.svcCtx.UserRpc.UpdateUserScoreLevel(l.ctx, &user.UpdateUserScoreLevelReq{
		Id:    updateUserId,
		Score: int32(userScore),
		Op:    "sub",
	}); err != nil {
		return err
	}
	return nil
}
