package logic

import (
	"context"

	"zjhx.com/pkg/cons"
	"zjhx.com/video/code"
	"zjhx.com/video/gmodel"
	"zjhx.com/video/rpc/internal/svc"
	"zjhx.com/video/rpc/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadVideoLogic {
	return &UploadVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 上传文件创建记录
func (l *UploadVideoLogic) UploadVideo(in *video.UploadRequest) (*video.UploadResponse, error) {

	//创建记录信息
	recordId, err := l.svcCtx.RecordModel.InsertRecord(l.ctx, &gmodel.Record{
		Path:      in.Path,
		Title:     in.Title,
		UserId:    uint64(in.UserId),
		DataType:  uint(in.DataType),
		Keyword:   in.Keyword,
		Summary:   in.Summary,
		Scanning:  in.Scanning,
		InnerText: in.Innertext,
		Review:    in.Review,
		TransText: in.TransText,
	})
	if err != nil {
		logx.Error(err)
		return nil, code.CreateRecordErr
	}

	return &video.UploadResponse{
		Code:     cons.ServiceOKCode,
		Msg:      cons.ServiceOKMsg,
		RecordId: int32(recordId),
	}, nil
}
