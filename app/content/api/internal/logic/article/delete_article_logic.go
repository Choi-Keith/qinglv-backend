package article

import (
	"context"
	"encoding/json"
	"errors"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/operation/rpc/operation"
	"qinglv-backend/pkg/event"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleLogic {
	return &DeleteArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteArticleLogic) DeleteArticle(req *types.DeleteArticleReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	roleId, _ := l.ctx.Value("roleId").(json.Number).Int64()
	articleResp, err := l.svcCtx.ArticleRpc.GetArticleDetail(l.ctx, &content.GetArticleDetailReq{
		Id: req.Id,
	})
	if err != nil {
		return err
	}
	if articleResp.Article.CreatorId != uint64(userId) && roleId > 2 {
		return errors.New("没有权限删除")
	}
	if _, err := l.svcCtx.ArticleRpc.DeleteArticle(l.ctx, &content.DeleteArticleReq{
		Id: req.Id,
	}); err != nil {
		return err
	}
	if err := l.checkAndDeleteComment(req.Id); err != nil {
		return err
	}
	if err := l.checkAndDeleteCollection(req.Id); err != nil {
		return err
	}
	if err := l.checkAndDeleteArticleShare(req.Id); err != nil {
		return err
	}
	event.Send(event.ArticleDeleteEvent{
		FollowingId: uint64(userId),
		ArticleId:   req.Id,
	})
	return nil
}

func (l *DeleteArticleLogic) checkAndDeleteComment(articleId uint64) error {
	commentListResp, err := l.svcCtx.CommentRpc.GetCommentAll(l.ctx, &operation.GetCommentAllReq{
		ArticleId: articleId,
		Type:      2,
	})
	if err != nil {
		return err
	}
	logx.Debugf("[checkAndDeleteComment] commentListResp: %+v\n", commentListResp)
	for _, commentItem := range commentListResp.Articles {
		commentThumbResp, err := l.svcCtx.ThumbRpc.GetCommentThumbDetail(l.ctx, &operation.GetCommentThumbDetailReq{
			CommentId: commentItem.Id,
			Type:      2,
		})
		if err != nil {
			return err
		}
		if _, err := l.svcCtx.CommentRpc.DeleteComment(l.ctx, &operation.DeleteCommentReq{
			CommentId:               commentItem.Id,
			ArticleId:               articleId,
			Type:                    2,
			ArticleCommentThumbList: commentThumbResp.Article,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (l *DeleteArticleLogic) checkAndDeleteCollection(articleId uint64) error {
	collectionListResp, err := l.svcCtx.CollectionRpc.GetCollectionAll(l.ctx, &operation.GetCollectionAllReq{
		TargetId: articleId,
	})
	if err != nil {
		return err
	}
	for _, collectionItem := range collectionListResp.Data {
		if _, err := l.svcCtx.CollectionRpc.DeleteCollection(l.ctx, &operation.DeleteCollectionReq{
			Id: collectionItem.Id,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (l *DeleteArticleLogic) checkAndDeleteArticleShare(articleId uint64) error {
	articleShareListResp, err := l.svcCtx.ShareRpc.GetArticleShareAll(l.ctx, &operation.GetArticleShareAllReq{
		ArticleId: articleId,
	})
	if err != nil {
		return err
	}
	for _, articleShareItem := range articleShareListResp.Data {
		if _, err := l.svcCtx.ShareRpc.DeleteArticleShare(l.ctx, &operation.DeleteArticleShareReq{
			Id: articleShareItem.Id,
		}); err != nil {
			return err
		}
	}
	return nil
}
