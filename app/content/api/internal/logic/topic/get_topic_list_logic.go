package topic

import (
	"context"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/content_client"
	"qinglv-backend/app/user/rpc/user_client"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type GetTopicListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTopicListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicListLogic {
	return &GetTopicListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTopicListLogic) GetTopicList(req *types.GetTopicListReq) (resp *types.GetTopicListResp, err error) {
	// todo: add your logic here and delete this line
	var userId uint64 = 0
	if req.Creator != "" {
		userResp, err := l.svcCtx.UserRpc.CheckNicknameExist(l.ctx, &user_client.CheckNicknameExistReq{
			Nickname: req.Creator,
		})
		if err != nil {
			logx.Debugf("[Content] GetTopicList user failed: %+v\n", err)
			return nil, err
		}
		userId = userResp.User.Id
	}
	topicListResp, err := l.svcCtx.ContentRpc.GetTopicList(l.ctx, &content_client.GetTopicListReq{
		Name:      req.Name,
		CreatorId: userId,
		PageNum:   uint64(req.PageNum),
		PageSize:  uint64(req.PageSize),
	})
	logx.Debugf("topicListResp: %+v\n", topicListResp)

	if err != nil {
		return nil, err
	}

	newTopicList, err := mr.MapReduce(func(source chan<- content.TopicItem) {
		for _, topicItem := range topicListResp.Data {
			source <- *topicItem
		}
	}, func(item content.TopicItem, writer mr.Writer[types.TopicItem], cancel func(error)) {
		userResp, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user_client.GetUserByIdReq{
			UserId: item.CreatorId,
		})
		if err != nil {
			cancel(err)
			return
		}
		var topicItem types.TopicItem
		_ = copier.Copy(&topicItem, &item)
		_ = copier.Copy(&topicItem.Creator, userResp.User)
		writer.Write(topicItem)
	}, func(pipe <-chan types.TopicItem, writer mr.Writer[[]types.TopicItem], cancel func(error)) {
		var r []types.TopicItem
		for p := range pipe {
			r = append(r, p)
		}
		writer.Write(r)
	})
	if err != nil {
		return nil, err
	}
	isEnd := false
	total := (req.PageNum-1)*req.PageSize + req.PageSize
	if topicListResp.Total < uint64(total) {
		isEnd = true
	}
	return &types.GetTopicListResp{
		List:  newTopicList,
		IsEnd: isEnd,
		Total: topicListResp.Total,
	}, nil
}
