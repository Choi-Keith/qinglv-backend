package tag

import (
	"context"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTagLogic {
	return &DeleteTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTagLogic) DeleteTag(req *types.DeleteTagReq) error {
	// todo: add your logic here and delete this line
	// userId, err := l.ctx.Value("userId").(json.Number).Int64()
	// if err != nil {
	// 	return err
	// }
	// tagResp, err := l.svcCtx.TagRpc.GetTagById(l.ctx, &content.GetTagByIdReq{
	// 	Id: req.Id,
	// })
	// if err != nil {
	// 	return err
	// }
	// roleId, _ := l.ctx.Value("roleId").(json.Number).Int64()
	// if roleId > 2 || userId != int64(tagResp.Tag.CreatorId) {
	// 	return errors.New("没有权限删除")
	// }
	// if _, err = l.svcCtx.TagRpc.DeleteTag(l.ctx, &content.DeleteTagReq{
	// 	Id: req.Id,
	// }); err != nil {
	// 	return err
	// }
	// return nil
	return nil
}
