package postclasslogic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	postModel "qinglv-backend/app/content/rpc/internal/model/post"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostFeedListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostFeedListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostFeedListLogic {
	return &GetPostFeedListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPostFeedListLogic) GetPostFeedList(in *content.GetPostFeedListReq) (*content.GetPostFeedListResp, error) {
	// todo: add your logic here and delete this line
	whereBuilder := l.svcCtx.PostFeedModel.SelectBuilder()
	if in.UserId != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"user_id": in.UserId,
		})
	}
	if in.AuthorId != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"author_id": in.AuthorId,
		})
	}
	if in.PostId != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"post_id": in.PostId,
		})
	}
	postFeedResp, total, err := l.svcCtx.PostFeedModel.FindPageListByPageWithTotal(l.ctx, whereBuilder, int64(in.PageNum), int64(in.PageSize), "")
	if err != nil {
		return nil, err
	}
	postFeedList := make([]*content.PostFeedItem, len(postFeedResp))
	for idx, postFeedItem := range postFeedResp {
		postFeedList[idx] = genPostFeedItem(postFeedItem)
	}
	return &content.GetPostFeedListResp{
		Data:  postFeedList,
		Total: uint64(total),
	}, nil
}

func genPostFeedItem(item *postModel.PostFeed) *content.PostFeedItem {
	return &content.PostFeedItem{
		Id:        item.Id,
		UserId:    item.UserId,
		AuthorId:  item.AuthorId,
		PostId:    item.PostId,
		CreatedAt: uint64(item.CreatedAt.Unix() * 1000),
		UpdatedAt: uint64(item.UpdatedAt.Unix() * 1000),
	}
}
