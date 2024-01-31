package svc

import (
	"net/http"
	"net/url"
	"qinglv-backend/app/content/api/internal/config"
	"qinglv-backend/app/content/rpc/client/categoryclass"
	"qinglv-backend/app/content/rpc/client/mediafileclass"
	"qinglv-backend/app/content/rpc/client/postclass"
	"qinglv-backend/app/content/rpc/client/topicclass"
	"qinglv-backend/app/user/rpc/user_client"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	CategoryRpc  categoryclass.CategoryClass
	MediaFileRpc mediafileclass.MediaFileClass
	PostRpc      postclass.PostClass
	TopicRpc     topicclass.TopicClass
	UserRpc      user_client.User
	CosClient    *cos.Client
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
	return &ServiceContext{
		Config:       c,
		CategoryRpc:  categoryclass.NewCategoryClass(zrpc.MustNewClient(c.ContentRpc)),
		MediaFileRpc: mediafileclass.NewMediaFileClass(zrpc.MustNewClient(c.ContentRpc)),
		PostRpc:      postclass.NewPostClass(zrpc.MustNewClient(c.ContentRpc)),
		TopicRpc:     topicclass.NewTopicClass(zrpc.MustNewClient(c.ContentRpc)),
		UserRpc:      user_client.NewUser(zrpc.MustNewClient(c.UserRpc)),
		CosClient:    client,
	}
}
