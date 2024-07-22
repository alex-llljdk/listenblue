package logic

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"zjhx.com/live/api/internal/svc"
	"zjhx.com/live/api/internal/types"
	"zjhx.com/live/rpc/live"
	"zjhx.com/pkg/cons"
	"zjhx.com/pkg/encrypt"
)

type CreateRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoomLogic {
	return &CreateRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateRoomLogic) CreateRoom(req *types.CreateRoomRequest) (resp *types.CreateRoomResponse, err error) {
	//创建房间，返回直播的推流地址
	//房间写入数据库和redis
	//先对房间号进行编码
	encRoomNumber, err := encrypt.EncRoomNumber(req.RoomNumber)
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, err
	}

	res, err := l.svcCtx.LiveRpc.CreateRoom(l.ctx, &live.CreateRoomRequest{
		RoomTitle:     req.RoomTitle,
		IsOpen:        req.IsOpen,
		Password:      req.Password,
		IsRecord:      req.IsRecord,
		RoomNumber:    req.RoomNumber,
		UserId:        int32(req.UserId),
		EncRoomNumber: encRoomNumber,
		UserName:      req.UserName,
		UserAvatar:    req.UserAvatar,
	})
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, err
	}

	//定时器1小时关闭主机推流, 关闭录制
	go func(encRoomNumber string) {
		timer := time.NewTimer(60 * 60 * time.Second)
		<-timer.C
		_, err := http.Get(fmt.Sprintf("%s/control/drop/publisher?app=live&name=%s", cons.RequestIp, encRoomNumber))
		if err != nil {
			logc.Error(l.ctx, err)
			return
		}
	}(encRoomNumber)

	return &types.CreateRoomResponse{
		Code: int(res.Code),
		Msg:  res.Msg,
		Url:  res.Url,
	}, nil
}
