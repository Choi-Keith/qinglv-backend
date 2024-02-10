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

type GetTagByNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTagByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagByNameLogic {
	return &GetTagByNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTagByNameLogic) GetTagByName(req *types.GetTagByNameReq) (resp *types.GetTagByNameResp, err error) {
	// todo: add your logic here and delete this line
	tagResp, err := l.svcCtx.TagRpc.GetTagByName(l.ctx, &content.GetTagByNameReq{
		Name: req.Name,
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
	return &types.GetTagByNameResp{
		Tag: tagItem,
	}, nil
}
