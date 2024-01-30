package svc

import (
	"net/http"
	"net/url"
	"qinglv-backend/app/content/rpc/content_client"
	"qinglv-backend/app/operation/api/internal/config"
	"qinglv-backend/app/operation/rpc/client/collectionclass"
	"qinglv-backend/app/operation/rpc/client/shareclass"
	"qinglv-backend/app/operation/rpc/client/thumbclass"
	"qinglv-backend/app/user/rpc/user_client"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	CollectionRpc collectionclass.CollectionClass
	ShareRpc      shareclass.ShareClass
	ThumbRpc      thumbclass.ThumbClass
	ContentRpc    content_client.Content
	UserRpc       user_client.User
	CosClient     *cos.Client
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
		Config:        c,
		CollectionRpc: collectionclass.NewCollectionClass(zrpc.MustNewClient(c.OperationRpc)),
		ShareRpc:      shareclass.NewShareClass(zrpc.MustNewClient(c.OperationRpc)),
		ThumbRpc:      thumbclass.NewThumbClass(zrpc.MustNewClient(c.OperationRpc)),
		ContentRpc:    content_client.NewContent(zrpc.MustNewClient(c.ContentRpc)),
		UserRpc:       user_client.NewUser(zrpc.MustNewClient(c.UserRpc)),
		CosClient:     client,
	}
}
