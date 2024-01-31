package userclasslogic

import (
	"context"
	"fmt"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"
	"qinglv-backend/pkg/sqls"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: user
func (l *GetUserListLogic) GetUserList(in *user.GetUserListReq) (*user.GetUserListResp, error) {
	// todo: add your logic here and delete this line
	whereBuilder := l.svcCtx.UserModel.SelectBuilder()
	if in.Status != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"status": in.Status,
		})
	}
	if in.MailStatus != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"mail_status": in.MailStatus,
		})
	}
	if in.Email != "" {
		whereBuilder = whereBuilder.Where("email LIKE ?", fmt.Sprint("%", in.Email, "%"))
	}
	if in.Nickname != "" {
		whereBuilder = whereBuilder.Where("nickname LIKE ?", fmt.Sprint("%", in.Nickname, "%"))
	}
	if in.Phone != "" {
		whereBuilder = whereBuilder.Where("phone LIKE ?", fmt.Sprint("%", in.Phone, "%"))
	}
	if in.WeChat != "" {
		whereBuilder = whereBuilder.Where("we_chat LIKE ?", fmt.Sprint("%", in.WeChat, "%"))
	}
	orderBy := sqls.HandleSort(in.Sort)
	userListResp, total, err := l.svcCtx.UserModel.FindPageListByPageWithTotal(l.ctx, whereBuilder, int64(in.PageNum), int64(in.PageSize), orderBy)
	if err != nil {
		return nil, err
	}
	userList := make([]*user.UserItem, len(userListResp))
	for idx, userItem := range userListResp {
		userList[idx] = genUserItem(userItem)
	}

	return &user.GetUserListResp{
		Data:  userList,
		Total: uint64(total),
	}, nil
}
