package logic

import (
	"context"
	"gozerotry/demo/apps/user/rpc/types/user"
	"gozerotry/demo/apps/video/api/internal/svc"
	"gozerotry/demo/apps/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoLogic {
	return &GetVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVideoLogic) GetVideo(req *types.VideoReq) (resp *types.VideoRes, err error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{Id: "1"})
	if err != nil {
		return nil, err
	}
	return &types.VideoRes{
		Id:   req.Id,
		Name: user.Name,
	}, nil
}
