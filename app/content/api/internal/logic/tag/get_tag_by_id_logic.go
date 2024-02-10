package tag

import (
	"context"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/user/rpc/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTagByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTagByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagByIdLogic {
	return &GetTagByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTagByIdLogic) GetTagById(req *types.GetTagByIdReq) (resp *types.GetTagByIdResp, err error) {
	// todo: add your logic here and delete this line
	tagResp, err := l.svcCtx.TagRpc.GetTagById(l.ctx, &content.GetTagByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	userResp, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.GetUserByIdReq{
		UserId: tagResp.Tag.CreatorId,
	})
	if err != nil {
		return nil, err
	}
	var tagItem types.TagItem
	_ = copier.Copy(&tagItem, tagResp.Tag)
	_ = copier.Copy(&tagItem.Creator, userResp.User)
	return &types.GetTagByIdResp{
		Tag: tagItem,
	}, nil
}
