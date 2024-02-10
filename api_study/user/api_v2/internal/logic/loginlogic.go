package logic

import (
	"context"

	"gozerotry/api_study/user/api_v2/internal/svc"
	"gozerotry/api_study/user/api_v2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp string, err error) {
	// todo: add your logic here and delete this line

	return "ok", nil
}
