package logic

import (
	"context"
	"fmt"

	"zjhx.com/user/api/internal/svc"
	"zjhx.com/user/api/internal/types"
	"zjhx.com/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{UserId: req.UserId, ActionId: req.ActionId})
	if err != nil {
		return nil, err
	}
	fmt.Println("取得用户信息")
	return &types.UserInfoResponse{
		Code: int(res.Code),
		Msg:  res.Msg,
		UserInfo: types.User{
			Id:              res.User.Id,
			Name:            res.User.Name,
			Avatar:          *res.User.Avatar,
			BackgroundImage: *res.User.BackgroundImage,
			FollowCount:     *res.User.FollowCount,
			FollowerCount:   *res.User.FollowerCount,
			TotalFavorited:  *res.User.TotalFavorited,
			WorkCount:       *res.User.WorkCount,
			FavoriteCount:   *res.User.FavoriteCount,
			IsFollow:        res.User.IsFollow,
		},
	}, nil
}
