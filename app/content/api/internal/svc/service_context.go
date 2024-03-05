package svc

import (
	"context"
	"net/http"
	"net/url"
	"qinglv-backend/app/content/api/internal/config"
	"qinglv-backend/app/content/api/internal/middleware"
	"qinglv-backend/app/content/rpc/client/articleclass"
	"qinglv-backend/app/content/rpc/client/categoryclass"
	"qinglv-backend/app/content/rpc/client/draftclass"
	"qinglv-backend/app/content/rpc/client/mediafileclass"
	"qinglv-backend/app/content/rpc/client/postclass"
	"qinglv-backend/app/content/rpc/client/tagclass"
	"qinglv-backend/app/content/rpc/client/topicclass"
	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/operation/rpc/client/collectionclass"
	"qinglv-backend/app/operation/rpc/client/commentclass"
	"qinglv-backend/app/operation/rpc/client/shareclass"
	"qinglv-backend/app/operation/rpc/client/thumbclass"
	"qinglv-backend/app/user/rpc/client/followingclass"
	"qinglv-backend/app/user/rpc/client/userclass"
	"qinglv-backend/app/user/rpc/user"
	"qinglv-backend/pkg/event"
	"qinglv-backend/pkg/snowflake"
	"reflect"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	Authority     rest.Middleware
	CategoryRpc   categoryclass.CategoryClass
	MediaFileRpc  mediafileclass.MediaFileClass
	PostRpc       postclass.PostClass
	TopicRpc      topicclass.TopicClass
	TagRpc        tagclass.TagClass
	ArticleRpc    articleclass.ArticleClass
	DraftRpc      draftclass.DraftClass
	UserRpc       userclass.UserClass
	FollowingRpc  followingclass.FollowingClass
	ThumbRpc      thumbclass.ThumbClass
	CollectionRpc collectionclass.CollectionClass
	CommentRpc    commentclass.CommentClass
	ShareRpc      shareclass.ShareClass
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
	svc := &ServiceContext{
		Config:        c,
		CategoryRpc:   categoryclass.NewCategoryClass(zrpc.MustNewClient(c.ContentRpc)),
		MediaFileRpc:  mediafileclass.NewMediaFileClass(zrpc.MustNewClient(c.ContentRpc)),
		PostRpc:       postclass.NewPostClass(zrpc.MustNewClient(c.ContentRpc)),
		TopicRpc:      topicclass.NewTopicClass(zrpc.MustNewClient(c.ContentRpc)),
		TagRpc:        tagclass.NewTagClass(zrpc.MustNewClient(c.ContentRpc)),
		ArticleRpc:    articleclass.NewArticleClass(zrpc.MustNewClient(c.ContentRpc)),
		DraftRpc:      draftclass.NewDraftClass(zrpc.MustNewClient(c.ContentRpc)),
		ThumbRpc:      thumbclass.NewThumbClass(zrpc.MustNewClient(c.OperationRpc)),
		CollectionRpc: collectionclass.NewCollectionClass(zrpc.MustNewClient(c.OperationRpc)),
		ShareRpc:      shareclass.NewShareClass(zrpc.MustNewClient(c.OperationRpc)),
		CommentRpc:    commentclass.NewCommentClass(zrpc.MustNewClient(c.OperationRpc)),
		UserRpc:       userclass.NewUserClass(zrpc.MustNewClient(c.UserRpc)),
		FollowingRpc:  followingclass.NewFollowingClass(zrpc.MustNewClient(c.UserRpc)),
		CosClient:     client,
	}
	svc.Authority = middleware.NewAuthorityMiddleware(svc.UserRpc, &c).Handle
	handleAddPostEvent(svc)
	handleAddArticleEvent(svc)
	handleDeletePostEvent(svc)
	handleDeleteArticleEvent(svc)
	return svc
}

func handleAddPostEvent(svc *ServiceContext) {
	event.RegHandler(reflect.TypeOf(event.PostAddEvent{}), func(i interface{}) {
		e := i.(event.PostAddEvent)
		logx.Debugf("[handleAddPostEvent] e: %+v\n", e)
		var pageNo int64
		for {
			followingListResp, err := svc.FollowingRpc.GetFollowingList(context.Background(), &user.GetFollowingListReq{
				FollowingId: e.FollowingId,
				PageNum:     int32(pageNo),
				PageSize:    100,
			})
			if err != nil {
				logx.Errorf("[GetFollowingList] failed: %+v\n", err)
			}
			if len(followingListResp.Data) == 0 {
				break
			}
			pageNo += 1
			id := snowflake.MustID()
			for _, followItem := range followingListResp.Data {
				if _, err := svc.PostRpc.AddPostFeed(context.Background(), &content.AddPostFeedReq{
					Id:       id,
					UserId:   followItem.UserId,
					AuthorId: e.FollowingId,
					PostId:   e.PostId,
				}); err != nil {
					logx.Errorf("[AddPostFeed] failed: %+v\n", err)
				}
			}

		}
	})
}

func handleDeletePostEvent(svc *ServiceContext) {
	event.RegHandler(reflect.TypeOf(event.PostDeleteEvent{}), func(i interface{}) {
		e := i.(event.PostDeleteEvent)
		logx.Debugf("[handleDeletePostEvent] e: %+v\n", e)
		svc.PostRpc.DeletePostFeedByIds(context.Background(), &content.DeletePostFeedByIdsReq{
			PostId: e.PostId,
		})
	})
}

func handleDeleteArticleEvent(svc *ServiceContext) {
	event.RegHandler(reflect.TypeOf(event.ArticleDeleteEvent{}), func(i interface{}) {
		e := i.(event.ArticleDeleteEvent)
		logx.Debugf("[handleDeleteArticleEvent] e: %+v\n", e)
		svc.ArticleRpc.DeleteArticleFeedByIds(context.Background(), &content.DeleteArticleFeedByIdsReq{
			ArticleId: e.ArticleId,
		})
	})
}

func handleAddArticleEvent(svc *ServiceContext) {
	event.RegHandler(reflect.TypeOf(event.ArticleAddEvent{}), func(i interface{}) {
		e := i.(event.ArticleAddEvent)
		logx.Debugf("[handleAddArticleEvent] e: %+v\n", e)

		var pageNo int64
		for {
			followingListResp, err := svc.FollowingRpc.GetFollowingList(context.Background(), &user.GetFollowingListReq{
				FollowingId: e.FollowingId,
				PageNum:     int32(pageNo),
				PageSize:    100,
			})
			if err != nil {
				logx.Errorf("[GetFollowingList] failed: %+v\n", err)
			}
			if len(followingListResp.Data) == 0 {
				break
			}
			pageNo += 1
			id := snowflake.MustID()
			for _, followItem := range followingListResp.Data {
				if _, err := svc.ArticleRpc.AddArticleFeed(context.Background(), &content.AddArticleFeedReq{
					Id:        id,
					UserId:    followItem.UserId,
					AuthorId:  e.FollowingId,
					ArticleId: e.ArticleId,
				}); err != nil {
					logx.Errorf("[AddArticleFeed] failed: %+v\n", err)
				}
			}

		}
	})
}
