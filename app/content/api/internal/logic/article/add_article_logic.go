package article

import (
	"context"
	"encoding/json"
	"strings"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/user/rpc/user"
	"qinglv-backend/pkg/event"
	"qinglv-backend/pkg/snowflake"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddArticleLogic {
	return &AddArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddArticleLogic) AddArticle(req *types.AddArticleReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	creatorName := l.ctx.Value("nickname").(string)
	if req.CategoryId != 0 {
		if _, err := l.svcCtx.CategoryRpc.GetCategoryDetail(l.ctx, &content.GetCategoryDetailReq{
			Id: req.CategoryId,
		}); err != nil {
			return err
		}
	}
	var tagList []*content.TagItem
	if len(req.Tag) > 0 {
		for _, tag := range req.Tag {
			tagResp, err := l.svcCtx.TagRpc.GetTagByName(l.ctx, &content.GetTagByNameReq{
				Name: tag,
			})
			if err != nil {
				return err
			}
			tagList = append(tagList, tagResp.Tag)
		}
	}
	id := snowflake.MustID()
	articleContentId := snowflake.MustID()
	tags := strings.Join(req.Tag, ",")
	if _, err := l.svcCtx.ArticleRpc.AddArticle(l.ctx, &content.AddArticleReq{
		Id:               id,
		CreatorId:        uint64(userId),
		Visibility:       int32(req.Visibility),
		ArticleContentId: articleContentId,
		Title:            req.Title,
		Introduction:     req.Introduction,
		Tags:             tags,
		CategoryId:       req.CategoryId,
		CoverImage:       req.CoverImage,
		Content:          req.Content,
		CreatorName:      creatorName,
	}); err != nil {
		return err
	}
	if _, err := l.svcCtx.UserRpc.UpdateUserScoreLevel(l.ctx, &user.UpdateUserScoreLevelReq{
		Id:    uint64(userId),
		Score: 10,
		Op:    "add",
	}); err != nil {
		return err
	}
	for _, tagItem := range tagList {
		if _, err := l.svcCtx.TagRpc.UpdateTag(l.ctx, &content.UpdateTagReq{
			Id:         tagItem.Id,
			QuoteCount: tagItem.QuoteCount + 1,
		}); err != nil {
			return err
		}
	}
	event.Send(event.ArticleAddEvent{
		FollowingId: uint64(userId),
		ArticleId:   id,
	})

	return nil
}
