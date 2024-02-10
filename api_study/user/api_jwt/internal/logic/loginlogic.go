package logic

import (
	"context"
	"gozerotry/common/jwt"

	"gozerotry/api_study/user/api_jwt/internal/svc"
	"gozerotry/api_study/user/api_jwt/internal/types"

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
	auth := l.svcCtx.Config.Auth
	token, err := jwt.GenTokens(jwt.JwtPayload{
		UserID:   1,
		UserName: "augustus",
		Role:     1,
	}, auth.AccessSecret, auth.AccessExpire)
	if err != nil {
		return "*", err
	}
	return token, nil
}
