package logic

import (
	"context"

	"zjhx.com/live/api/internal/svc"
	"zjhx.com/live/api/internal/types"
	"zjhx.com/live/rpc/live"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJoinRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinRoomLogic {
	return &JoinRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JoinRoomLogic) JoinRoom(req *types.JoinRoomRequest) (resp *types.JoinRoomResponse, err error) {
	// todo: add your logic here and delete this line

	//调用rpc检查房间是否存在,同时返回房间信息
	res, err := l.svcCtx.LiveRpc.JoinRoom(l.ctx, &live.JoinRoomRequest{
		RoomNumber: req.RoomNumber,
		Password:   req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.JoinRoomResponse{
		Code: int(res.Code),
		Msg:  res.Msg,
		Data: types.JoinData{
			UserID:     int(res.RoomInfo.UserId),
			RoomTitle:  res.RoomInfo.RoomTitle,
			IsOpen:     res.RoomInfo.IsOpen,
			Password:   res.RoomInfo.Password,
			IsRecord:   res.RoomInfo.IsRecord,
			RoomNumBer: res.RoomInfo.RoomNumber,
			Url:        res.RoomInfo.Url,
		},
	}, nil
}
