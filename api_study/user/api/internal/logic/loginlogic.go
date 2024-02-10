package logic

import (
	"context"
	"fmt"

	"gozerotry/api_study/user/api/internal/svc"
	"gozerotry/api_study/user/api/internal/types"

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

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRep, err error) {
	// todo: add your logic here and delete this line
	fmt.Printf(req.Password, req.Username)
	return &types.LoginRep{
		Code: 200,
		Data: "ok",
		Msg:  "ok",
	}, nil
}
