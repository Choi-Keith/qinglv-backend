package article

import (
	"context"
	"net/http"
	"strings"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/common/globalKey"
	"qinglv-backend/pkg/jwtx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type GetUserArticleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetUserArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetUserArticleListLogic {
	return &GetUserArticleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *GetUserArticleListLogic) GetUserArticleList(req *types.GetUserArticleListReq) (resp *types.GetUserArticleListResp, err error) {
	// todo: add your logic here and delete this line
	args := []int32{globalKey.PostVisitPublic}

	m, err := jwtx.ParseToken(l.r, l.svcCtx.Config.JWTAuth.AccessSecret)
	if err == nil {
		roleId, ok := m["roleId"]
		userId := m["userId"]
		if ok && (roleId <= 2 || userId == req.CreatorId) {
			// 如果是管理员或者自己，则都可以看见
			args = append(args, globalKey.PostVisitFriend, globalKey.PostVisitPrivate)
		}
		// TODO：如果是朋友，则需要再处理
	}

	articleListResp, err := l.svcCtx.ArticleRpc.GetArticleList(l.ctx, &content.GetArticleListReq{
		CreatorId:  req.CreatorId,
		Sort:       req.Sort,
		Visibility: args,
		PageNum:    uint64(req.PageNum),
		PageSize:   uint64(req.PageSize),
	})
	if err != nil {
		return nil, err
	}
	postList, err := mr.MapReduce(func(source chan<- content.ArticleItem) {
		for _, articleItem := range articleListResp.Data {
			source <- *articleItem
		}
	}, func(item content.ArticleItem, writer mr.Writer[types.ArticleItem], cancel func(error)) {
		articleContentListResp, err := l.svcCtx.ArticleRpc.GetArticleContentByArticleId(l.ctx, &content.GetArticleContentDetailReq{
			Id: item.Id,
		})
		if err != nil {
			cancel(err)
			return
		}
		var articleItem types.ArticleItem
		_ = copier.Copy(&articleItem, &item)
		var categoryItem types.ArticleCategory
		if articleContentListResp.ArticleContent.CategoryId != 0 {
			categoryResp, err := l.svcCtx.CategoryRpc.GetCategoryDetail(l.ctx, &content.GetCategoryDetailReq{
				Id: articleContentListResp.ArticleContent.CategoryId,
			})
			if err != nil {
				cancel(err)
				return
			}
			_ = copier.Copy(&categoryItem, categoryResp.Category)
		}
		articleItem.Category = categoryItem
		topics := strings.Split(articleContentListResp.ArticleContent.Tags, ",")
		articleItem.Tag = topics
		articleItem.Content = articleContentListResp.ArticleContent.Content
		articleItem.CoverImage = articleContentListResp.ArticleContent.CoverImage
		writer.Write(articleItem)
	}, func(pipe <-chan types.ArticleItem, writer mr.Writer[[]types.ArticleItem], cancel func(error)) {
		var r []types.ArticleItem
		m := make(map[uint64]types.ArticleItem, len(articleListResp.Data))
		for p := range pipe {
			m[p.Id] = p
		}
		for _, postItem := range articleListResp.Data {
			r = append(r, m[postItem.Id])
		}
		writer.Write(r)
	})
	if err != nil {
		return nil, err
	}
	isEnd := false
	total := (req.PageNum-1)*req.PageSize + req.PageSize
	if articleListResp.Total < uint64(total) {
		isEnd = true
	}
	return &types.GetUserArticleListResp{
		List:  postList,
		Total: articleListResp.Total,
		IsEnd: isEnd,
	}, nil
}
