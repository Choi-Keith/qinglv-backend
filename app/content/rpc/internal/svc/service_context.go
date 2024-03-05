package svc

import (
	"qinglv-backend/app/content/rpc/internal/config"
	"qinglv-backend/app/content/rpc/internal/model/article"
	"qinglv-backend/app/content/rpc/internal/model/category"
	"qinglv-backend/app/content/rpc/internal/model/draft"
	"qinglv-backend/app/content/rpc/internal/model/media_file"
	"qinglv-backend/app/content/rpc/internal/model/post"
	"qinglv-backend/app/content/rpc/internal/model/tag"
	"qinglv-backend/app/content/rpc/internal/model/topic"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config              config.Config
	TopicModel          topic.TopicModel
	CategoryModel       category.CategoryModel
	MediaFileModel      media_file.MediaFileModel
	PostModel           post.PostModel
	PostContentModel    post.PostContentModel
	PostFeedModel       post.PostFeedModel
	TagModel            tag.TagModel
	ArticleModel        article.ArticleModel
	ArticleContentModel article.ArticleContentModel
	ArticleFeedModel    article.ArticleFeedModel
	DraftModel          draft.ArticleDraftModel
	RedisClient         *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config:              c,
		TopicModel:          topic.NewTopicModel(sqlConn, c.Cache),
		CategoryModel:       category.NewCategoryModel(sqlConn, c.Cache),
		PostModel:           post.NewPostModel(sqlConn, c.Cache),
		PostContentModel:    post.NewPostContentModel(sqlConn, c.Cache),
		PostFeedModel:       post.NewPostFeedModel(sqlConn, c.Cache),
		MediaFileModel:      media_file.NewMediaFileModel(sqlConn, c.Cache),
		TagModel:            tag.NewTagModel(sqlConn, c.Cache),
		ArticleModel:        article.NewArticleModel(sqlConn, c.Cache),
		ArticleContentModel: article.NewArticleContentModel(sqlConn, c.Cache),
		ArticleFeedModel:    article.NewArticleFeedModel(sqlConn, c.Cache),
		DraftModel:          draft.NewArticleDraftModel(sqlConn),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Pass = c.Redis.Pass
			r.Type = c.Redis.Type
		}),
	}
}
