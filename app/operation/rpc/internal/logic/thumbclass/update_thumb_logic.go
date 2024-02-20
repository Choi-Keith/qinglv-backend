package thumbclasslogic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateThumbLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateThumbLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateThumbLogic {
	return &UpdateThumbLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: PostThumb
func (l *UpdateThumbLogic) UpdateThumb(in *operation.UpdateThumbReq) (*operation.OkResp, error) {
	// todo: add your logic here and delete this line
	if in.Type == 1 {
		thumbResp, err := l.svcCtx.PostThumbModel.FindOne(l.ctx, in.Id)
		if err != nil {
			return nil, err
		}
		thumbResp.Like = uint64(in.Like)
		thumbResp.Dislike = uint64(in.Dislike)
		err = l.svcCtx.PostThumbModel.UpdateWithVersion(l.ctx, nil, thumbResp)
		if err != nil {
			return nil, err
		}
	}
	if in.Type == 2 {
		thumbResp, err := l.svcCtx.ArticleThumbModel.FindOne(l.ctx, in.Id)
		if err != nil {
			return nil, err
		}
		thumbResp.Like = uint64(in.Like)
		thumbResp.Dislike = uint64(in.Dislike)
		err = l.svcCtx.ArticleThumbModel.UpdateWithVersion(l.ctx, nil, thumbResp)
		if err != nil {
			return nil, err
		}
	}
	return &operation.OkResp{}, nil
}
