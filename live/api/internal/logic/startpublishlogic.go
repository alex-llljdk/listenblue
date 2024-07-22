package logic

import (
	"context"
	"encoding/json"

	"zjhx.com/live/api/internal/svc"
	"zjhx.com/live/api/internal/types"
	"zjhx.com/live/gmodel"
	"zjhx.com/live/rpc/live"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartPublishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStartPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartPublishLogic {
	return &StartPublishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartPublishLogic) StartPublish(req *types.OnPublishRequest) (resp *types.OnPublishResponse, err error) {
	// todo: add your logic here and delete this line

	res, err := l.svcCtx.LiveRpc.StartPublish(l.ctx, &live.StartPublishRequest{
		EncRoomNumber: req.Name,
	})
	if err != nil {
		return nil, err
	}

	var resStruct gmodel.LiveRoom
	err = json.Unmarshal([]byte(res.Res), &resStruct)
	if err != nil {
		return nil, err
	}

	return &types.OnPublishResponse{
		Code:     int(res.Code),
		Msg:      res.Msg,
		IsRecord: resStruct.IsRecord,
	}, nil
}
