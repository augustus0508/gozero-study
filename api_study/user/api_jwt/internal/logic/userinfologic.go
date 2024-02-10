package logic

import (
	"context"

	"gozerotry/api_study/user/api_jwt/internal/svc"
	"gozerotry/api_study/user/api_jwt/internal/types"

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
	//user_id := l.ctx.Value("user_id").(uint)
	username := l.ctx.Value("user_name").(string)
	return &types.UserInfoRep{
		UserId:   1,
		UserName: username,
	}, nil
}
