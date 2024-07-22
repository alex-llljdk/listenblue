package logic

import (
	"context"

	"zjhx.com/live/api/internal/svc"
	"zjhx.com/live/api/internal/types"
	"zjhx.com/live/rpc/live"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetLiveListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLiveListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLiveListLogic {
	return &GetLiveListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLiveListLogic) GetLiveList() (resp *types.GetLiveListResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.LiveRpc.GetLiveList(l.ctx, &live.GetLiveListRequest{})
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, err
	}

	ret := make([]types.LiveListData, len(res.LiveList))
	for i, v := range res.LiveList {
		ret[i].UserName = v.UserName
		ret[i].RoomNumBer = v.RoomNumber
		ret[i].RoomTitle = v.RoomTitle
		ret[i].CoverUrl = v.CoverUrl
		ret[i].UserAvatar = v.UserAvatar
	}

	return &types.GetLiveListResponse{
		Code:     int(res.Code),
		Msg:      res.Msg,
		LiveList: ret,
	}, nil
}
