package logic

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"zjhx.com/pkg/cons"
	"zjhx.com/pkg/encrypt"
	"zjhx.com/user/code"
	"zjhx.com/user/gmodel"
	"zjhx.com/user/rpc/internal/svc"
	"zjhx.com/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注册
func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// todo: add your logic here and delete this line
	// 判断邮箱是否注册
	_, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err == nil {
		return nil, code.UserExistErr
	}

	if err == gorm.ErrRecordNotFound {
		newUser := gmodel.User{
			Email:           in.Email,
			Name:            in.Name,
			Password:        encrypt.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
			Avatar:          in.Avatar,
			Signature:       "",
			BackgroundImage: "",
		}
		
		//inser user err
		err = l.svcCtx.UserModel.InsertUser(l.ctx, &newUser)
		if err != nil {
			l.Logger.Errorf("插入数据库失败")
			return nil, err
		}

		return &user.RegisterResponse{
			Code: cons.ServiceOKCode,
			Msg:  cons.ServiceOKMsg,
		}, nil
	}
	return nil, errors.New("数据库出错")
}
