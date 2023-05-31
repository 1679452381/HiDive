package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"hiDive-server/app/user/cmd/api/internal/svc"
	"hiDive-server/app/user/cmd/api/internal/types"
	"hiDive-server/app/user/cmd/rpc/userclient"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.Request) (resp *types.Response, err error) {
	res, err := l.svcCtx.UserRpc.Ping(l.ctx, &userclient.Request{Ping: "zxc"})
	if err != nil {
		return nil, err
	}
	resp = &types.Response{Message: res.Pong}
	return
}
