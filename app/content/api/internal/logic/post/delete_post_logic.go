package post

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

type DeletePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePostLogic) DeletePost(req *types.DeletePostReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	roleId, _ := l.ctx.Value("roleId").(json.Number).Int64()
	postResp, err := l.svcCtx.PostRpc.GetPostDetail(l.ctx, &content.GetPostDetailReq{
		Id: req.Id,
	})
	if err != nil {
		return err
	}
	if userId != int64(postResp.Post.CreatorId) && roleId > 2 {
		return errors.New("没有权限删除")
	}
	postContentResp, err := l.svcCtx.PostRpc.GetPostContentByPostId(l.ctx, &content.GetPostContentDetailReq{
		Id: postResp.Post.Id,
	})
	if err != nil {
		return err
	}
	_, err = l.svcCtx.PostRpc.DeletePost(l.ctx, &content.DeletePostReq{
		Id:            req.Id,
		PostContentId: postContentResp.PostContent.Id,
	})
	if err != nil {
		return err
	}
	if err := l.checkAndDeleteComment(req.Id); err != nil {
		return err
	}
	if err := l.checkAndDeleteCollection(req.Id); err != nil {
		return err
	}
	if err := l.checkAndDeletePostShare(req.Id); err != nil {
		return err
	}
	event.Send(event.PostDeleteEvent{
		FollowingId: postResp.Post.CreatorId,
		PostId:      req.Id,
	})
	return nil
}

func (l *DeletePostLogic) checkAndDeleteComment(postId uint64) error {
	commentListResp, err := l.svcCtx.CommentRpc.GetCommentAll(l.ctx, &operation.GetCommentAllReq{
		PostId: postId,
		Type:   1,
	})
	if err != nil {
		return err
	}

	logx.Debugf("[checkAndDeleteComment] commentListResp.Posts: %+v\n", commentListResp.Posts)

	for _, commentItem := range commentListResp.Posts {
		commentThumbResp, err := l.svcCtx.ThumbRpc.GetCommentThumbDetail(l.ctx, &operation.GetCommentThumbDetailReq{
			CommentId: commentItem.Id,
			Type:      1,
		})
		if err != nil {
			return err
		}
		if _, err := l.svcCtx.CommentRpc.DeleteComment(l.ctx, &operation.DeleteCommentReq{
			CommentId:            commentItem.Id,
			PostId:               postId,
			Type:                 1,
			PostCommentThumbList: commentThumbResp.Post,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (l *DeletePostLogic) checkAndDeleteCollection(postId uint64) error {
	collectionListResp, err := l.svcCtx.CollectionRpc.GetCollectionAll(l.ctx, &operation.GetCollectionAllReq{
		TargetId: postId,
	})
	if err != nil {
		return err
	}
	logx.Debugf("[checkAndDeleteCollection] collectionListResp.Data: %+v\n", collectionListResp.Data)
	for _, collectionItem := range collectionListResp.Data {
		if _, err := l.svcCtx.CollectionRpc.DeleteCollection(l.ctx, &operation.DeleteCollectionReq{
			Id: collectionItem.Id,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (l *DeletePostLogic) checkAndDeletePostShare(postId uint64) error {
	postShareListResp, err := l.svcCtx.ShareRpc.GetPostShareAll(l.ctx, &operation.GetPostShareAllReq{
		PostId: postId,
	})
	if err != nil {
		return err
	}
	logx.Debugf("[checkAndDeletePostShare] postShareListResp.Data: %+v\n", postShareListResp.Data)
	for _, postShareItem := range postShareListResp.Data {
		if _, err := l.svcCtx.ShareRpc.DeletePostShare(l.ctx, &operation.DeletePostShareReq{
			Id: postShareItem.Id,
		}); err != nil {
			return err
		}
	}
	return nil
}
