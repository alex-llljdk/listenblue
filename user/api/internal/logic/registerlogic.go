package logic

import (
	"context"

	"zjhx.com/user/api/internal/svc"
	"zjhx.com/user/api/internal/types"
	"zjhx.com/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Avatar:   req.Avatar,
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	return &types.RegisterResponse{
		Code: int(res.Code),
		Msg:  res.Msg,
	}, nil
}
