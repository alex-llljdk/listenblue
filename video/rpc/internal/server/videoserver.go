// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package server

import (
	"context"

	"zjhx.com/video/rpc/internal/logic"
	"zjhx.com/video/rpc/internal/svc"
	"zjhx.com/video/rpc/video"
)

type VideoServer struct {
	svcCtx *svc.ServiceContext
	video.UnimplementedVideoServer
}

func NewVideoServer(svcCtx *svc.ServiceContext) *VideoServer {
	return &VideoServer{
		svcCtx: svcCtx,
	}
}

// 上传文件创建记录
func (s *VideoServer) UploadVideo(ctx context.Context, in *video.UploadRequest) (*video.UploadResponse, error) {
	l := logic.NewUploadVideoLogic(ctx, s.svcCtx)
	return l.UploadVideo(in)
}
