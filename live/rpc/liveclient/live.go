// Code generated by goctl. DO NOT EDIT.
// Source: live.proto

package liveclient

import (
	"context"

	"zjhx.com/live/rpc/live"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ClosePublishRequest  = live.ClosePublishRequest
	ClosePublishResponse = live.ClosePublishResponse
	CreateRoomRequest    = live.CreateRoomRequest
	CreateRoomResponse   = live.CreateRoomResponse
	GetLiveListInfo      = live.GetLiveListInfo
	GetLiveListRequest   = live.GetLiveListRequest
	GetLiveListResponse  = live.GetLiveListResponse
	JoinRoomInfo         = live.JoinRoomInfo
	JoinRoomRequest      = live.JoinRoomRequest
	JoinRoomResponse     = live.JoinRoomResponse
	SaveLiveRequest      = live.SaveLiveRequest
	SaveLiveResponse     = live.SaveLiveResponse
	SaveRecordRequest    = live.SaveRecordRequest
	SaveRecordResponse   = live.SaveRecordResponse
	StartPublishRequest  = live.StartPublishRequest
	StartPublishResponse = live.StartPublishResponse

	Live interface {
		CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomResponse, error)
		StartPublish(ctx context.Context, in *StartPublishRequest, opts ...grpc.CallOption) (*StartPublishResponse, error)
		ClosePublish(ctx context.Context, in *ClosePublishRequest, opts ...grpc.CallOption) (*ClosePublishResponse, error)
		JoinRoom(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (*JoinRoomResponse, error)
		SaveToRecord(ctx context.Context, in *SaveRecordRequest, opts ...grpc.CallOption) (*SaveRecordResponse, error)
		SaveLive(ctx context.Context, in *SaveLiveRequest, opts ...grpc.CallOption) (*SaveLiveResponse, error)
		GetLiveList(ctx context.Context, in *GetLiveListRequest, opts ...grpc.CallOption) (*GetLiveListResponse, error)
	}

	defaultLive struct {
		cli zrpc.Client
	}
)

func NewLive(cli zrpc.Client) Live {
	return &defaultLive{
		cli: cli,
	}
}

func (m *defaultLive) CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomResponse, error) {
	client := live.NewLiveClient(m.cli.Conn())
	return client.CreateRoom(ctx, in, opts...)
}

func (m *defaultLive) StartPublish(ctx context.Context, in *StartPublishRequest, opts ...grpc.CallOption) (*StartPublishResponse, error) {
	client := live.NewLiveClient(m.cli.Conn())
	return client.StartPublish(ctx, in, opts...)
}

func (m *defaultLive) ClosePublish(ctx context.Context, in *ClosePublishRequest, opts ...grpc.CallOption) (*ClosePublishResponse, error) {
	client := live.NewLiveClient(m.cli.Conn())
	return client.ClosePublish(ctx, in, opts...)
}

func (m *defaultLive) JoinRoom(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (*JoinRoomResponse, error) {
	client := live.NewLiveClient(m.cli.Conn())
	return client.JoinRoom(ctx, in, opts...)
}

func (m *defaultLive) SaveToRecord(ctx context.Context, in *SaveRecordRequest, opts ...grpc.CallOption) (*SaveRecordResponse, error) {
	client := live.NewLiveClient(m.cli.Conn())
	return client.SaveToRecord(ctx, in, opts...)
}

func (m *defaultLive) SaveLive(ctx context.Context, in *SaveLiveRequest, opts ...grpc.CallOption) (*SaveLiveResponse, error) {
	client := live.NewLiveClient(m.cli.Conn())
	return client.SaveLive(ctx, in, opts...)
}

func (m *defaultLive) GetLiveList(ctx context.Context, in *GetLiveListRequest, opts ...grpc.CallOption) (*GetLiveListResponse, error) {
	client := live.NewLiveClient(m.cli.Conn())
	return client.GetLiveList(ctx, in, opts...)
}