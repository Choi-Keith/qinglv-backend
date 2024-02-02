package commentclasslogic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentByIdLogic {
	return &GetCommentByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: Comment
func (l *GetCommentByIdLogic) GetCommentById(in *operation.GetCommentByIdReq) (*operation.GetCommentByIdResp, error) {
	// todo: add your logic here and delete this line
	if in.Type == 1 {
		commentResp, err := l.svcCtx.PostCommentModel.FindOne(l.ctx, in.Id)
		if err != nil {
			return nil, err
		}
		commentItem := genCommentItem(commentResp)
		return &operation.GetCommentByIdResp{
			PostComment:    commentItem,
			ArticleComment: nil,
		}, nil
	}
	return &operation.GetCommentByIdResp{}, nil
}
