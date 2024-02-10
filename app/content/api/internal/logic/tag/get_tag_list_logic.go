package tag

import (
	"context"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTagListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagListLogic {
	return &GetTagListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTagListLogic) GetTagList(req *types.GetTagListReq) (resp *types.GetTagListResp, err error) {
	// todo: add your logic here and delete this line
	tagResp, err := l.svcCtx.TagRpc.GetTagList(l.ctx, &content.GetTagListReq{
		Name:       req.Name,
		Type:       int32(req.Type),
		QuoteCount: req.QuoteCount,
		Sort:       req.Sort,
		PageNum:    uint64(req.PageNum),
		PageSize:   uint64(req.PageSize),
	})
	if err != nil {
		return nil, err
	}
	tagList := make([]types.TagItem, len(tagResp.Data))
	for idx, tagItem := range tagResp.Data {
		_ = copier.Copy(&tagList[idx], tagItem)
	}
	isEnd := false
	total := (req.PageNum-1)*req.PageSize + req.PageSize
	if tagResp.Total <= uint64(total) {
		isEnd = true
	}

	return &types.GetTagListResp{
		List:  tagList,
		Total: tagResp.Total,
		IsEnd: isEnd,
	}, nil
}
