package commentclasslogic

import (
	"context"

	postCommentModel "qinglv-backend/app/operation/rpc/internal/model/comment"
	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"
	"qinglv-backend/pkg/sqls"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: Comment
func (l *GetCommentListLogic) GetCommentList(in *operation.GetCommentListReq) (*operation.GetCommentListResp, error) {
	// todo: add your logic here and delete this line
	if in.Type == 1 {
		whereBuilder := l.svcCtx.PostCommentModel.SelectBuilder()
		if in.PostId != 0 {
			whereBuilder = whereBuilder.Where(squirrel.Eq{
				"post_id": in.PostId,
			})
		}
		orderBy := sqls.HandleSort(in.Sort)
		commentListResp, total, err := l.svcCtx.PostCommentModel.FindPageListByPageWithTotal(l.ctx, whereBuilder, int64(in.PageNum), int64(in.PageSize), orderBy)
		if err != nil {
			return nil, err
		}
		commentList := make([]*operation.PostCommentItem, len(commentListResp))
		for idx, commentItem := range commentListResp {
			commentList[idx] = genCommentItem(commentItem)
		}
		return &operation.GetCommentListResp{
			Post: &operation.PostCommentResp{
				Data:  commentList,
				Total: uint64(total),
			},
			Article: nil,
		}, nil

	}
	return &operation.GetCommentListResp{}, nil
}

func genCommentItem(item *postCommentModel.PostComment) *operation.PostCommentItem {
	return &operation.PostCommentItem{
		Id:           item.Id,
		PostId:       item.PostId,
		CreatorId:    item.CreatorId,
		CreatorName:  item.CreatorName,
		Content:      item.Content,
		LikeCount:    item.LikeCount,
		DislikeCount: item.DislikeCount,
		Location:     item.Location,
		Score:        item.Score,
		CreatedAt:    uint64(item.CreatedAt.Unix() * 1000),
	}
}
