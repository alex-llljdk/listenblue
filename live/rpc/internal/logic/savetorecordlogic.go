package logic

import (
	"context"
	"fmt"

	"zjhx.com/live/code"
	"zjhx.com/live/rpc/internal/svc"
	"zjhx.com/live/rpc/live"
	"zjhx.com/pkg/cons"
	"zjhx.com/record/gmodel"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveToRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveToRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveToRecordLogic {
	return &SaveToRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveToRecordLogic) SaveToRecord(in *live.SaveRecordRequest) (*live.SaveRecordResponse, error) {
	err := l.svcCtx.RecordModel.InsertRecord(l.ctx, &gmodel.Record{
		Path:      in.Path,
		Title:     in.Title,
		UserId:    in.UserId,
		DataType:  uint(in.DataType),
		Keyword:   in.Keyword,
		Summary:   in.Summary,
		Scanning:  in.Scanning,
		Review:    in.Review,
		TransText: in.TransText,
		InnerText: in.InnerText,
	})
	if err != nil {
		fmt.Println(err)
		return nil, code.InsertRecordErr
	}

	return &live.SaveRecordResponse{
		Code: cons.ServiceOKCode,
		Msg:  cons.ServiceOKMsg,
	}, nil
}
