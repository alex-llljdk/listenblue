package logic

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"zjhx.com/pkg/cons"
	"zjhx.com/pkg/encrypt"
	"zjhx.com/user/code"
	"zjhx.com/user/rpc/internal/svc"
	"zjhx.com/user/rpc/user"

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

// 登录
func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	//查找是否拥有该用户
	res, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logx.Error(err)
			return nil, code.UserNotExistErr
		}
		logx.Error(err)
		return nil, err
	}
	fmt.Println(1)
	//判断密码是否正确
	password := encrypt.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if password != res.Password {
		l.Logger.Infof("密码错误")
		return nil, code.UserNotExistErr
	}
	fmt.Println(3)
	return &user.LoginResponse{
		Code:   cons.ServiceOKCode,
		Msg:    cons.ServiceOKMsg,
		UserId: int64(res.Id),
		Avatar: res.Avatar,
		Name:   res.Name,
	}, nil
}
