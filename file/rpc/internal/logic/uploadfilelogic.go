package logic

import (
	"context"
	"fmt"

	"zjhx.com/file/code"
	"zjhx.com/file/gmodel"
	"zjhx.com/file/rpc/file"
	"zjhx.com/file/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 上传文件创建记录
func (l *UploadFileLogic) UploadFile(in *file.FileRecord) (*file.UploadResponse, error) {
	//先检查文件路径是否存在，如果存在则创建添加记录信息
	// _, err := os.Stat(in.Path)
	// if err != nil {
	// 	if os.IsNotExist(err) {
	// 		return nil, code.UploadErr
	// 	} else {
	// 		return nil, err
	// 	}
	// }

	//创建记录信息
	recordId, err := l.svcCtx.RecordModel.InsertRecord(l.ctx, &gmodel.Record{
		Path:     in.Path,
		Title:    in.Title,
		UserId:   uint64(in.UserId),
		DataType: uint(in.DataType),
		Keyword:  in.Keyword,
		Summary:  in.Summary,
		Scanning: in.Scanning,
	})
	if err != nil {
		return nil, code.CreateRecordErr
	}

	fmt.Println("recordid--------------", recordId)
	return &file.UploadResponse{
		Code:     200,
		Msg:      "记录保存成功",
		RecordId: int32(recordId),
	}, nil
}
