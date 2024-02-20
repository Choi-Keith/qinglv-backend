package shareclasslogic

import (
	"context"
	"time"

	"qinglv-backend/app/operation/rpc/internal/model/share"
	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddArticleShareLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddArticleShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddArticleShareLogic {
	return &AddArticleShareLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddArticleShareLogic) AddArticleShare(in *operation.AddArticleShareReq) (*operation.OkResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.ArticleShareModel.Insert(l.ctx, nil, &share.ArticleShare{
		Id:        in.Id,
		CreatorId: in.CreatorId,
		ArticleId: in.ArticleId,
		Type:      uint64(in.Type),
		DeletedAt: time.Now(),
		Version:   1,
	})
	if err != nil {
		return nil, err
	}
	return &operation.OkResp{}, nil
}
