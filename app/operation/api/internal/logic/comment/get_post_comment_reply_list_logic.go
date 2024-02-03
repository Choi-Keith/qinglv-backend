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

type GetPostCommentReplyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetPostCommentReplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetPostCommentReplyListLogic {
	return &GetPostCommentReplyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *GetPostCommentReplyListLogic) GetPostCommentReplyList(req *types.GetPostCommentReplyListReq) (resp *types.GetPostCommentReplyListResp, err error) {
	// todo: add your logic here and delete this line
	commentReplyListResp, err := l.svcCtx.CommentRpc.GetCommentReplyList(l.ctx, &operation.GetCommentReplyListReq{
		PostId:    req.PostId,
		CommentId: req.CommentId,
		PageNum:   uint64(req.PageNum),
		PageSize:  uint64(req.PageSize),
		Type:      1,
	})
	if err != nil {
		return nil, err
	}
	commentReplyList, err := mr.MapReduce(func(source chan<- operation.PostCommentReplyItem) {
		for _, commentReplyItem := range commentReplyListResp.Post.Data {
			source <- *commentReplyItem
		}
	}, func(item operation.PostCommentReplyItem, writer mr.Writer[types.PostCommentReplyItem], cancel func(error)) {
		var commentReplyItem types.PostCommentReplyItem
		_ = copier.Copy(&commentReplyItem, &item)
		commentReplyItem.Like = globalKey.ThumbNo
		commentReplyItem.Dislike = globalKey.ThumbNo
		m, err := jwtx.ParseToken(l.r, l.svcCtx.Config.JWTAuth.AccessSecret)
		if err == nil {
			userId, ok := m["userId"]
			if ok && userId != 0 {
				postCommentReplyThumbResp, err := l.svcCtx.ThumbRpc.GetCommentThumbDetail(l.ctx, &operation.GetCommentThumbDetailReq{
					CreatorId:   userId,
					ReplyId:     commentReplyItem.Id,
					CommentId:   commentReplyItem.CommentId,
					PostId:      commentReplyItem.PostId,
					CommentType: 2,
					Type:        1,
				})
				if err != nil {
					cancel(err)
					return
				}
				if len(postCommentReplyThumbResp.Post) != 0 {
					commentReplyItem.Like = postCommentReplyThumbResp.Post[0].Like
					commentReplyItem.Dislike = postCommentReplyThumbResp.Post[0].Dislike
				}
			}
		}
		writer.Write(commentReplyItem)

	}, func(pipe <-chan types.PostCommentReplyItem, writer mr.Writer[[]types.PostCommentReplyItem], cancel func(error)) {
		var r []types.PostCommentReplyItem
		m := make(map[uint64]types.PostCommentReplyItem, len(commentReplyListResp.Post.Data))
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

	return &types.GetPostCommentReplyListResp{
		List:  commentReplyList,
		IsEnd: isEnd,
		Total: commentReplyListResp.Post.Total,
	}, nil
}
