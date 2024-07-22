package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"zjhx.com/pkg/cons"
	"zjhx.com/pkg/utils"
	"zjhx.com/voice/api/internal/svc"
	"zjhx.com/voice/api/internal/types"
	"zjhx.com/voice/code"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type TransLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTransLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransLogic {
	return &TransLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TransLogic) Trans(req *types.TransRequest) (resp *types.TransResponse, err error) {

	headers, _ := utils.GenSignHeaders(cons.AppId, cons.AppKey, "POST", code.TransUri, map[string]string{})
	data := map[string]string{
		"from":      req.From,
		"to":        req.To,
		"text":      req.Text,
		"app":       "test",
		"requestId": uuid.New().String(),
	}
	rurl := fmt.Sprintf("%s%s", code.Domain, code.TransUri)
	response, err := utils.PostUrlencodedRequest(rurl, data, headers)
	if err != nil {
		return nil, err
	}
	var TransTextResponse types.TransTextResponse
	err = json.Unmarshal([]byte(response), &TransTextResponse)
	if err != nil {
		return nil, err
	}

	return &types.TransResponse{
		Code: cons.ServiceOKCode,
		Msg:  cons.ServiceOKMsg,
		Data: TransTextResponse.Data.Translation,
	}, nil
}
