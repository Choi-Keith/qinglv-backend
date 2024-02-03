package thumbclasslogic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentThumbDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentThumbDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentThumbDetailLogic {
	return &GetCommentThumbDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: PostCommentThumb
func (l *GetCommentThumbDetailLogic) GetCommentThumbDetail(in *operation.GetCommentThumbDetailReq) (*operation.GetCommentThumbDetailResp, error) {
	// todo: add your logic here and delete this line
	if in.Type == 1 {
		whereBuilder := l.svcCtx.PostCommentThumbModel.SelectBuilder()
		if in.CreatorId != 0 {
			whereBuilder = whereBuilder.Where(squirrel.Eq{
				"creator_id": in.CreatorId,
			})
		}
		if in.PostId != 0 {
			whereBuilder = whereBuilder.Where(squirrel.Eq{
				"post_id": in.PostId,
			})
		}
		if in.CommentId != 0 {
			whereBuilder = whereBuilder.Where(squirrel.Eq{
				"comment_id": in.CommentId,
			})
		}
		if in.ReplyId != 0 {
			whereBuilder = whereBuilder.Where(squirrel.Eq{
				"reply_id": in.ReplyId,
			})
		}
		if in.CommentType != 0 {
			whereBuilder = whereBuilder.Where(squirrel.Eq{
				"comment_type": in.CommentType,
			})
		}
		postCommentResp, err := l.svcCtx.PostCommentThumbModel.FindAll(l.ctx, whereBuilder, "")
		if err != nil {
			return nil, err
		}
		postCommentThumbList := make([]*operation.PostCommentThumbItem, len(postCommentResp))
		for idx, postThumbItem := range postCommentResp {
			postCommentThumbList[idx] = &operation.PostCommentThumbItem{
				Id:          postThumbItem.Id,
				PostId:      postThumbItem.PostId,
				CreatorId:   postThumbItem.CreatorId,
				CommentType: int32(postThumbItem.CommentType),
				CommentId:   postThumbItem.CommentId,
				ReplyId:     uint64(postThumbItem.ReplyId.Int64),
				Like:        int32(postThumbItem.Like),
				Dislike:     int32(postThumbItem.Dislike),
				CreatedAt:   uint64(postThumbItem.CreatedAt.Unix() * 1000),
				UpdatedAt:   uint64(postThumbItem.UpdatedAt.Unix() * 1000),
			}
		}
		logx.Debugf("[postCommentThumbList]: %+v\n", postCommentThumbList)
		articleCommentThumbList := make([]*operation.ArticleCommentThumbItem, 0)
		return &operation.GetCommentThumbDetailResp{
			Post:    postCommentThumbList,
			Article: articleCommentThumbList,
		}, nil

	}
	return &operation.GetCommentThumbDetailResp{}, nil
}
