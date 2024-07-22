package logic

import (
	"context"

	"zjhx.com/record/api/internal/svc"
	"zjhx.com/record/api/internal/types"
	"zjhx.com/record/rpc/record"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecentRecordByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecentRecordByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecentRecordByIdLogic {
	return &RecentRecordByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecentRecordByIdLogic) RecentRecordById(req *types.RecordByIdRequest) (resp *types.RecordByUserIdRespone, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.RecordRpc.GetRecentIdAndTitles(l.ctx, &record.RecordByUserIdReq{
		UserId: req.Id,
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	var singleRecord []types.SingleRecord
	for _, v := range res.IdTitles {
		singleRecord = append(singleRecord, types.SingleRecord{
			Id:         v.Id,
			Title:      v.Title,
			UpdateTime: v.UpdateTime,
			Keyword:    v.Keyword,
			DataType:   v.DataType,
		})
	}
	return &types.RecordByUserIdRespone{
		Msg:           res.Msg,
		Code:          int(res.Code),
		SingleRecords: singleRecord,
	}, nil
}
