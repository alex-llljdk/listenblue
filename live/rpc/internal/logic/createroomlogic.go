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

type CreateRoomLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoomLogic {
	return &CreateRoomLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateRoomLogic) CreateRoom(in *live.CreateRoomRequest) (*live.CreateRoomResponse, error) {
	// todo: add your logic here and delete this line

	//先检查是否有相同房间号 redis
	exist, err := l.svcCtx.BizRedis.ExistsCtx(l.ctx, cons.LiveKey+in.RoomNumber)
	if err != nil {
		logc.Error(l.ctx, err)
	}
	if exist {
		// 记录存在
		logc.Infof(l.ctx, "房间已存在")
		return nil, code.RoomExistErr
	}

	//没有相同房间号则创建数据库记录和redis记录
	LR := &gmodel.LiveRoom{
		RoomNumBer: in.RoomNumber,
		RoomTitle:  in.RoomTitle,
		IsOpen:     in.IsOpen,
		Password:   in.Password,
		IsRecord:   in.IsRecord,
		UserID:     int(in.UserId),
		Url:        cons.LiveUrl + in.EncRoomNumber,
		CoverUrl:   code.LiveNetUrl + "/" + in.EncRoomNumber + ".jpg",
		UserName:   in.UserName,
		Avatar:     in.UserAvatar,
	}
	err = l.svcCtx.LiveModel.InsertLiveRoom(l.ctx, LR)
	if err != nil {
		logc.Error(l.ctx, "InsertLiveRoom req: %v error %v", in, err)
		return nil, err
	}

	rdsLr, _ := json.Marshal(LR)

	//创建记录
	err = l.svcCtx.BizRedis.SetexCtx(l.ctx, cons.LiveKey+in.RoomNumber, string(rdsLr), 60*65)
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, err
	}

	//房间列表添加记录
	_, err = l.svcCtx.BizRedis.LpushCtx(l.ctx, "RoomList", in.RoomNumber)
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, err
	}

	return &live.CreateRoomResponse{
		Code: cons.ServiceOKCode,
		Msg:  cons.ServiceOKMsg,
		Url:  cons.RtmpIP + "/" + in.EncRoomNumber,
	}, nil
}
