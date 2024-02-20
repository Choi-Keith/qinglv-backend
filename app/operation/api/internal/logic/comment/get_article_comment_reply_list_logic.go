package comment

import (
	"context"
	"net/http"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"
	"qinglv-backend/app/operation/rpc/operation"
	"qinglv-backend/common/globalKey"
	"qinglv-backend/pkg/jwtx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type GetArticleCommentReplyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetArticleCommentReplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetArticleCommentReplyListLogic {
	return &GetArticleCommentReplyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *GetArticleCommentReplyListLogic) GetArticleCommentReplyList(req *types.GetArticleCommentReplyListReq) (resp *types.GetArticleCommentReplyListResp, err error) {
	// todo: add your logic here and delete this line
	commentReplyListResp, err := l.svcCtx.CommentRpc.GetCommentReplyList(l.ctx, &operation.GetCommentReplyListReq{
		ArticleId: req.ArticleId,
		CommentId: req.CommentId,
		PageNum:   uint64(req.PageNum),
		PageSize:  uint64(req.PageSize),
		Type:      2,
	})
	if err != nil {
		return nil, err
	}
	commentReplyList, err := mr.MapReduce(func(source chan<- operation.ArticleCommentReplyItem) {
		for _, commentReplyItem := range commentReplyListResp.Article.Data {
			source <- *commentReplyItem
		}
	}, func(item operation.ArticleCommentReplyItem, writer mr.Writer[types.ArticleCommentReplyItem], cancel func(error)) {
		var commentReplyItem types.ArticleCommentReplyItem
		_ = copier.Copy(&commentReplyItem, &item)
		commentReplyItem.Like = globalKey.ThumbNo
		commentReplyItem.Dislike = globalKey.ThumbNo
		m, err := jwtx.ParseToken(l.r, l.svcCtx.Config.JWTAuth.AccessSecret)
		if err == nil {
			userId, ok := m["userId"]
			if ok && userId != 0 {
				articleCommentReplyThumbResp, err := l.svcCtx.ThumbRpc.GetCommentThumbDetail(l.ctx, &operation.GetCommentThumbDetailReq{
					CreatorId:   userId,
					ReplyId:     commentReplyItem.Id,
					CommentId:   commentReplyItem.CommentId,
					ArticleId:   commentReplyItem.ArticleId,
					CommentType: 2,
					Type:        2,
				})
				if err != nil {
					cancel(err)
					return
				}
				if len(articleCommentReplyThumbResp.Article) != 0 {
					commentReplyItem.Like = articleCommentReplyThumbResp.Article[0].Like
					commentReplyItem.Dislike = articleCommentReplyThumbResp.Article[0].Dislike
				}
			}
		}
		writer.Write(commentReplyItem)

	}, func(pipe <-chan types.ArticleCommentReplyItem, writer mr.Writer[[]types.ArticleCommentReplyItem], cancel func(error)) {
		var r []types.ArticleCommentReplyItem
		m := make(map[uint64]types.ArticleCommentReplyItem, len(commentReplyListResp.Post.Data))
		for p := range pipe {
			m[p.Id] = p
		}
		for _, commentItem := range commentReplyListResp.Post.Data {
			r = append(r, m[commentItem.Id])
		}
		writer.Write(r)
	})
	if err != nil {
		return nil, err
	}
	isEnd := false
	total := (req.PageNum-1)*req.PageSize + req.PageSize
	if int32(commentReplyListResp.Post.Total) < total {
		isEnd = true
	}

	return &types.GetArticleCommentReplyListResp{
		List:  commentReplyList,
		IsEnd: isEnd,
		Total: commentReplyListResp.Post.Total,
	}, nil
}
