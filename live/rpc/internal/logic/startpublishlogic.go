package logic

import (
	"context"
	"errors"

	"zjhx.com/live/code"
	"zjhx.com/live/rpc/internal/svc"
	"zjhx.com/live/rpc/live"
	"zjhx.com/pkg/cons"
	"zjhx.com/pkg/encrypt"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type StartPublishLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStartPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartPublishLogic {
	return &StartPublishLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StartPublishLogic) StartPublish(in *live.StartPublishRequest) (*live.StartPublishResponse, error) {
	// todo: add your logic here and delete this line
	//查询redis中是否有房间记录，如果没有则返回错误，不许开房间
	roomNumber, err := encrypt.DecRoomNumber(in.EncRoomNumber)
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, errors.New("解码错误")
	}
	exist, err := l.svcCtx.BizRedis.ExistsCtx(l.ctx, cons.LiveKey+roomNumber)
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, code.RedisErr
	}
	if !exist {
		logc.Error(l.ctx, code.RoomNotExistErr)
		return nil, code.RoomNotExistErr
	}

	res, err := l.svcCtx.BizRedis.GetCtx(l.ctx, cons.LiveKey+roomNumber)
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, code.RedisErr
	}

	return &live.StartPublishResponse{
		Code: cons.ServiceOKCode,
		Msg:  cons.ServiceOKMsg,
		Res:  res,
	}, nil
}
