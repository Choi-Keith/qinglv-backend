package postclasslogic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostFeedByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePostFeedByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostFeedByIdsLogic {
	return &DeletePostFeedByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePostFeedByIdsLogic) DeletePostFeedByIds(in *content.DeletePostFeedByIdsReq) (*content.OkResp, error) {
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
	list, _ := l.svcCtx.PostFeedModel.FindAll(l.ctx, whereBuilder, "")
	for _, postItem := range list {
		l.svcCtx.PostFeedModel.Delete(l.ctx, nil, postItem.Id)
	}
	return &content.OkResp{}, nil
}
