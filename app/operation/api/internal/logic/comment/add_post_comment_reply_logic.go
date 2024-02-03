package comment

import (
	"context"
	"encoding/json"
	"net/http"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"
	"qinglv-backend/app/operation/rpc/operation"
	"qinglv-backend/app/user/rpc/user"
	"qinglv-backend/pkg/snowflake"
	"qinglv-backend/pkg/utils"

	"github.com/techxmind/ip2location"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddPostCommentReplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewAddPostCommentReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *AddPostCommentReplyLogic {
	return &AddPostCommentReplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *AddPostCommentReplyLogic) AddPostCommentReply(req *types.AddPostCommentReplyReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	creatorName := l.ctx.Value("nickname").(string)
	id := snowflake.MustID()
	userResp, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.GetUserByIdReq{
		UserId: req.AtUserId,
	})
	if err != nil {
		return err
	}
	ip := utils.GetClientIP(l.r)
	loc, err := ip2location.Get(ip)
	if err != nil {
		return err
	}
	if _, err = l.svcCtx.CommentRpc.AddCommentReply(l.ctx, &operation.AddCommentReplyReq{
		Id:          id,
		PostId:      req.PostId,
		CommentId:   req.CommentId,
		AtUserId:    req.AtUserId,
		AtUserName:  userResp.User.Nickname,
		CreatorId:   uint64(userId),
		CreatorName: creatorName,
		Content:     req.Content,
		Type:        1,
		Location:    loc.Province,
	}); err != nil {
		return err
	}
	return nil
}
