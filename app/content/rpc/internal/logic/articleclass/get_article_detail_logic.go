package articleclasslogic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	articleModel "qinglv-backend/app/content/rpc/internal/model/article"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleDetailLogic {
	return &GetArticleDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetArticleDetailLogic) GetArticleDetail(in *content.GetArticleDetailReq) (*content.GetArticleDetailResp, error) {
	// todo: add your logic here and delete this line
	articleResp, err := l.svcCtx.ArticleModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	articleItem := genArticleItem(articleResp)
	return &content.GetArticleDetailResp{
		Article: articleItem,
	}, nil
}

func genArticleItem(item *articleModel.Article) *content.ArticleItem {
	return &content.ArticleItem{
		Id:              item.Id,
		CreatorId:       item.CreatorId,
		Visibility:      int32(item.Visibility),
		IsTop:           int32(item.IsTop),
		Score:           item.Score,
		ShareCount:      item.ShareCount,
		DislikeCount:    item.DislikeCount,
		LikeCount:       item.LikeCount,
		CommentCount:    item.CommentCount,
		CollectionCount: item.CollectionCount,
		CreatedAt:       uint64(item.CreatedAt.Unix() * 1000),
		UpdatedAt:       uint64(item.UpdatedAt.Unix() * 1000),
	}
}
