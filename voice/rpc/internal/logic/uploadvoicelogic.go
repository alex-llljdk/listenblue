package logic

import (
	"context"

	"zjhx.com/pkg/cons"
	"zjhx.com/record/gmodel"
	"zjhx.com/voice/code"
	"zjhx.com/voice/rpc/internal/svc"
	"zjhx.com/voice/rpc/voice"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadVoiceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadVoiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadVoiceLogic {
	return &UploadVoiceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 上传文件创建记录
func (l *UploadVoiceLogic) UploadVoice(in *voice.UploadRequest) (*voice.UploadResponse, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.RecordModel.InsertRecord(l.ctx, &gmodel.Record{
		Path:      in.Path,
		UserId:    in.UserId,
		DataType:  uint(in.DataType),
		Keyword:   in.Keyword,
		Summary:   in.Summary,
		Scanning:  in.Scanning,
		Review:    in.Review,
		TransText: in.TransText,
		InnerText: in.Innertext,
		Title:     in.Title,
	})
	if err != nil {
		return nil, code.CreateRecordErr
	}
	return &voice.UploadResponse{
		Code: cons.ServiceOKCode,
		Msg:  cons.ServiceOKMsg,
	}, nil
}
