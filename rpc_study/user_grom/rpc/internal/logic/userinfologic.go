package logic

import (
	"context"
	"errors"
	"gozerotry/rpc_study/user_grom/models"

	"gozerotry/rpc_study/user_grom/rpc/internal/svc"
	"gozerotry/rpc_study/user_grom/rpc/types/user"

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
	var model models.UserModel
	err := l.svcCtx.DB.Take(&model, in.UserId).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	return &user.UserInfoRep{
		UserId:   uint32(model.ID),
		Username: model.Username,
	}, nil
}
