package shareclasslogic

import (
	"context"

	shareModel "qinglv-backend/app/operation/rpc/internal/model/share"
	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleShareAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleShareAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleShareAllLogic {
	return &GetArticleShareAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetArticleShareAllLogic) GetArticleShareAll(in *operation.GetArticleShareAllReq) (*operation.GetArticleShareAllResp, error) {
	// todo: add your logic here and delete this line
	whereBuilder := l.svcCtx.ArticleShareModel.SelectBuilder()
	if in.ArticleId != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"article_id": in.ArticleId,
		})
	}
	articleShareResp, err := l.svcCtx.ArticleShareModel.FindAll(l.ctx, whereBuilder, "")
	if err != nil {
		return nil, err
	}
	articleShareList := make([]*operation.ArticleShareItem, len(articleShareResp))
	for idx, articleShareItem := range articleShareResp {
		articleShareList[idx] = genArticleShareItem(articleShareItem)
	}

	return &operation.GetArticleShareAllResp{
		Data: articleShareList,
	}, nil
}

func genArticleShareItem(item *shareModel.ArticleShare) *operation.ArticleShareItem {
	return &operation.ArticleShareItem{
		Id:        item.Id,
		ArticleId: item.ArticleId,
	}
}
