package logic

import (
	"context"

	"gorm.io/gorm"
	"zjhx.com/pkg/cons"
	"zjhx.com/user/code"
	"zjhx.com/user/rpc/internal/svc"
	"zjhx.com/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户信息
func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	// todo: add your logic here and delete this line

	//检查请求者id
	_, err := l.svcCtx.UserModel.GetUserByID(l.ctx, uint(in.UserId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, code.GetUserError
		}
		logc.Error(l.ctx, err)
		return nil, err
	}

	//获取被请求者信息
	userInfo, err := l.svcCtx.UserModel.GetUserByID(l.ctx, uint(in.ActionId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, code.GetUserError
		}
		logc.Error(l.ctx, err)
		return nil, err
	}

	//获取关注者数量，被关注者数量，是否被关注，点赞数量，记录数量，被点赞数量
	FollowCount := uint32(0)
	FollowedCount := uint32(0)
	IsFollow := true
	if in.ActionId == in.UserId {
		IsFollow = false
	}
	TotalFavorited := uint32(0)
	WorkCount := uint32(0)
	FavoriteCount := uint32(0)
	//
	users := user.UserInfo{
		Id:              uint32(userInfo.Id),
		Name:            userInfo.Name,
		Avatar:          &userInfo.Avatar,
		BackgroundImage: &userInfo.BackgroundImage,
		Signature:       &userInfo.Signature,
		FollowCount:     &FollowCount,
		FollowerCount:   &FollowedCount,
		TotalFavorited:  &TotalFavorited,
		WorkCount:       &WorkCount,
		FavoriteCount:   &FavoriteCount,
		IsFollow:        IsFollow,
	}

	return &user.UserInfoResponse{
		Code: cons.ServiceOKCode,
		Msg:  cons.ServiceOKMsg,
		User: &users,
	}, nil
}
