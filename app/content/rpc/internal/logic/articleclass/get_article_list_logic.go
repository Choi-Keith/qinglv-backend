package articleclasslogic

import (
	"context"
	"fmt"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"
	"qinglv-backend/pkg/sqls"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleListLogic {
	return &GetArticleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetArticleListLogic) GetArticleList(in *content.GetArticleListReq) (*content.GetArticleListResp, error) {
	// todo: add your logic here and delete this line
	whereBuilder := l.svcCtx.ArticleModel.SelectBuilder()
	if in.Keyword != "" {
		whereBuilder = whereBuilder.Where("title Like ?", fmt.Sprint("%", in.Keyword, "%"))
	}
	if in.Status != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"status": in.Status,
		})
	}
	if in.CreatorId != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"creator_id": in.CreatorId,
		})
	}
	if in.CreatorName != "" {
		whereBuilder = whereBuilder.Where("creator_name Like ?", fmt.Sprint("%", in.CreatorName, "%"))
	}
	if len(in.Visibility) != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"visibility": in.Visibility,
		})
	}
	if in.IsTop != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"is_top": in.IsTop,
		})
	}
	orderBy := sqls.HandleSort(in.Sort)
	articleResp, total, err := l.svcCtx.ArticleModel.FindPageListByPageWithTotal(l.ctx, whereBuilder, int64(in.PageNum), int64(in.PageSize), orderBy)
	if err != nil {
		return nil, err
	}
	articleList := make([]*content.ArticleItem, len(articleResp))
	for idx, articleItem := range articleResp {
		articleList[idx] = genArticleItem(articleItem)
	}

	return &content.GetArticleListResp{
		Data:  articleList,
		Total: uint64(total),
	}, nil
}
