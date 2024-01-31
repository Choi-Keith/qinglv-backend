package postclasslogic

import (
	"context"
	"fmt"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"
	"qinglv-backend/pkg/sqls"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostListLogic {
	return &GetPostListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: post
func (l *GetPostListLogic) GetPostList(in *content.GetPostListReq) (*content.GetPostListResp, error) {
	// todo: add your logic here and delete this line
	whereBuilder := l.svcCtx.PostModel.SelectBuilder()
	if in.Status != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"status": in.Status,
		})
	}
	if len(in.Visibility) != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"visibility": in.Visibility,
		})
	}
	if in.IsTop != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"is_top": in.IsTop,
		})
	}
	if in.Score != 0 {
		whereBuilder = whereBuilder.Where(squirrel.GtOrEq{
			"score": in.Score,
		})
	}
	if in.CreatorId != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"creator_id": in.CreatorId,
		})
	}
	if in.Location != "" {
		whereBuilder = whereBuilder.Where("location Like ?", fmt.Sprint("%", in.Location, "%"))
	}
	orderBy := fmt.Sprintf("%s, %s", "is_top Desc", sqls.HandleSort(in.Sort))
	postListResp, total, err := l.svcCtx.PostModel.FindPageListByPageWithTotal(l.ctx, whereBuilder, int64(in.PageNum), int64(in.PageSize), orderBy)
	if err != nil {
		return nil, err
	}
	postList := make([]*content.PostItem, len(postListResp))
	for idx, postItem := range postListResp {
		postList[idx] = genPostItem(postItem)
	}
	logx.Debugf("[Post] GetPostList postList: %+v\n", postList)
	return &content.GetPostListResp{
		Data:  postList,
		Total: uint64(total),
	}, nil
}
