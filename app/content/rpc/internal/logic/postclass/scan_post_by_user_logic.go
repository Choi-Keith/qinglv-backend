package postclasslogic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/model/post"
	"qinglv-backend/app/content/rpc/internal/svc"
	"qinglv-backend/pkg/snowflake"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type ScanPostByUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewScanPostByUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScanPostByUserLogic {
	return &ScanPostByUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ScanPostByUserLogic) ScanPostByUser(in *content.ScanPostByUserReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line
	var pageNo int64 = 0
	for {
		whereBuilder := l.svcCtx.PostModel.SelectBuilder().Where(squirrel.Eq{
			"creator_id": in.FollowingId,
			"status":     1,
			"visibility": 1,
		}).Limit(1000)
		list, _, _ := l.svcCtx.PostModel.FindPageListByPageWithTotal(l.ctx, whereBuilder, pageNo, 1000, "")
		if len(list) == 0 {
			break
		}
		pageNo += 1
		for _, postItem := range list {
			id := snowflake.MustID()
			if _, err := l.svcCtx.PostFeedModel.Insert(l.ctx, nil, &post.PostFeed{
				Id:       id,
				UserId:   in.UserId,
				PostId:   postItem.Id,
				AuthorId: postItem.CreatorId,
				Version:  1,
				IsDel:    0,
			}); err != nil {
				logx.Errorf("[ScanPostByUser] Insert failed:%+v\n", err)
			}
		}

	}
	return &content.OkResp{}, nil
}
