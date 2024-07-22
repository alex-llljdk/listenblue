package logic

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"zjhx.com/pkg/cons"
	"zjhx.com/user/api/internal/svc"
	"zjhx.com/user/api/internal/types"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadImgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadImgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadImgLogic {
	return &UploadImgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadImgLogic) UploadImg(filename string, file multipart.File) (resp *types.UploadImageResponse, err error) {
	// todo: add your logic here and delete this line
	err = l.CheckAndCreate(fmt.Sprintf("%s/ava", l.svcCtx.Config.FilePath))
	if err != nil {
		return nil, err
	}

	spFileName := strings.Split(filename, ".")
	datatype := spFileName[len(spFileName)-1]
	name := uuid.New().String()
	path := fmt.Sprintf("%s/ava/%s.%s", l.svcCtx.Config.FilePath, name, datatype)

	err = l.WriteFile(&file, path)
	if err != nil {
		return nil, err
	}

	return &types.UploadImageResponse{
		Code: cons.ServiceOKCode,
		Msg:  cons.ServiceOKMsg,
		Url:  fmt.Sprintf("%s/ava/%s.%s", cons.RequestIp, name, datatype),
	}, err
}

func (l *UploadImgLogic) CheckAndCreate(path string) error {
	//检查用户文件夹是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 目录不存在，创建目录
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *UploadImgLogic) WriteFile(file *multipart.File, path string) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	//写入文件
	defer out.Close()
	_, err = io.Copy(out, *file)
	if err != nil {
		return err
	}
	return nil
}
