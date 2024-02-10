package tagclasslogic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTagLogic {
	return &DeleteTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteTagLogic) DeleteTag(in *content.DeleteTagReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line
	tagItem, err := l.svcCtx.TagModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	if err = l.svcCtx.TagModel.DeleteSoft(l.ctx, nil, tagItem); err != nil {
		return nil, err
	}
	return &content.OkResp{}, nil
}
