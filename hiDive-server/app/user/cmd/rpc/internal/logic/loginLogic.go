package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"hiDive-server/app/user/cmd/rpc/internal/svc"
	"hiDive-server/app/user/cmd/rpc/pb/user"
	"hiDive-server/app/user/model"
	"hiDive-server/common/errx"
	"hiDive-server/common/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	var err error
	var userId int64
	//登录 成功拿到userId
	userId, err = l.loginByUsername(in.Username, in.Password)
	if err != nil {
		return nil, err
	}

	//生成token
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	generateTokenResp, err := generateTokenLogic.GenerateToken(&user.GenerateTokenReq{UserId: userId})
	if err != nil {
		return nil, err
	}
	return &user.LoginResp{
		AccessToken:  generateTokenResp.AccessToken,
		AccessExpire: generateTokenResp.AccessExpire,
		RefreshAfter: generateTokenResp.RefreshAfter,
	}, nil
}

func (l *LoginLogic) loginByUsername(username, password string) (int64, error) {
	//根据用户名查询结果
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, username)

	if err != nil && err != model.ErrNotFound {
		return 0, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "根据账号查询失败,账号:%s,错误:%s", username, err)
	}
	if user == nil {
		return 0, errors.Wrapf(errx.NewErrMsg("用户不存在"), "username:%s", username)
	}
	fmt.Println(password)
	fmt.Println(tool.Md5ByString(password))
	fmt.Println(user.Password)
	if tool.Md5ByString(password) != user.Password {
		return 0, errors.Wrap(errx.NewErrMsg("账号或密码不正确"), "密码错误")
	}
	return user.Id, nil
}
