package articleclasslogic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	articleModel "qinglv-backend/app/content/rpc/internal/model/article"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleFeedListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleFeedListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleFeedListLogic {
	return &GetArticleFeedListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetArticleFeedListLogic) GetArticleFeedList(in *content.GetArticleFeedListReq) (*content.GetArticleFeedListResp, error) {
	// todo: add your logic here and delete this line
	whereBuilder := l.svcCtx.ArticleFeedModel.SelectBuilder()
	if in.UserId != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"user_id": in.UserId,
		})
	}
	if in.AuthorId != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"author_id": in.AuthorId,
		})
	}
	if in.ArticleId != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"article_id": in.ArticleId,
		})
	}
	articleFeedResp, total, err := l.svcCtx.ArticleFeedModel.FindPageListByPageWithTotal(l.ctx, whereBuilder, int64(in.PageNum), int64(in.PageSize), "")
	if err != nil {
		return nil, err
	}
	articleFeedList := make([]*content.ArticleFeedItem, len(articleFeedResp))
	for idx, articleFeedItem := range articleFeedResp {
		articleFeedList[idx] = genArticleFeedItem(articleFeedItem)
	}
	return &content.GetArticleFeedListResp{
		Data:  articleFeedList,
		Total: uint64(total),
	}, nil
}

func genArticleFeedItem(item *articleModel.ArticleFeed) *content.ArticleFeedItem {
	return &content.ArticleFeedItem{
		Id:        item.Id,
		UserId:    item.UserId,
		AuthorId:  item.AuthorId,
		ArticleId: item.ArticleId,
		CreatedAt: uint64(item.CreatedAt.Unix() * 1000),
		UpdatedAt: uint64(item.UpdatedAt.Unix() * 1000),
	}
}
