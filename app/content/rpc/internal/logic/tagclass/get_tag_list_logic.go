package tagclasslogic

import (
	"context"
	"fmt"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"
	"qinglv-backend/pkg/sqls"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTagListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagListLogic {
	return &GetTagListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTagListLogic) GetTagList(in *content.GetTagListReq) (*content.GetTagListResp, error) {
	// todo: add your logic here and delete this line
	whereBuilder := l.svcCtx.TagModel.SelectBuilder()
	if in.Name != "" {
		whereBuilder = whereBuilder.Where("name LIKE ?", fmt.Sprint("%", in.Name, "%"))
	}
	if in.CreatorName != "" {
		whereBuilder = whereBuilder.Where("creator_name LIKE ?", fmt.Sprint("%", in.CreatorName, "%"))
	}
	orderBy := sqls.HandleSort(in.Sort)
	tagResp, total, err := l.svcCtx.TagModel.FindPageListByPageWithTotal(l.ctx, whereBuilder, int64(in.PageNum), int64(in.PageSize), orderBy)
	if err != nil {
		return nil, err
	}
	tagList := make([]*content.TagItem, len(tagResp))
	for idx, tagItem := range tagResp {
		tagList[idx] = genTagItem(tagItem)
	}
	return &content.GetTagListResp{
		Total: uint64(total),
		Data:  tagList,
	}, nil
}
