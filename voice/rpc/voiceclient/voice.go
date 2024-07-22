// Code generated by goctl. DO NOT EDIT.
// Source: voice.proto

package voiceclient

import (
	"context"

	"zjhx.com/voice/rpc/voice"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	UploadRequest  = voice.UploadRequest
	UploadResponse = voice.UploadResponse

	Voice interface {
		// 上传文件创建记录
		UploadVoice(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error)
	}

	defaultVoice struct {
		cli zrpc.Client
	}
)

func NewVoice(cli zrpc.Client) Voice {
	return &defaultVoice{
		cli: cli,
	}
}

// 上传文件创建记录
func (m *defaultVoice) UploadVoice(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error) {
	client := voice.NewVoiceClient(m.cli.Conn())
	return client.UploadVoice(ctx, in, opts...)
}
