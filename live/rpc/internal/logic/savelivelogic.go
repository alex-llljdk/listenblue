package logic

import (
	"context"

	"zjhx.com/live/rpc/internal/svc"
	"zjhx.com/live/rpc/live"
	"zjhx.com/pkg/cons"
	"zjhx.com/record/gmodel"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveLiveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveLiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveLiveLogic {
	return &SaveLiveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveLiveLogic) SaveLive(in *live.SaveLiveRequest) (*live.SaveLiveResponse, error) {
	// todo: add your logic here and delete this line

	result, err := l.svcCtx.RecordModel.FindOneByPath(l.ctx, in.Path)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	var record = gmodel.Record{
		Path:      in.Path,
		Title:     result.Title,
		UserId:    in.UserId,
		DataType:  result.DataType,
		Keyword:   result.Keyword,
		Summary:   result.Summary,
		Scanning:  result.Scanning,
		InnerText: in.InnerText,
		Review:    result.Review,
		TransText: result.TransText,
	}

	err = l.svcCtx.RecordModel.InsertRecord(l.ctx, &record)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &live.SaveLiveResponse{
		Code: cons.ServiceOKCode,
		Msg:  cons.ServiceOKMsg,
	}, nil
}
