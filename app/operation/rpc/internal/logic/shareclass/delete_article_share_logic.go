package shareclasslogic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteArticleShareLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteArticleShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleShareLogic {
	return &DeleteArticleShareLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteArticleShareLogic) DeleteArticleShare(in *operation.DeleteArticleShareReq) (*operation.OkResp, error) {
	// todo: add your logic here and delete this line
	if err := l.svcCtx.ArticleShareModel.Delete(l.ctx, nil, in.Id); err != nil {
		return nil, err
	}
	return &operation.OkResp{}, nil
}
