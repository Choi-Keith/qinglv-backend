package postclasslogic

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/model/post"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPostContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPostContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPostContentLogic {
	return &AddPostContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: postContent
func (l *AddPostContentLogic) AddPostContent(in *content.AddPostContentReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line
	images, err := json.Marshal(in.Images)
	if err != nil {
		return nil, err
	}
	logx.Debugf("[Post] AddPostcontent images: %+v\n", images)
	_, err = l.svcCtx.PostContentModel.Insert(l.ctx, nil, &post.PostContent{
		Id:         in.Id,
		PostId:     in.PostId,
		CreatorId:  in.CreatorId,
		CategoryId: sql.NullInt64{Int64: int64(in.CategoryId), Valid: true},
		Topics:     sql.NullString{String: in.Topics, Valid: true},
		Content:    in.Content,
		Images:     sql.NullString{String: string(images), Valid: true},
		Version:    1,
		DeletedAt:  time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &content.OkResp{}, nil
}
