package thumbclasslogic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetThumbDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetThumbDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetThumbDetailLogic {
	return &GetThumbDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: PostThumb
func (l *GetThumbDetailLogic) GetThumbDetail(in *operation.GetThumbDetailReq) (*operation.GetThumbDetailResp, error) {
	// todo: add your logic here and delete this line
	if in.Type == 1 {
		whereBuilder := l.svcCtx.PostThumbModel.SelectBuilder()
		if in.CreatorId != 0 {
			whereBuilder = whereBuilder.Where(squirrel.Eq{
				"creator_id": in.CreatorId,
			})
		}
		if in.PostId != 0 {
			whereBuilder = whereBuilder.Where(squirrel.Eq{
				"post_id": in.PostId,
			})
		}
		postResp, err := l.svcCtx.PostThumbModel.FindAll(l.ctx, whereBuilder, "")
		if err != nil {
			return nil, err
		}
		postThumbList := make([]*operation.PostThumbItem, len(postResp))
		for idx, postThumbItem := range postResp {
			postThumbList[idx] = &operation.PostThumbItem{
				Id:        postThumbItem.Id,
				PostId:    postThumbItem.PostId,
				CreatorId: postThumbItem.CreatorId,
				Like:      int32(postThumbItem.Like),
				Dislike:   int32(postThumbItem.Dislike),
				CreatedAt: uint64(postThumbItem.CreatedAt.Unix() * 1000),
				UpdatedAt: uint64(postThumbItem.UpdatedAt.Unix() * 1000),
			}
		}
		articleThumbList := make([]*operation.ArticleThumbItem, 0)
		return &operation.GetThumbDetailResp{
			Post:    postThumbList,
			Article: articleThumbList,
		}, nil

	}
	return &operation.GetThumbDetailResp{}, nil
}
