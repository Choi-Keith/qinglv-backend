package comment

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"
	"qinglv-backend/app/operation/rpc/operation"
	"qinglv-backend/pkg/snowflake"
	"qinglv-backend/pkg/utils"

	"github.com/techxmind/ip2location"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddPostCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewAddPostCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *AddPostCommentLogic {
	return &AddPostCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *AddPostCommentLogic) AddPostComment(req *types.AddPostCommentReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	creatorName := l.ctx.Value("nickname").(string)
	fmt.Printf("userId, %d, creatorName: %s\n", userId, creatorName)
	id := snowflake.MustID()
	ip := utils.GetClientIP(l.r)
	loc, err := ip2location.Get(ip)
	if err != nil {
		return err
	}
	postResp, err := l.svcCtx.PostRpc.GetPostDetail(l.ctx, &content.GetPostDetailReq{
		Id: req.PostId,
	})
	if err != nil {
		return err
	}
	if _, err = l.svcCtx.CommentRpc.AddComment(l.ctx, &operation.AddCommentReq{
		Id:          id,
		PostId:      req.PostId,
		Content:     req.Content,
		Location:    loc.Province,
		CreatorId:   uint64(userId),
		CreatorName: creatorName,
		Type:        1,
	}); err != nil {
		return err
	}
	if _, err = l.svcCtx.PostRpc.UpdatePost(l.ctx, &content.UpdatePostReq{
		Id:           req.PostId,
		CommentCount: postResp.Post.CommentCount + 1,
		Score:        postResp.Post.Score + 1,
	}); err != nil {
		return err
	}
	return nil
}
