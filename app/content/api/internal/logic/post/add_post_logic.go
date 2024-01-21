package post

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content_client"
	"qinglv-backend/pkg/snowflake"
	"qinglv-backend/pkg/utils"

	"github.com/techxmind/ip2location"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewAddPostLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *AddPostLogic {
	return &AddPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *AddPostLogic) AddPost(req *types.AddPostReq, r *http.Request) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	if req.CategoryId != 0 {
		if _, err := l.svcCtx.ContentRpc.GetCategoryDetail(l.ctx, &content_client.GetCategoryDetailReq{
			Id: req.CategoryId,
		}); err != nil {
			return err
		}
	}
	if len(req.Topics) != 0 {
		if err := l.checkAndCreateTopic(req, uint64(userId)); err != nil {
			return err
		}
	}
	id := snowflake.MustID()
	ip := utils.GetClientIP(r)
	loc, err := ip2location.Get(ip)
	if err != nil {
		return err
	}
	logx.Debugf("[Post] AddPost loc: %+v\n", loc)
	location := loc.Province
	_, err = l.svcCtx.ContentRpc.AddPost(l.ctx, &content_client.AddPostReq{
		Id:         id,
		CreatorId:  uint64(userId),
		Location:   location,
		Status:     1,
		IsTop:      int32(req.IsTop),
		Visibility: int32(req.Visibility),
	})
	if err != nil {
		return err
	}
	topics := strings.Join(req.Topics, ",")
	contentId := snowflake.MustID()
	_, err = l.svcCtx.ContentRpc.AddPostContent(l.ctx, &content_client.AddPostContentReq{
		Id:         contentId,
		PostId:     id,
		CategoryId: req.CategoryId,
		Topics:     topics,
		CreatorId:  uint64(userId),
		Content:    req.Content,
		Images:     req.Images,
	})
	if err != nil {
		return err
	}
	return nil
}

func (l *AddPostLogic) checkAndCreateTopic(req *types.AddPostReq, userId uint64) error {
	for _, topic := range req.Topics {
		if _, err := l.svcCtx.ContentRpc.GetTopicByName(l.ctx, &content_client.GetTopicByNameReq{
			Name: topic,
		}); err != nil {
			topicId := snowflake.MustID()
			_, err := l.svcCtx.ContentRpc.AddTopic(l.ctx, &content_client.AddTopicReq{
				Id:          topicId,
				CreatorId:   uint64(userId),
				Name:        topic,
				Type:        2,
				Description: "",
				Bg:          "https://qinglv-1304086226.cos.ap-guangzhou.myqcloud.com/images/topic/default.png",
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}
