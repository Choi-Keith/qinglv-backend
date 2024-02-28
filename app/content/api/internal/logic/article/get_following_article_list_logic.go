package article

import (
	"context"
	"encoding/json"
	"strings"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/user/rpc/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type GetFollowingArticleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFollowingArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowingArticleListLogic {
	return &GetFollowingArticleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFollowingArticleListLogic) GetFollowingArticleList(req *types.GetArticleListReq) (resp *types.GetArticleListResp, err error) {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	articleFeedListResp, err := l.svcCtx.ArticleRpc.GetArticleFeedList(l.ctx, &content.GetArticleFeedListReq{
		UserId: uint64(userId),
	})
	if err != nil {
		return nil, err
	}
	articleList, err := mr.MapReduce(func(source chan<- content.ArticleFeedItem) {
		for _, articleItem := range articleFeedListResp.Data {
			source <- *articleItem
		}
	}, func(item content.ArticleFeedItem, writer mr.Writer[types.ArticleItem], cancel func(error)) {
		articleResp, err := l.svcCtx.ArticleRpc.GetArticleDetail(l.ctx, &content.GetArticleDetailReq{
			Id: item.ArticleId,
		})
		if err != nil {
			cancel(err)
			return
		}
		articleContentResp, err := l.svcCtx.ArticleRpc.GetArticleContentByArticleId(l.ctx, &content.GetArticleContentDetailReq{
			Id: item.ArticleId,
		})
		if err != nil {
			cancel(err)
			return
		}
		userResp, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.GetUserByIdReq{
			UserId: item.UserId,
		})
		if err != nil {
			cancel(err)
			return
		}
		var articleItem types.ArticleItem
		_ = copier.Copy(&articleItem, &articleResp)
		_ = copier.Copy(&articleItem.Creator, userResp.User)
		articleItem.Creator.IsFollowing = true
		var categoryItem types.ArticleCategory
		if articleContentResp.ArticleContent.CategoryId != 0 {
			categoryResp, err := l.svcCtx.CategoryRpc.GetCategoryDetail(l.ctx, &content.GetCategoryDetailReq{
				Id: articleContentResp.ArticleContent.CategoryId,
			})
			if err != nil {
				cancel(err)
				return
			}
			_ = copier.Copy(&categoryItem, categoryResp.Category)
		}
		articleItem.Category = categoryItem
		topics := strings.Split(articleContentResp.ArticleContent.Tags, ",")
		articleItem.Tag = topics
		articleItem.Content = articleContentResp.ArticleContent.Content
		articleItem.CoverImage = articleContentResp.ArticleContent.CoverImage
		writer.Write(articleItem)
	}, func(pipe <-chan types.ArticleItem, writer mr.Writer[[]types.ArticleItem], cancel func(error)) {
		var r []types.ArticleItem
		m := make(map[uint64]types.ArticleItem, len(articleFeedListResp.Data))
		for p := range pipe {
			m[p.Id] = p
		}
		for _, ArticleItem := range articleFeedListResp.Data {
			r = append(r, m[ArticleItem.Id])
		}
		writer.Write(r)
	})
	if err != nil {
		return nil, err
	}
	isEnd := false
	total := (req.PageNum-1)*req.PageSize + req.PageSize
	if articleFeedListResp.Total <= uint64(total) {
		isEnd = true
	}
	return &types.GetArticleListResp{
		List:  articleList,
		Total: articleFeedListResp.Total,
		IsEnd: isEnd,
	}, nil
}
