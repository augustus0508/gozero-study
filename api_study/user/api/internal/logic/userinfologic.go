package logic

import (
	"context"

	"gozerotry/api_study/user/api/internal/svc"
	"gozerotry/api_study/user/api/internal/types"

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
		Code: 200,
		Data: types.UserInfo{
			UserId:   1,
			UserName: "augustus",
		},
		Msg: "",
	}, nil
}
