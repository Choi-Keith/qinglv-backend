package userclasslogic

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserScoreLevelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserScoreLevelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserScoreLevelLogic {
	return &UpdateUserScoreLevelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserScoreLevelLogic) UpdateUserScoreLevel(in *user.UpdateUserScoreLevelReq) (*user.UpdateUserScoreLevelResp, error) {
	// todo: add your logic here and delete this line
	userResp, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	score := userResp.Score
	if in.Op == "add" {
		score += int64(in.Score)
	} else {
		score -= int64(in.Score)
	}
	level := handleLevelByScore(score)
	userResp.Score = score
	userResp.Level = level
	if err = l.svcCtx.UserModel.UpdateWithVersion(l.ctx, nil, userResp); err != nil {
		return nil, err
	}
	return &user.UpdateUserScoreLevelResp{}, nil
}

func handleLevelByScore(score int64) int64 {
	if score >= 0 && score < 20 {
		return 1
	}
	if score >= 20 && score < 100 {
		return 2
	}
	if score >= 100 && score < 800 {
		return 3
	}
	if score >= 800 && score < 4000 {
		return 4
	}
	if score >= 4000 && score < 8000 {
		return 5
	}
	if score >= 8000 && score < 24000 {
		return 6
	}
	if score >= 24000 && score < 48000 {
		return 7
	}
	if score >= 48000 && score < 96000 {
		return 8
	}
	return 9
}
