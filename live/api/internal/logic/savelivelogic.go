package logic

import (
	"context"
	"fmt"
	"strings"

	"zjhx.com/live/api/internal/svc"
	"zjhx.com/live/api/internal/types"
	"zjhx.com/live/code"
	"zjhx.com/live/rpc/live"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveLiveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveLiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveLiveLogic {
	return &SaveLiveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveLiveLogic) SaveLive(req *types.SaveRequest) (resp *types.SaveResponse, err error) {
	// todo: add your logic here and delete this line

	sppath := strings.Split(req.Path, "=")
	path := fmt.Sprintf("%s/%s.flv", code.LiveNetUrl, sppath[len(sppath)-1])
	fmt.Println(path)
	res, err := l.svcCtx.LiveRpc.SaveLive(l.ctx, &live.SaveLiveRequest{
		Path:      path,
		UserId:    uint64(req.UserID),
		InnerText: req.InnerText,
	})
	if err != nil {
		return nil, err
	}
	return &types.SaveResponse{
		Code: int(res.Code),
		Msg:  res.Msg,
	}, nil
}
