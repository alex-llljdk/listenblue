package logic

import (
	"context"

	"zjhx.com/pkg/cons"
	"zjhx.com/record/rpc/internal/svc"
	"zjhx.com/record/rpc/record"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIdAndTitlesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIdAndTitlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIdAndTitlesLogic {
	return &GetIdAndTitlesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户对应的记录和标题
func (l *GetIdAndTitlesLogic) GetIdAndTitles(in *record.RecordByUserIdReq) (*record.RecordByUserIdResp, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.RecordModel.GetListRecorById(l.ctx, int32(in.UserId))
	if err != nil {
		return nil, err
	}
	resplen := len(resp)
	var final_res = make([]*record.IDandTitle, resplen) // 使用 make 创建切片
	for index, v := range resp {
		final_res[index] = &record.IDandTitle{
			Id:         v.Id,
			Title:      v.Title,
			DataType:   uint32(v.DataType),
			UpdateTime: v.UpdateTime.String(),
			Keyword:    v.Keyword,
		}
	}

	return &record.RecordByUserIdResp{
		Code:     cons.ServiceOKCode,
		Msg:      cons.ServiceOKMsg,
		IdTitles: final_res,
	}, nil
}
