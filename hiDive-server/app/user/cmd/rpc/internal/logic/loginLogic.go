package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"hiDive-server/app/user/cmd/rpc/internal/svc"
	"hiDive-server/app/user/cmd/rpc/pb"
	"hiDive-server/app/user/model"

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

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	var err error
	var userId int64
	fmt.Println("rpc.................................")
	userId, err = l.loginByUsername(in.Username, in.Password)
	if err != nil {
		return nil, err
	}
	fmt.Println("rpc.................................,", userId)
	return &pb.LoginResp{
		AccessExpire: userId,
	}, nil
}

func (l *LoginLogic) loginByUsername(username, password string) (int64, error) {
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, username)

	if err != nil && err != model.ErrNotFound {
		return 0, err
	}
	if user == nil {
		return 0, errors.New("用户不存在")
	}
	if password != user.Password {
		return 0, errors.New("密码错误")
	}
	return user.Id, nil
}
