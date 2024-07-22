package logic

import (
	"context"

	"zjhx.com/record/rpc/internal/svc"
	"zjhx.com/record/rpc/record"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDetailRecordByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDetailRecordByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDetailRecordByIdLogic {
	return &UpdateDetailRecordByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新某条详细的记录
func (l *UpdateDetailRecordByIdLogic) UpdateDetailRecordById(in *record.Records) (*record.BasicRespone, error) {
	// todo: add your logic here and delete this line

	return &record.BasicRespone{}, nil
}
