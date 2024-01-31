package postclasslogic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	postModel "qinglv-backend/app/content/rpc/internal/model/post"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostDetailLogic {
	return &GetPostDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: post
func (l *GetPostDetailLogic) GetPostDetail(in *content.GetPostDetailReq) (*content.GetPostDetailResp, error) {
	// todo: add your logic here and delete this line
	postResp, err := l.svcCtx.PostModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	postItem := genPostItem(postResp)
	return &content.GetPostDetailResp{
		Post: postItem,
	}, nil
}

func genPostItem(postItem *postModel.Post) *content.PostItem {
	return &content.PostItem{
		Id:              postItem.Id,
		CreatorId:       postItem.CreatorId,
		Status:          int32(postItem.Status),
		Visibility:      int32(postItem.Visibility),
		IsTop:           int32(postItem.IsTop),
		Location:        postItem.Location,
		Score:           postItem.Score,
		CommentCount:    postItem.CommentCount,
		LikeCount:       postItem.LikeCount,
		DislikeCount:    postItem.DislikeCount,
		ShareCount:      postItem.ShareCount,
		CollectionCount: postItem.CollectionCount,
		LastReplyTime:   uint64(postItem.LatestRepliedOn.Unix() * 1000),
		CreatedAt:       uint64(postItem.CreatedAt.Unix() * 1000),
		UpdatedAt:       uint64(postItem.UpdatedAt.Unix() * 1000),
	}
}
