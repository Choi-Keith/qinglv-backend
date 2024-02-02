package comment

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostCommentListLogic {
	return &GetPostCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
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
	commentList := make([]types.PostCommentItem, len(commentListResp.Post.Data))
	for idx, commentItem := range commentListResp.Post.Data {
		_ = copier.Copy(&commentList[idx], commentItem)
	}
	isEnd := false
	if commentListResp.Post.Total <= uint64(req.PageNum-1)*uint64(req.PageSize)+uint64(req.PageSize) {
		isEnd = true
	}

	return &types.GetPostCommentListResp{
		List:  commentList,
		IsEnd: isEnd,
		Total: commentListResp.Post.Total,
	}, nil
}
