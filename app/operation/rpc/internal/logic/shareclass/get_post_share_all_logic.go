package shareclasslogic

import (
	"context"
	shareModel "qinglv-backend/app/operation/rpc/internal/model/share"
	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostShareAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostShareAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostShareAllLogic {
	return &GetPostShareAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPostShareAllLogic) GetPostShareAll(in *operation.GetPostShareAllReq) (*operation.GetPostShareAllResp, error) {
	// todo: add your logic here and delete this line
	whereBuilder := l.svcCtx.PostShareModel.SelectBuilder()
	if in.PostId != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"postId_id": in.PostId,
		})
	}
	postShareResp, err := l.svcCtx.PostShareModel.FindAll(l.ctx, whereBuilder, "")
	if err != nil {
		return nil, err
	}
	postShareList := make([]*operation.PostShareItem, len(postShareResp))
	for idx, postShareItem := range postShareResp {
		postShareList[idx] = genPostShareItem(postShareItem)
	}
	return &operation.GetPostShareAllResp{}, nil
}

func genPostShareItem(item *shareModel.PostShare) *operation.PostShareItem {
	return &operation.PostShareItem{
		Id:     item.Id,
		PostId: item.PostId,
	}
}
