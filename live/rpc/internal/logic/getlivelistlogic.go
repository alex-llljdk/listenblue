package logic

import (
	"context"
	"encoding/json"

	"zjhx.com/live/gmodel"
	"zjhx.com/live/rpc/internal/svc"
	"zjhx.com/live/rpc/live"
	"zjhx.com/pkg/cons"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetLiveListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLiveListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLiveListLogic {
	return &GetLiveListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLiveListLogic) GetLiveList(in *live.GetLiveListRequest) (*live.GetLiveListResponse, error) {
	// todo: add your logic here and delete this line

	//获取目前房间列表
	res, err := l.svcCtx.BizRedis.LrangeCtx(l.ctx, "RoomList", 0, -1)
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, err
	}

	length := len(res)
	var liveList = make([]*live.GetLiveListInfo, length)
	//遍历房间列表，返回所有房间信息
	for i, v := range res {
		res, err := l.svcCtx.BizRedis.GetCtx(l.ctx, cons.LiveKey+v)
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
		if !roomInfo.IsOpen {
			continue
		}
		liveList[i] = &live.GetLiveListInfo{
			UserName:   roomInfo.UserName,
			RoomNumber: roomInfo.RoomNumBer,
			RoomTitle:  roomInfo.RoomTitle,
			CoverUrl:   roomInfo.CoverUrl,
			UserAvatar: roomInfo.Avatar,
		}
	}

	return &live.GetLiveListResponse{
		Code:     cons.ServiceOKCode,
		Msg:      cons.ServiceOKMsg,
		LiveList: liveList,
	}, nil
}
