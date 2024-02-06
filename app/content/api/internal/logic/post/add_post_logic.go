package post

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"
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
		if _, err := l.svcCtx.CategoryRpc.GetCategoryDetail(l.ctx, &content.GetCategoryDetailReq{
			Id: req.CategoryId,
		}); err != nil {
			return err
		}
	}
	var topicList []*content.TopicItem
	if len(req.Topics) != 0 {
		topicList, err = l.checkAndCreateTopic(req, uint64(userId))
		if err != nil {
			return err
		}
	}
	id := snowflake.MustID()
	ip := utils.GetClientIP(r)
	loc, err := ip2location.Get(ip)
	if err != nil {
		return err
	}
	location := loc.Province
	_, err = l.svcCtx.PostRpc.AddPost(l.ctx, &content.AddPostReq{
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
	_, err = l.svcCtx.PostRpc.AddPostContent(l.ctx, &content.AddPostContentReq{
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
	for _, topicItem := range topicList {
		score := utils.AddScore(topicItem.CreatedAt, 5, 1.5)
		if _, err := l.svcCtx.TopicRpc.UpdateTopic(l.ctx, &content.UpdateTopicReq{
			Id:         topicItem.Id,
			Score:      topicItem.Score + score,
			QuoteCount: topicItem.QuoteCount + 1,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (l *AddPostLogic) checkAndCreateTopic(req *types.AddPostReq, userId uint64) ([]*content.TopicItem, error) {
	var scores []*content.TopicItem
	for _, topic := range req.Topics {
		topicResp, err := l.svcCtx.TopicRpc.GetTopicByName(l.ctx, &content.GetTopicByNameReq{
			Name: topic,
		})
		if err != nil {
			topicId := snowflake.MustID()
			_, err := l.svcCtx.TopicRpc.AddTopic(l.ctx, &content.AddTopicReq{
				Id:          topicId,
				CreatorId:   uint64(userId),
				Name:        topic,
				Type:        2,
				Description: "",
				Bg:          "https://qinglv-1304086226.cos.ap-guangzhou.myqcloud.com/images/topic/default.png",
				Score:       10,
			})
			if err != nil {
				return scores, err
			}
			scores = append(scores, &content.TopicItem{
				Id:         topicId,
				Score:      10,
				CreatedAt:  uint64(time.Now().Unix() * 1000),
				QuoteCount: 0,
			})
		} else {
			scores = append(scores, topicResp.Topic)
		}
	}
	return scores, nil
}
