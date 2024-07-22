package logic

import (
	"context"
	"encoding/json"

	"zjhx.com/live/code"
	"zjhx.com/live/gmodel"
	"zjhx.com/live/rpc/internal/svc"
	"zjhx.com/live/rpc/live"
	"zjhx.com/pkg/cons"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type JoinRoomLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJoinRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinRoomLogic {
	return &JoinRoomLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JoinRoomLogic) JoinRoom(in *live.JoinRoomRequest) (*live.JoinRoomResponse, error) {
	// todo: add your logic here and delete this line

	exist, err := l.svcCtx.BizRedis.ExistsCtx(l.ctx, cons.LiveKey+in.RoomNumber)
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, err
	}
	//房间不存在
	if !exist {
		logc.Error(l.ctx, code.RoomNotExistErr)
		return nil, code.RoomNotExistErr
	}

	//房间存在则获取到房间信息，并返回
	res, err := l.svcCtx.BizRedis.GetCtx(l.ctx, cons.LiveKey+in.RoomNumber)
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, err
	}
	var roomInfo gmodel.LiveRoom
	err = json.Unmarshal([]byte(res), &roomInfo)
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, err
	}

	//判断密码是否正确,密码不正确，返回错误
	if !roomInfo.IsOpen && in.Password != roomInfo.Password {
		logc.Error(l.ctx, err)
		return nil, code.PasswordErr
	}

	return &live.JoinRoomResponse{
		Code: cons.ServiceOKCode,
		Msg:  cons.ServiceOKMsg,
		RoomInfo: &live.JoinRoomInfo{
			RoomTitle:  roomInfo.RoomTitle,
			IsOpen:     roomInfo.IsOpen,
			Password:   roomInfo.Password,
			IsRecord:   roomInfo.IsRecord,
			RoomNumber: roomInfo.RoomNumBer,
			UserId:     int32(roomInfo.UserID),
			Url:        roomInfo.Url,
		},
	}, nil
}
