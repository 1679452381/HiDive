package logic

import (
	"context"
	"hiDive-server/app/user/cmd/rpc/userclient"
	"net/http"

	"hiDive-server/app/user/cmd/api/internal/svc"
	"hiDive-server/app/user/cmd/api/internal/types"

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

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	if req.Username == "" || req.Password == "" {
		return nil, err
	}
	loginResp, err := l.svcCtx.UserRpc.Login(l.ctx, &userclient.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.LoginResp{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data: types.TokenInfo{
			AccessToken:  loginResp.AccessToken,
			AccessExpire: loginResp.AccessExpire,
			RefreshAfter: loginResp.RefreshAfter,
		},
	}
	return
}
