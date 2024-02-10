package svc

import (
	"net/http"
	"net/url"
	"qinglv-backend/app/content/api/internal/config"
	"qinglv-backend/app/content/api/internal/middleware"
	"qinglv-backend/app/content/rpc/client/articleclass"
	"qinglv-backend/app/content/rpc/client/categoryclass"
	"qinglv-backend/app/content/rpc/client/mediafileclass"
	"qinglv-backend/app/content/rpc/client/postclass"
	"qinglv-backend/app/content/rpc/client/tagclass"
	"qinglv-backend/app/content/rpc/client/topicclass"
	"qinglv-backend/app/user/rpc/client/followingclass"
	"qinglv-backend/app/user/rpc/client/userclass"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	Authority    rest.Middleware
	CategoryRpc  categoryclass.CategoryClass
	MediaFileRpc mediafileclass.MediaFileClass
	PostRpc      postclass.PostClass
	TopicRpc     topicclass.TopicClass
	TagRpc       tagclass.TagClass
	ArticleRpc   articleclass.ArticleClass
	UserRpc      userclass.UserClass
	FollowingRpc followingclass.FollowingClass
	// ThumbRpc      thumbclass.ThumbClass
	// CollectionRpc collectionclass.CollectionClass
	// CommentRpc    commentclass.CommentClass
	// ShareRpc      shareclass.ShareClass
	CosClient *cos.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	u, _ := url.Parse(c.Cos.Endpoint)
	su, _ := url.Parse(c.Cos.Service)
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  c.Cos.SecretID,
			SecretKey: c.Cos.SecretKey,
		},
	})
	svc := &ServiceContext{
		Config:       c,
		CategoryRpc:  categoryclass.NewCategoryClass(zrpc.MustNewClient(c.ContentRpc)),
		MediaFileRpc: mediafileclass.NewMediaFileClass(zrpc.MustNewClient(c.ContentRpc)),
		PostRpc:      postclass.NewPostClass(zrpc.MustNewClient(c.ContentRpc)),
		TopicRpc:     topicclass.NewTopicClass(zrpc.MustNewClient(c.ContentRpc)),
		TagRpc:       tagclass.NewTagClass(zrpc.MustNewClient(c.ContentRpc)),
		ArticleRpc:   articleclass.NewArticleClass(zrpc.MustNewClient(c.ContentRpc)),
		UserRpc:      userclass.NewUserClass(zrpc.MustNewClient(c.UserRpc)),
		FollowingRpc: followingclass.NewFollowingClass(zrpc.MustNewClient(c.UserRpc)),
		CosClient:    client,
	}
	svc.Authority = middleware.NewAuthorityMiddleware(svc.UserRpc, &c).Handle
	return svc
}
