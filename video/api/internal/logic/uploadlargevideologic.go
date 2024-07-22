package logic

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"zjhx.com/pkg/cons"
	"zjhx.com/video/api/internal/svc"
	"zjhx.com/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLargeVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLargeVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLargeVideoLogic {
	return &UploadLargeVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLargeVideoLogic) UploadLargeVideo(filename, userId, dataType, md5, current, total string, file multipart.File) (resp *types.UploadLargeVideoReponse, err error) {
	// todo: add your logic here and delete this line
	dir_path, err := l.CheckAndCreate(l.svcCtx.Config.FilePath, userId, md5, dataType, total)
	if err != nil {
		return nil, err
	}
	//拼接文件名和路径
	filename_split := strings.Split(filename, ".")
	path := fmt.Sprintf("%s/%s_%s.%s", dir_path, filename_split[0], current, filename_split[1])
	//写入文件
	err = l.WriteFile(&file, path)
	if err != nil {
		return nil, err
	}

	return &types.UploadLargeVideoReponse{
		Code: cons.ServiceOKCode,
		Msg:  cons.ServiceOKMsg,
	}, nil
}

func (l *UploadLargeVideoLogic) CheckAndCreate(path, userId, hash, dataType, total string) (string, error) {
	//检查用户文件夹是否存在
	switch dataType {
	case "0":
		path = strings.Join([]string{path, "voice", userId, hash + "_" + total}, "/")
	case "1":
		path = strings.Join([]string{path, "video", userId, hash + "_" + total}, "/")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 目录不存在，创建目录
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	return path, nil
}

func (l *UploadLargeVideoLogic) WriteFile(file *multipart.File, path string) error {
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
