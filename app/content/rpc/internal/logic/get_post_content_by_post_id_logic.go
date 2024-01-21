package logic

import (
	"context"
	"encoding/json"

	"qinglv-backend/app/content/rpc/content"
	postContentModel "qinglv-backend/app/content/rpc/internal/model/post_content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostContentByPostIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostContentByPostIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostContentByPostIdLogic {
	return &GetPostContentByPostIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: postContent
func (l *GetPostContentByPostIdLogic) GetPostContentByPostId(in *content.GetPostContentDetailReq) (*content.GetPostContentDetailResp, error) {
	// todo: add your logic here and delete this line
	postContentResp, err := l.svcCtx.PostContentModel.FindOneByPostId(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	postContentItem, err := genPostContent(postContentResp)
	if err != nil {
		return nil, err
	}

	return &content.GetPostContentDetailResp{
		PostContent: postContentItem,
	}, nil
}

func genPostContent(postContent *postContentModel.PostContent) (*content.PostContentItem, error) {
	images := make([]string, 0)
	err := json.Unmarshal([]byte(postContent.Images.String), &images)
	if err != nil {
		return nil, err
	}
	return &content.PostContentItem{
		Id:         postContent.Id,
		PostId:     postContent.PostId,
		CreatorId:  postContent.CreatorId,
		CategoryId: uint64(postContent.CategoryId.Int64),
		Topics:     postContent.Topics.String,
		Content:    postContent.Content,
		Images:     images,
		CreatedAt:  uint64(postContent.CreatedAt.Unix() * 1000),
		UpdatedAt:  uint64(postContent.UpdatedAt.Unix() * 1000),
	}, nil
}
