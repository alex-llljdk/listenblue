// Code generated by goctl. DO NOT EDIT.
// Source: live.proto

package server

import (
	"context"

	"zjhx.com/live/rpc/internal/logic"
	"zjhx.com/live/rpc/internal/svc"
	"zjhx.com/live/rpc/live"
)

type LiveServer struct {
	svcCtx *svc.ServiceContext
	live.UnimplementedLiveServer
}

func NewLiveServer(svcCtx *svc.ServiceContext) *LiveServer {
	return &LiveServer{
		svcCtx: svcCtx,
	}
}

func (s *LiveServer) CreateRoom(ctx context.Context, in *live.CreateRoomRequest) (*live.CreateRoomResponse, error) {
	l := logic.NewCreateRoomLogic(ctx, s.svcCtx)
	return l.CreateRoom(in)
}

func (s *LiveServer) StartPublish(ctx context.Context, in *live.StartPublishRequest) (*live.StartPublishResponse, error) {
	l := logic.NewStartPublishLogic(ctx, s.svcCtx)
	return l.StartPublish(in)
}

func (s *LiveServer) ClosePublish(ctx context.Context, in *live.ClosePublishRequest) (*live.ClosePublishResponse, error) {
	l := logic.NewClosePublishLogic(ctx, s.svcCtx)
	return l.ClosePublish(in)
}

func (s *LiveServer) JoinRoom(ctx context.Context, in *live.JoinRoomRequest) (*live.JoinRoomResponse, error) {
	l := logic.NewJoinRoomLogic(ctx, s.svcCtx)
	return l.JoinRoom(in)
}

func (s *LiveServer) SaveToRecord(ctx context.Context, in *live.SaveRecordRequest) (*live.SaveRecordResponse, error) {
	l := logic.NewSaveToRecordLogic(ctx, s.svcCtx)
	return l.SaveToRecord(in)
}

func (s *LiveServer) SaveLive(ctx context.Context, in *live.SaveLiveRequest) (*live.SaveLiveResponse, error) {
	l := logic.NewSaveLiveLogic(ctx, s.svcCtx)
	return l.SaveLive(in)
}

func (s *LiveServer) GetLiveList(ctx context.Context, in *live.GetLiveListRequest) (*live.GetLiveListResponse, error) {
	l := logic.NewGetLiveListLogic(ctx, s.svcCtx)
	return l.GetLiveList(in)
}
