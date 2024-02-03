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

type GetPostCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetPostCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetPostCommentListLogic {
	return &GetPostCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *GetPostCommentListLogic) GetPostCommentList(req *types.GetPostCommentListReq) (resp *types.GetPostCommentListResp, err error) {
	// todo: add your logic here and delete this line
	commentListResp, err := l.svcCtx.CommentRpc.GetCommentList(l.ctx, &operation.GetCommentListReq{
		PostId:   req.PostId,
		Sort:     req.Sort,
		PageNum:  uint64(req.PageNum),
		PageSize: uint64(req.PageSize),
		Type:     1,
	})
	if err != nil {
		return nil, err
	}
	commentReplyResp, err := l.svcCtx.CommentRpc.GetCommentReplyList(l.ctx, &operation.GetCommentReplyListReq{
		PostId:   req.PostId,
		Sort:     req.Sort,
		PageNum:  uint64(req.PageNum),
		PageSize: uint64(req.PageSize),
		Type:     1,
	})
	if err != nil {
		return nil, err
	}
	commentList, err := mr.MapReduce(func(source chan<- operation.PostCommentItem) {
		for _, item := range commentListResp.Post.Data {
			source <- *item
		}
	}, func(item operation.PostCommentItem, writer mr.Writer[types.PostCommentItem], cancel func(error)) {
		var commentItem types.PostCommentItem
		_ = copier.Copy(&commentItem, &item)
		commentItem.Like = globalKey.ThumbNo
		commentItem.Dislike = globalKey.ThumbNo
		m, err := jwtx.ParseToken(l.r, l.svcCtx.Config.JWTAuth.AccessSecret)
		if err == nil {
			userId, ok := m["userId"]
			if ok && userId != 0 {
				postThumbResp, err := l.svcCtx.ThumbRpc.GetCommentThumbDetail(l.ctx, &operation.GetCommentThumbDetailReq{
					CreatorId:   userId,
					CommentId:   commentItem.Id,
					PostId:      commentItem.PostId,
					CommentType: 1,
					Type:        1,
				})
				if err != nil {
					cancel(err)
					return
				}
				logx.Debugf("parseToken post: %+v\n", postThumbResp)
				logx.Debugf("parseToken dislike: %v\n", postThumbResp.Post)

				if len(postThumbResp.Post) != 0 {
					commentItem.Like = postThumbResp.Post[0].Like
					commentItem.Dislike = postThumbResp.Post[0].Dislike
				}
			}

		}
		writer.Write(commentItem)
	}, func(pipe <-chan types.PostCommentItem, writer mr.Writer[[]types.PostCommentItem], cancel func(error)) {
		var r []types.PostCommentItem
		m := make(map[uint64]types.PostCommentItem, len(commentListResp.Post.Data))
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

	return &types.GetPostCommentListResp{
		List:  commentList,
		IsEnd: isEnd,
		Total: commentListResp.Post.Total,
		Count: commentReplyResp.Post.Total + commentListResp.Post.Total,
	}, nil
}
