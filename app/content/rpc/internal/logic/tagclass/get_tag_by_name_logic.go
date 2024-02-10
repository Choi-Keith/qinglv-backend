package tagclasslogic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTagByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTagByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagByNameLogic {
	return &GetTagByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTagByNameLogic) GetTagByName(in *content.GetTagByNameReq) (*content.GetTagByNameResp, error) {
	// todo: add your logic here and delete this line
	tagResp, err := l.svcCtx.TagModel.FindOneByName(l.ctx, in.Name)
	if err != nil {
		return nil, err
	}
	tagItem := genTagItem(tagResp)
	return &content.GetTagByNameResp{
		Tag: tagItem,
	}, nil
}
