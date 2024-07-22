package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"zjhx.com/pkg/cons"
	"zjhx.com/pkg/utils"
	"zjhx.com/record/api/internal/svc"
	"zjhx.com/record/api/internal/types"
	"zjhx.com/record/code"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAiRewritingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAiRewritingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAiRewritingLogic {
	return &GetAiRewritingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAiRewritingLogic) GetAiRewriting(req *types.AiRewritingRequest) (resp *types.AiRewritingResponse, err error) {
	// todo: add your logic here and delete this line
	//查询出来的结果，进行问答
	var rep []string
	for _, v := range req.Text {
		query := map[string]string{"requestId": uuid.New().String()}
		headers, signParams := utils.GenSignHeaders(cons.AppId, cons.AppKey, "POST", code.GptUri, query)
		headers["Content-Type"] = "application/json"

		var systemPrompt strings.Builder
		systemPrompt.WriteString("请你作为一名文学方面的专家,结合以下一段文本内容回答我所提出的问题,回答需要遵循指定的格式。")
		systemPrompt.WriteString("<Context>")
		systemPrompt.WriteString(v)
		systemPrompt.WriteString("</Context>")

		sessionId := uuid.New().String()
		rurl := fmt.Sprintf("%s%s?%s", code.Domain, code.GptUri, signParams)
		dataAsk := map[string]any{
			"messages": []map[string]string{
				{
					"role":    "system",
					"content": systemPrompt.String(),
				},
				{
					"role":    "user",
					"content": "请对进行用词和语法上的润色,要求改写的内容更加流畅和生动、并确保内容逻辑清晰、易于阅读、不要过分补充。回答的格式为\"{\"content\":\"润色的内容\"}\"",
				},
			},
			"model":     "vivo-BlueLM-TB",
			"sessionId": sessionId,
			"extra": map[string]string{
				"temperature": "0.9",
			},
		}
		//角色预设
		var ChatResponse types.ChatResponse
		response, err := utils.PostBodyRequest(rurl, dataAsk, headers)
		if err != nil {
			l.Logger.Error(err)
			return nil, err
		}
		err = json.Unmarshal([]byte(response), &ChatResponse)
		if err != nil {
			l.Logger.Error(err)
			return nil, err
		}
		rep = append(rep, ChatResponse.Data.Content)
	}

	repbyte, err := json.Marshal(rep)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}

	return &types.AiRewritingResponse{
		Code: cons.ServiceOKCode,
		Msg:  cons.ServiceOKMsg,
		Data: string(repbyte),
	}, nil
}
