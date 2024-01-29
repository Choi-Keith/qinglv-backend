package collection

import (
	"context"
	"encoding/json"
	"errors"

	"qinglv-backend/app/content/rpc/content_client"
	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"
	"qinglv-backend/app/operation/rpc/operation_client"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCollectionLogic {
	return &DeleteCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCollectionLogic) DeleteCollection(req *types.DeleteCollectionReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	roleId, _ := l.ctx.Value("roleId").(json.Number).Int64()

	collectionResp, err := l.svcCtx.OperationRpc.GetCollectionById(l.ctx, &operation_client.GetCollectionByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return err
	}
	if collectionResp.Collection.CreatorId != uint64(userId) && roleId > 2 {
		return errors.New("没有权限删除")
	}
	if _, err = l.svcCtx.OperationRpc.DeleteCollection(l.ctx, &operation_client.DeleteCollectionReq{
		Id: req.Id,
	}); err != nil {
		return err
	}
	groupResp, err := l.svcCtx.OperationRpc.GetCollectionGroupById(l.ctx, &operation_client.GetCollectionGroupByIdReq{
		Id: collectionResp.Collection.GroupId,
	})
	if err != nil {
		return err
	}
	if groupResp.CollectionGroup.BizType == 1 {
		postResp, err := l.svcCtx.ContentRpc.GetPostDetail(l.ctx, &content_client.GetPostDetailReq{
			Id: collectionResp.Collection.TargetId,
		})
		if err != nil {
			return err
		}
		if _, err = l.svcCtx.ContentRpc.UpdatePost(l.ctx, &content_client.UpdatePostReq{
			Id:              collectionResp.Collection.TargetId,
			CollectionCount: postResp.Post.CollectionCount - 1,
			Score:           postResp.Post.Score - 1,
		}); err != nil {
			return err
		}
	}
	return nil
}
