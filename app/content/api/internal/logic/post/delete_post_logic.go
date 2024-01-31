package post

import (
	"context"
	"encoding/json"
	"errors"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePostLogic) DeletePost(req *types.DeletePostReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	roleId, _ := l.ctx.Value("roleId").(json.Number).Int64()
	postResp, err := l.svcCtx.PostRpc.GetPostDetail(l.ctx, &content.GetPostDetailReq{
		Id: req.Id,
	})
	if err != nil {
		return err
	}
	if userId != int64(postResp.Post.CreatorId) && roleId > 2 {
		return errors.New("没有权限删除")
	}
	postContentResp, err := l.svcCtx.PostRpc.GetPostContentByPostId(l.ctx, &content.GetPostContentDetailReq{
		Id: postResp.Post.Id,
	})
	_, err = l.svcCtx.PostRpc.DeletePost(l.ctx, &content.DeletePostReq{
		Id:            req.Id,
		PostContentId: postContentResp.PostContent.Id,
	})
	if err != nil {
		return err
	}
	return nil
}
