package comment

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostCommentReplyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostCommentReplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostCommentReplyListLogic {
	return &GetPostCommentReplyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
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
	commentReplyList := make([]types.PostCommentReplyItem, len(commentReplyListResp.Post.Data))
	for idx, commentReplyItem := range commentReplyListResp.Post.Data {
		_ = copier.Copy(&commentReplyList[idx], commentReplyItem)
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
