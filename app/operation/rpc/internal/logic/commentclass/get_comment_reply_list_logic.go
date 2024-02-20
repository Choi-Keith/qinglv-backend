package commentclasslogic

import (
	"context"

	commentReplyModel "qinglv-backend/app/operation/rpc/internal/model/comment"
	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"
	"qinglv-backend/pkg/sqls"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentReplyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentReplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentReplyListLogic {
	return &GetCommentReplyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: Comment
func (l *GetCommentReplyListLogic) GetCommentReplyList(in *operation.GetCommentReplyListReq) (*operation.GetCommentReplyListResp, error) {
	// todo: add your logic here and delete this line
	if in.Type == 1 {
		whereBuilder := l.svcCtx.PostCommentReplyModel.SelectBuilder()
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
		orderBy := sqls.HandleSort(in.Sort)
		commentReplyResp, total, err := l.svcCtx.PostCommentReplyModel.FindPageListByPageWithTotal(l.ctx, whereBuilder, int64(in.PageNum), int64(in.PageSize), orderBy)
		if err != nil {
			return nil, err
		}
		commentReplyList := make([]*operation.PostCommentReplyItem, len(commentReplyResp))
		for idx, commentReplyItem := range commentReplyResp {
			commentReplyList[idx] = genPostCommentReplyItem(commentReplyItem)
		}
		return &operation.GetCommentReplyListResp{
			Post: &operation.PostCommentReplyResp{
				Data:  commentReplyList,
				Total: uint64(total),
			},
		}, nil

	}
	if in.Type == 2 {
		whereBuilder := l.svcCtx.ArticleCommentReplyModel.SelectBuilder()
		if in.ArticleId != 0 {
			whereBuilder = whereBuilder.Where(squirrel.Eq{
				"article_id": in.PostId,
			})
		}
		if in.CommentId != 0 {
			whereBuilder = whereBuilder.Where(squirrel.Eq{
				"comment_id": in.CommentId,
			})
		}
		orderBy := sqls.HandleSort(in.Sort)
		commentReplyResp, total, err := l.svcCtx.ArticleCommentReplyModel.FindPageListByPageWithTotal(l.ctx, whereBuilder, int64(in.PageNum), int64(in.PageSize), orderBy)
		if err != nil {
			return nil, err
		}
		commentReplyList := make([]*operation.ArticleCommentReplyItem, len(commentReplyResp))
		for idx, commentReplyItem := range commentReplyResp {
			commentReplyList[idx] = genArticleCommentReplyItem(commentReplyItem)
		}
		return &operation.GetCommentReplyListResp{
			Article: &operation.ArticleCommentReplyResp{
				Data:  commentReplyList,
				Total: uint64(total),
			},
		}, nil
	}
	return &operation.GetCommentReplyListResp{}, nil
}

func genPostCommentReplyItem(item *commentReplyModel.PostCommentReply) *operation.PostCommentReplyItem {
	return &operation.PostCommentReplyItem{
		Id:           item.Id,
		PostId:       item.PostId,
		CommentId:    item.CommentId,
		CreatorId:    item.CreatorId,
		CreatorName:  item.CreatorName,
		AtUserId:     item.AtUserId,
		AtUserName:   item.AtUserName,
		Content:      item.Content,
		LikeCount:    item.LikeCount,
		DislikeCount: item.DislikeCount,
		Location:     item.Location,
		Score:        item.Score,
		CreatedAt:    uint64(item.CreatedAt.Unix() * 1000),
	}
}

func genArticleCommentReplyItem(item *commentReplyModel.ArticleCommentReply) *operation.ArticleCommentReplyItem {
	return &operation.ArticleCommentReplyItem{
		Id:           item.Id,
		ArticleId:    item.ArticleId,
		CommentId:    item.CommentId,
		CreatorId:    item.CreatorId,
		CreatorName:  item.CreatorName,
		AtUserId:     item.AtUserId,
		AtUserName:   item.AtUserName,
		Content:      item.Content,
		LikeCount:    item.LikeCount,
		DislikeCount: item.DislikeCount,
		Location:     item.Location,
		Score:        item.Score,
		CreatedAt:    uint64(item.CreatedAt.Unix() * 1000),
	}
}
