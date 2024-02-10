package logic

import (
	"context"

	"gozerotry/api_study/user/api_v2/internal/svc"
	"gozerotry/api_study/user/api_v2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoRep, err error) {
	// todo: add your logic here and delete this line

	return &types.UserInfoRep{
		UserId:   200,
		UserName: "augustus",
	}, nil
}
