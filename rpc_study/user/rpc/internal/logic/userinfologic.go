package logic

import (
	"context"

	"gozerotry/rpc_study/user/rpc/internal/svc"
	"gozerotry/rpc_study/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoReq) (*user.UserInfoRep, error) {
	// todo: add your logic here and delete this line

	return &user.UserInfoRep{
		UserId:   in.UserId,
		Username: "augustus",
	}, nil
}
