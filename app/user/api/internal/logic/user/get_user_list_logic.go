package user

import (
	"context"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user_client"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.UserListReq) (resp *types.UserListResp, err error) {
	// todo: add your logic here and delete this line
	userListResp, err := l.svcCtx.UserRpc.GetUserList(l.ctx, &user_client.GetUserListReq{
		Email:      req.Email,
		Nickname:   req.Nickname,
		PageNum:    int32(req.PageNum),
		PageSize:   int32(req.PageSize),
		Status:     int32(req.Status),
		MailStatus: int32(req.MailStatus),
		Sort:       req.Sort,
		WeChat:     req.WeChat,
		Phone:      req.Phone,
	})
	if err != nil {
		return nil, err
	}
	logx.Debugf("[User] GetUserList userlistresp: %+v\n", userListResp)
	userList := make([]types.User, len(userListResp.Data))
	for idx, userItem := range userListResp.Data {
		_ = copier.Copy(&userList[idx], userItem)
	}
	isEnd := false
	pageSize := uint64(req.PageSize)
	pageNum := uint64(req.PageNum)
	total := (pageNum-1)*pageSize + pageSize
	if userListResp.Total <= total {
		isEnd = true
	}
	resp = &types.UserListResp{
		List:  userList,
		Total: userListResp.Total,
		IsEnd: isEnd,
	}
	return
}
