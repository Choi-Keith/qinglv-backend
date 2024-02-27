package articleclasslogic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/model/article"
	"qinglv-backend/app/content/rpc/internal/svc"
	"qinglv-backend/pkg/snowflake"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type ScanArticleByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewScanArticleByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScanArticleByUserIdLogic {
	return &ScanArticleByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ScanArticleByUserIdLogic) ScanArticleByUserId(in *content.ScanArticleByUserIdReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line
	var pageNo int64 = 0
	for {
		whereBuilder := l.svcCtx.ArticleModel.SelectBuilder().Where(squirrel.Eq{
			"creator_id": in.FollowingId,
			"status":     1,
			"visibility": 1,
		}).Limit(1000)
		list, _, _ := l.svcCtx.ArticleModel.FindPageListByPageWithTotal(l.ctx, whereBuilder, pageNo, 1000, "")
		if len(list) == 0 {
			break
		}
		pageNo += 1
		for _, articleItem := range list {
			id := snowflake.MustID()
			if _, err := l.svcCtx.ArticleFeedModel.Insert(l.ctx, nil, &article.ArticleFeed{
				Id:        id,
				UserId:    in.UserId,
				ArticleId: articleItem.Id,
				AuthorId:  articleItem.CreatorId,
				Version:   1,
				IsDel:     0,
			}); err != nil {
				logx.Errorf("[ScanArticleByUser] Insert failed:%+v\n", err)
			}
		}

	}
	return &content.OkResp{}, nil
}
