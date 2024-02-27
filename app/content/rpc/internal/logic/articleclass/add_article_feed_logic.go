package articleclasslogic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/model/article"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddArticleFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddArticleFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddArticleFeedLogic {
	return &AddArticleFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddArticleFeedLogic) AddArticleFeed(in *content.AddArticleFeedReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line
	if _, err := l.svcCtx.ArticleFeedModel.Insert(l.ctx, nil, &article.ArticleFeed{
		Id:        in.Id,
		UserId:    in.UserId,
		AuthorId:  in.AuthorId,
		ArticleId: in.ArticleId,
		IsDel:     0,
		Version:   1,
	}); err != nil {
		return nil, err
	}
	return &content.OkResp{}, nil
}
