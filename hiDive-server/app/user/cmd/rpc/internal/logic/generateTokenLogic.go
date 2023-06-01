package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"hiDive-server/app/user/cmd/rpc/internal/svc"
	"hiDive-server/app/user/cmd/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// rpc register(LoginReq) returns(LoginResp);
func (l *GenerateTokenLogic) GenerateToken(in *user.GenerateTokenReq) (*user.GenerateTokenResp, error) {
	//获取时间戳
	now := time.Now().Unix()
	accessSecret := l.svcCtx.Config.Auth.AccessSecret
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	token, err := l.getJwtToken(accessSecret, now, accessExpire, in.UserId)
	if err != nil {
		return nil, err
	}
	return &user.GenerateTokenResp{
		AccessToken:  token,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

// 生成token
func (l *GenerateTokenLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
