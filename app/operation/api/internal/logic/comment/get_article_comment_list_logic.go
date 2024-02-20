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

type GetArticleCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetArticleCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetArticleCommentListLogic {
	return &GetArticleCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *GetArticleCommentListLogic) GetArticleCommentList(req *types.GetArticleCommentListReq) (resp *types.GetArticleCommentListResp, err error) {
	// todo: add your logic here and delete this line
	commentListResp, err := l.svcCtx.CommentRpc.GetCommentList(l.ctx, &operation.GetCommentListReq{
		ArticleId: req.ArticleId,
		Sort:      req.Sort,
		PageNum:   uint64(req.PageNum),
		PageSize:  uint64(req.PageSize),
		Type:      2,
	})
	if err != nil {
		return nil, err
	}
	commentReplyResp, err := l.svcCtx.CommentRpc.GetCommentReplyList(l.ctx, &operation.GetCommentReplyListReq{
		ArticleId: req.ArticleId,
		Sort:      req.Sort,
		PageNum:   uint64(req.PageNum),
		PageSize:  uint64(req.PageSize),
		Type:      2,
	})
	if err != nil {
		return nil, err
	}
	commentList, err := mr.MapReduce(func(source chan<- operation.ArticleCommentItem) {
		for _, item := range commentListResp.Article.Data {
			source <- *item
		}
	}, func(item operation.ArticleCommentItem, writer mr.Writer[types.ArticleCommentItem], cancel func(error)) {
		var commentItem types.ArticleCommentItem
		_ = copier.Copy(&commentItem, &item)
		commentItem.Like = globalKey.ThumbNo
		commentItem.Dislike = globalKey.ThumbNo
		m, err := jwtx.ParseToken(l.r, l.svcCtx.Config.JWTAuth.AccessSecret)
		if err == nil {
			userId, ok := m["userId"]
			if ok && userId != 0 {
				articleThumbResp, err := l.svcCtx.ThumbRpc.GetCommentThumbDetail(l.ctx, &operation.GetCommentThumbDetailReq{
					CreatorId:   userId,
					CommentId:   commentItem.Id,
					ArticleId:   commentItem.ArticleId,
					CommentType: 1,
					Type:        2,
				})
				if err != nil {
					cancel(err)
					return
				}
				logx.Debugf("parseToken post: %+v\n", articleThumbResp)
				logx.Debugf("parseToken dislike: %v\n", articleThumbResp.Article)

				if len(articleThumbResp.Article) != 0 {
					commentItem.Like = articleThumbResp.Article[0].Like
					commentItem.Dislike = articleThumbResp.Article[0].Dislike
				}
			}

		}
		writer.Write(commentItem)
	}, func(pipe <-chan types.ArticleCommentItem, writer mr.Writer[[]types.ArticleCommentItem], cancel func(error)) {
		var r []types.ArticleCommentItem
		m := make(map[uint64]types.ArticleCommentItem, len(commentListResp.Post.Data))
		for p := range pipe {
			m[p.Id] = p
		}
		for _, commentItem := range commentListResp.Post.Data {
			r = append(r, m[commentItem.Id])
		}
		writer.Write(r)
	})
	if err != nil {
		return nil, err
	}
	isEnd := false
	if commentListResp.Post.Total <= uint64(req.PageNum-1)*uint64(req.PageSize)+uint64(req.PageSize) {
		isEnd = true
	}

	return &types.GetArticleCommentListResp{
		List:  commentList,
		IsEnd: isEnd,
		Total: commentListResp.Post.Total,
		Count: commentReplyResp.Post.Total + commentListResp.Post.Total,
	}, nil
}
