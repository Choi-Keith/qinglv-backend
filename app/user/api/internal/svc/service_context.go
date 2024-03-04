package svc

import (
	"context"
	"net/http"
	"net/url"
	"qinglv-backend/app/content/rpc/client/articleclass"
	"qinglv-backend/app/content/rpc/client/postclass"
	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/user/api/internal/config"
	"qinglv-backend/app/user/api/internal/middleware"
	"qinglv-backend/app/user/rpc/client/blacklistclass"
	"qinglv-backend/app/user/rpc/client/captchaclass"
	"qinglv-backend/app/user/rpc/client/emailclass"
	"qinglv-backend/app/user/rpc/client/followingclass"
	"qinglv-backend/app/user/rpc/client/notificationclass"
	"qinglv-backend/app/user/rpc/client/roleclass"
	"qinglv-backend/app/user/rpc/client/userclass"
	"qinglv-backend/pkg/event"
	"reflect"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	Authority       rest.Middleware
	UserRpc         userclass.UserClass
	RoleRpc         roleclass.RoleClass
	EmailRpc        emailclass.EmailClass
	CaptchaRpc      captchaclass.CaptchaClass
	FollowingRpc    followingclass.FollowingClass
	BlacklistRpc    blacklistclass.BlacklistClass
	NotificationRpc notificationclass.NotificationClass
	PostRpc         postclass.PostClass
	ArticleRpc      articleclass.ArticleClass
	CosClient       *cos.Client
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
		Config:          c,
		UserRpc:         userclass.NewUserClass(zrpc.MustNewClient(c.UserRpc)),
		RoleRpc:         roleclass.NewRoleClass(zrpc.MustNewClient(c.UserRpc)),
		EmailRpc:        emailclass.NewEmailClass(zrpc.MustNewClient(c.UserRpc)),
		CaptchaRpc:      captchaclass.NewCaptchaClass(zrpc.MustNewClient(c.UserRpc)),
		FollowingRpc:    followingclass.NewFollowingClass(zrpc.MustNewClient(c.UserRpc)),
		BlacklistRpc:    blacklistclass.NewBlacklistClass(zrpc.MustNewClient(c.UserRpc)),
		NotificationRpc: notificationclass.NewNotificationClass(zrpc.MustNewClient(c.UserRpc)),
		PostRpc:         postclass.NewPostClass(zrpc.MustNewClient(c.ContentRpc)),
		ArticleRpc:      articleclass.NewArticleClass(zrpc.MustNewClient(c.ContentRpc)),
		CosClient:       client,
	}
	svc.Authority = middleware.NewAuthorityMiddleware(svc.UserRpc, &c).Handle
	event.RegHandler(reflect.TypeOf(event.FollowEvent{}), func(i interface{}) {
		e := i.(event.FollowEvent)
		logx.Debugf("[handleFollowEvent] e: %+v\n", e)
		svc.PostRpc.ScanPostByUser(context.Background(), &content.ScanPostByUserReq{
			UserId:      e.UserId,
			FollowingId: e.OtherId,
		})
		svc.ArticleRpc.ScanArticleByUserId(context.Background(), &content.ScanArticleByUserIdReq{
			UserId:      e.UserId,
			FollowingId: e.OtherId,
		})
	})
	event.RegHandler(reflect.TypeOf(event.UnFollowEvent{}), func(i interface{}) {
		e := i.(event.UnFollowEvent)
		logx.Debugf("[handleUnFollowEvent] e: %+v\n", e)
		svc.PostRpc.DeletePostFeedByIds(context.Background(), &content.DeletePostFeedByIdsReq{
			UserId:   e.UserId,
			AuthorId: e.OtherId,
		})
		svc.ArticleRpc.DeleteArticleFeedByIds(context.Background(), &content.DeleteArticleFeedByIdsReq{
			UserId:   e.UserId,
			AuthorId: e.OtherId,
		})
	})
	return svc
}
