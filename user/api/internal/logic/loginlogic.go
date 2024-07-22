package logic

import (
	"context"

	"zjhx.com/pkg/cons"
	"zjhx.com/pkg/jwt"
	"zjhx.com/user/api/internal/svc"
	"zjhx.com/user/api/internal/types"
	"zjhx.com/user/rpc/user"

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

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	opts := jwt.TokenOptions{
		AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
		Fields: map[string]interface{}{
			"Email":    req.Email,
			"Password": req.Password,
		},
	}
	accessToken, err := jwt.BuildTokens(opts)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &types.LoginResponse{
		Code:   cons.ServiceOKCode,
		Msg:    cons.ServiceOKMsg,
		UserId: res.UserId,
		Name:   res.Name,
		Avatar: res.Avatar,
		Token:  types.Token(accessToken),
	}, nil
}
