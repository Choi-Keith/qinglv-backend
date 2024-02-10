package tagclasslogic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	tagModel "qinglv-backend/app/content/rpc/internal/model/tag"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTagByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTagByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagByIdLogic {
	return &GetTagByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTagByIdLogic) GetTagById(in *content.GetTagByIdReq) (*content.GetTagByIdResp, error) {
	// todo: add your logic here and delete this line
	tagResp, err := l.svcCtx.TagModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	tagItem := genTagItem(tagResp)
	return &content.GetTagByIdResp{
		Tag: tagItem,
	}, nil
}

func genTagItem(item *tagModel.Tag) *content.TagItem {
	return &content.TagItem{
		Id:          item.Id,
		Name:        item.Name,
		Image:       item.Image,
		Description: item.Description.String,
		QuoteCount:  item.QuoteCount,
		CreatorId:   item.CreatorId,
		CreatedAt:   uint64(item.CreatedAt.Unix() * 1000),
		UpdatedAt:   uint64(item.UpdatedAt.Unix() * 1000),
	}
}
