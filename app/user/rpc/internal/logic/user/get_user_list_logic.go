package user

import (
	"context"
	"fmt"
	"strings"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"
	"qinglv-backend/pkg/sqlike"

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
		whereBuilder = whereBuilder.Where(squirrel.Like{"email": sqlike.GenSqlike(in.Email)})
	}
	if in.Nickname != "" {
		whereBuilder = whereBuilder.Where(squirrel.Like{"nickname": sqlike.GenSqlike(in.Nickname)})
	}
	if in.Phone != "" {
		whereBuilder = whereBuilder.Where(squirrel.Like{"phone": sqlike.GenSqlike(in.Phone)})
	}
	if in.WeChat != "" {
		whereBuilder = whereBuilder.Where(squirrel.Like{"we_chat": sqlike.GenSqlike(in.WeChat)})
	}
	sortField := "created_at"
	sort := "DESC"
	if in.Sort != "" {
		sortStrs := strings.Split(in.Sort, "|")
		sortField = sortStrs[0]
		sort = strings.ToUpper(sort)
	}
	orderBy := fmt.Sprintf("%s %s", sortField, sort)
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
