package logic

import (
	"context"
	"fmt"
	"net/http"

	"zjhx.com/live/rpc/internal/svc"
	"zjhx.com/live/rpc/live"
	"zjhx.com/pkg/cons"
	"zjhx.com/pkg/encrypt"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type ClosePublishLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClosePublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClosePublishLogic {
	return &ClosePublishLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClosePublishLogic) ClosePublish(in *live.ClosePublishRequest) (*live.ClosePublishResponse, error) {
	// todo: add your logic here and delete this line
	//删除redis中的房间记录，同时中断直播推流
	//查询redis中是否有房间记录，如果没有则返回错误
	roomNumber, err := encrypt.DecRoomNumber(in.EncRoomNumber)
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, err
	}

	//获取房间信息
	res, err := l.svcCtx.BizRedis.GetCtx(l.ctx, cons.LiveKey+roomNumber)
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, err
	}

	//删除房间内容
	_, err = l.svcCtx.BizRedis.DelCtx(l.ctx, cons.LiveKey+roomNumber)
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, err
	}

	//删除房间列表信息
	_, err = l.svcCtx.BizRedis.LremCtx(l.ctx, "RoomList", 1, roomNumber)
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, err
	}

	//中断推流
	_, err = http.Get(fmt.Sprintf("%s/control/record/stop?app=live&name=%s", cons.RequestIp, in.EncRoomNumber))
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, err
	}
	// _, err = http.Get(fmt.Sprintf("%s/control/drop/publisher?app=live&name=%s", cons.RequestIp, in.EncRoomNumber))
	// if err != nil {
	// 	logc.Error(l.ctx, err)
	// 	return nil, err
	// }

	return &live.ClosePublishResponse{
		Code: cons.ServiceOKCode,
		Msg:  cons.ServiceOKMsg,
		Res:  res,
	}, nil
}
