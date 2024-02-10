package tagclasslogic

import (
	"context"
	"database/sql"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/model/tag"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTagLogic {
	return &AddTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddTagLogic) AddTag(in *content.AddTagReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.TagModel.Insert(l.ctx, nil, &tag.Tag{
		Id:          in.Id,
		CreatorId:   in.CreatorId,
		Name:        in.Name,
		Description: sql.NullString{String: in.Description, Valid: true},
		Image:       in.Image,
		Type:        int64(in.Type),
		Version:     1,
		CreatorName: in.CreatorName,
	})
	if err != nil {
		return nil, err
	}
	return &content.OkResp{}, nil
}
