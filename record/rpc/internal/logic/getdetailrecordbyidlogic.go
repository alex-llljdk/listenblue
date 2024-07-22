package logic

import (
	"context"

	"zjhx.com/pkg/cons"
	"zjhx.com/record/rpc/internal/svc"
	"zjhx.com/record/rpc/record"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDetailRecordByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDetailRecordByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDetailRecordByIdLogic {
	return &GetDetailRecordByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询某条详细的记录
func (l *GetDetailRecordByIdLogic) GetDetailRecordById(in *record.DetailRecordReq) (*record.Records, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.RecordModel.GetRecorById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &record.Records{
		Code:      cons.ServiceOKCode,
		Msg:       cons.ServiceOKMsg,
		Id:        resp.Id,
		Path:      resp.Path,
		Title:     resp.Title,
		UserId:    resp.UserId,
		DataType:  uint32(resp.DataType),
		Keyword:   resp.Keyword,
		Scanning:  resp.Scanning,
		Summary:   resp.Summary,
		InnerText: resp.InnerText,
		Review:    resp.Review,
		TransText: resp.TransText,
	}, nil
}
