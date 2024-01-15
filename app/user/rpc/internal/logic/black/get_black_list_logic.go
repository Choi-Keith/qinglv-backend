package black

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetBlackListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBlackListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBlackListLogic {
	return &GetBlackListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: blacklist
func (l *GetBlackListLogic) GetBlackList(in *user.GetBlackListReq) (*user.GetBlackListResp, error) {
	// todo: add your logic here and delete this line
	whereBuilder := l.svcCtx.BlacklistModel.SelectBuilder()
	if in.BlackItemId != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"black_id": in.BlackItemId,
		})
	}
	if in.UserId != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"user_id": in.UserId,
		})
	}
	blacklistResp, total, err := l.svcCtx.BlacklistModel.FindPageListByPageWithTotal(l.ctx, whereBuilder, int64(in.PageNum), int64(in.PageSize), "")
	if err != nil {
		return nil, err
	}
	blacklist := make([]*user.BlackItem, len(blacklistResp))
	for idx, blackItem := range blacklistResp {
		blacklist[idx] = &user.BlackItem{
			Id:          blackItem.Id,
			UserId:      blackItem.UserId,
			BlackItemId: blackItem.BlackId,
			CreatedAt:   uint64(blackItem.CreatedAt.Unix() * 1000),
			UpdatedAt:   uint64(blackItem.UpdatedAt.Unix() * 1000),
		}
	}
	logx.Debugf("blackListResp: %+v\n", blacklist)
	return &user.GetBlackListResp{
		Total: uint64(total),
		Data:  blacklist,
	}, nil
}
