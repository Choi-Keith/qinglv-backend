package category

import (
	"context"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/content_client"
	"qinglv-backend/app/user/rpc/user_client"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type GetCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryListLogic {
	return &GetCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryListLogic) GetCategoryList(req *types.GetCategoryListReq) (resp *types.GetCategoryListResp, err error) {
	// todo: add your logic here and delete this line
	var creatorId uint64 = 0
	if req.Creator != "" {
		userResp, err := l.svcCtx.UserRpc.CheckNicknameExist(l.ctx, &user_client.CheckNicknameExistReq{
			Nickname: req.Creator,
		})
		if err != nil {
			return nil, err
		}
		if userResp.IsExist {
			creatorId = userResp.User.Id
		}
	}
	categoryListResp, err := l.svcCtx.ContentRpc.GetCategoryList(l.ctx, &content_client.GetCategoryListReq{
		CreatorId:  creatorId,
		Name:       req.Name,
		QuoteCount: req.QuoteCount,
		PageNum:    uint64(req.PageNum),
		PageSize:   uint64(req.PageSize),
	})
	if err != nil {
		return nil, err
	}
	categoryList, err := mr.MapReduce(func(source chan<- content.CategoryItem) {
		for _, categoryItem := range categoryListResp.Data {
			source <- *categoryItem
		}
	}, func(item content.CategoryItem, writer mr.Writer[types.CategoryItem], cancel func(error)) {
		userResp, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user_client.GetUserByIdReq{
			UserId: item.CreatorId,
		})
		if err != nil {
			cancel(err)
			return
		}
		var categoryItem types.CategoryItem
		_ = copier.Copy(&categoryItem, &item)
		_ = copier.Copy(&categoryItem.Creator, userResp.User)
		writer.Write(categoryItem)
	}, func(pipe <-chan types.CategoryItem, writer mr.Writer[[]types.CategoryItem], cancel func(error)) {
		var r []types.CategoryItem
		m := make(map[uint64]types.CategoryItem, len(categoryListResp.Data))
		for p := range pipe {
			m[p.Id] = p
		}
		// 为了避免mapReduce多线程导致排序的问题
		for _, categoryItem := range categoryListResp.Data {
			r = append(r, m[categoryItem.Id])
		}
		writer.Write(r)
	})
	if err != nil {
		return nil, err
	}

	isEnd := false
	total := (req.PageNum-1)*req.PageSize + req.PageSize
	if categoryListResp.Total < uint64(total) {
		isEnd = true
	}
	return &types.GetCategoryListResp{
		List:  categoryList,
		IsEnd: isEnd,
		Total: categoryListResp.Total,
	}, nil
}
