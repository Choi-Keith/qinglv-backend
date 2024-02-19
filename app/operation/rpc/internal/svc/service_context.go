package svc

import (
	"qinglv-backend/app/operation/rpc/internal/config"
	"qinglv-backend/app/operation/rpc/internal/model/collection"
	"qinglv-backend/app/operation/rpc/internal/model/collection_group"
	"qinglv-backend/app/operation/rpc/internal/model/comment"
	"qinglv-backend/app/operation/rpc/internal/model/share"
	"qinglv-backend/app/operation/rpc/internal/model/thumb"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config                   config.Config
	CollectionModel          collection.CollectionModel
	CollectionGroupModel     collection_group.CollectionGroupModel
	PostShareModel           share.PostShareModel
	PostCommentModel         comment.PostCommentModel
	PostCommentReplyModel    comment.PostCommentReplyModel
	PostCommentThumbModel    thumb.PostCommentThumbModel
	PostThumbModel           thumb.PostThumbModel
	ArticleShareModel        share.ArticleShareModel
	ArticleCommentModel      comment.ArticleCommentModel
	ArticleCommentReplyModel comment.ArticleCommentReplyModel
	ArticleCommentThumbModel thumb.ArticleCommentThumbModel
	ArticleThumbModel        thumb.ArticleThumbModel
	RedisClient              *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config:                   c,
		CollectionModel:          collection.NewCollectionModel(sqlConn, c.Cache),
		CollectionGroupModel:     collection_group.NewCollectionGroupModel(sqlConn, c.Cache),
		PostShareModel:           share.NewPostShareModel(sqlConn, c.Cache),
		PostCommentThumbModel:    thumb.NewPostCommentThumbModel(sqlConn, c.Cache),
		PostThumbModel:           thumb.NewPostThumbModel(sqlConn, c.Cache),
		PostCommentModel:         comment.NewPostCommentModel(sqlConn, c.Cache),
		PostCommentReplyModel:    comment.NewPostCommentReplyModel(sqlConn, c.Cache),
		ArticleShareModel:        share.NewArticleShareModel(sqlConn, c.Cache),
		ArticleCommentModel:      comment.NewArticleCommentModel(sqlConn, c.Cache),
		ArticleCommentReplyModel: comment.NewArticleCommentReplyModel(sqlConn, c.Cache),
		ArticleCommentThumbModel: thumb.NewArticleCommentThumbModel(sqlConn, c.Cache),
		ArticleThumbModel:        thumb.NewArticleThumbModel(sqlConn, c.Cache),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Pass = c.Redis.Pass
			r.Type = c.Redis.Type
		}),
	}
}
