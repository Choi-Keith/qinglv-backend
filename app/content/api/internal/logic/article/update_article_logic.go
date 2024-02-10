package article

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleLogic {
	return &UpdateArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateArticleLogic) UpdateArticle(req *types.UpdateArticleReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	roleId, _ := l.ctx.Value("roleId").(json.Number).Int64()
	if err != nil {
		return err
	}
	articleContentResp, err := l.svcCtx.ArticleRpc.GetArticleContentByArticleId(l.ctx, &content.GetArticleContentDetailReq{
		Id: req.Id,
	})
	if err != nil {
		return err
	}
	if userId != int64(articleContentResp.ArticleContent.CreatorId) && roleId > 2 {
		return errors.New("没有权限删除")
	}
	if req.Visibility != 0 {
		if _, err := l.svcCtx.ArticleRpc.UpdateArticle(l.ctx, &content.UpdateArticleReq{
			Id:         req.Id,
			Visibility: int32(req.Visibility),
		}); err != nil {
			return err
		}
	}
	tag := strings.Join(req.Tag, ",")
	if _, err := l.svcCtx.ArticleRpc.UpdateArticleContent(l.ctx, &content.UpdateArticleContentReq{
		Id:           articleContentResp.ArticleContent.Id,
		Title:        req.Title,
		Introduction: req.Desc,
		CoverImage:   req.CoverImage,
		Content:      req.Content,
		CategoryId:   req.CategoryId,
		Tags:         tag,
	}); err != nil {
		return err
	}
	return nil
}
