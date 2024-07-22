package logic

import (
	"context"

	"zjhx.com/record/api/internal/svc"
	"zjhx.com/record/api/internal/types"
	"zjhx.com/record/rpc/record"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecordByUseridLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecordByUseridLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecordByUseridLogic {
	return &RecordByUseridLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecordByUseridLogic) RecordByUserid(req *types.RecordByUserIdRequest) (resp *types.RecordByUserIdRespone, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.RecordRpc.GetIdAndTitles(l.ctx, &record.RecordByUserIdReq{UserId: req.UserId})
	//
	if err != nil {
		return nil, err
	}
	resplen := len(res.IdTitles)
	var idtitles = make([]types.SingleRecord, resplen) // 使用 make 创建切片
	for index, v := range res.IdTitles {
		idtitles[index] = types.SingleRecord{
			Id:         v.Id,
			Title:      v.Title,
			UpdateTime: v.UpdateTime,
			Keyword:    v.Keyword,
			DataType:   v.DataType,
		}
	}

	return &types.RecordByUserIdRespone{
		SingleRecords: idtitles,
		Code:          200,
		Msg:           "请求成功",
	}, nil
}
