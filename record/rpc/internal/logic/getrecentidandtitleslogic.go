package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"zjhx.com/pkg/cons"
	"zjhx.com/record/rpc/internal/svc"
	"zjhx.com/record/rpc/record"
)

type GetRecentIdAndTitlesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRecentIdAndTitlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecentIdAndTitlesLogic {
	return &GetRecentIdAndTitlesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询最近记录
func (l *GetRecentIdAndTitlesLogic) GetRecentIdAndTitles(in *record.RecordByUserIdReq) (*record.RecordByUserIdResp, error) {
	// todo: add your logic here and delete this line

	res, err := l.svcCtx.RecordModel.GetRecentRecorById(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	var IdTitles []*record.IDandTitle
	for _, v := range *res {
		IdTitles = append(IdTitles, &record.IDandTitle{
			Id:         v.Id,
			Title:      v.Title,
			DataType:   uint32(v.DataType),
			UpdateTime: v.UpdateTime.String(),
			Keyword:    v.Keyword,
		})
	}
	return &record.RecordByUserIdResp{
		Code:     cons.ServiceOKCode,
		Msg:      cons.ServiceOKMsg,
		IdTitles: IdTitles,
	}, nil
}
