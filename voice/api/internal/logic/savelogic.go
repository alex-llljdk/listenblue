package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"

	"zjhx.com/pkg/cons"
	"zjhx.com/pkg/utils"
	"zjhx.com/voice/api/internal/svc"
	"zjhx.com/voice/api/internal/types"
	"zjhx.com/voice/code"
	"zjhx.com/voice/rpc/voice"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type SaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveLogic {
	return &SaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveLogic) Save(filename, userId, dataType, sourceContent, innerHtml, title string, file multipart.File) (resp *types.SaveResponse, err error) {
	// todo: add your logic here and delete this line

	//将文件进行md5编码
	//获取地址的hash码
	hasher := md5.New()
	hasher.Write([]byte(sourceContent))
	hashByte := hasher.Sum(nil)
	hash := hex.EncodeToString(hashByte)

	urlPath := strings.Join([]string{cons.RequestIp, "voice", userId, hash, filename + ".mp3"}, "/")
	//检查创建目录
	dir_path, err := l.CheckAndCreate(l.svcCtx.Config.FilePath, userId, hash, dataType)
	if err != nil {
		return nil, err
	}
	path := strings.Join([]string{dir_path, fmt.Sprintf("%s.mp3", filename)}, "/")
	// 写入文件
	fmt.Println(file)
	err = l.WriteFile(&file, path)
	if err != nil {
		return nil, err
	}

	var SC []types.SourceContent
	// 解码JSON字符串到结构体
	err = json.Unmarshal([]byte(sourceContent), &SC)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}

	// 打印解码后的结构体内容

	query := map[string]string{"requestId": uuid.New().String()}
	headers, signParams := utils.GenSignHeaders(cons.AppId, cons.AppKey, "POST", code.GptUri, query)
	headers["Content-Type"] = "application/json"

	var systemPrompt strings.Builder
	systemPrompt.WriteString("请你作为一名文学方面的专家,结合以下一段文本内容回答我所提出的问题,回答需要遵循指定的格式。")
	systemPrompt.WriteString("<Context>")
	for _, v := range SC {
		systemPrompt.WriteString(v.Content)
	}
	systemPrompt.WriteString("</Context>")
	fmt.Println(systemPrompt.String())
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
				"content": "请帮我提炼出关键词(不多于20个关键词),请返回json格式\"{\"answer\": [\"关键词1\",\"关键词2\",...\"关键词n\"]}\",如果不需要额外回复，则对应字段输出空。",
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
	response = strings.Replace(response, " ", "", -1)
	response = strings.Replace(response, "\n", "", -1)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	err = json.Unmarshal([]byte(response), &ChatResponse)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	keyword := ChatResponse.Data.Content
	fmt.Println(response)

	ChatResponse = types.ChatResponse{}
	//全文概要
	dataAsk = map[string]any{
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": systemPrompt.String(),
			},
			{
				"role":    "user",
				"content": "请对内容进行总结得出一个简略的全文概要(不超过512个字),请返回json格式\"{\"answer\": \"给用户的回复\"}\",如果不需要额外回复，则对应字段输出空。",
			},
		},
		"model":     "vivo-BlueLM-TB",
		"sessionId": sessionId,
		"extra": map[string]string{
			"temperature": "0.9",
		},
	}
	response, err = utils.PostBodyRequest(rurl, dataAsk, headers)
	response = strings.Replace(response, " ", "", -1)
	response = strings.Replace(response, "\n", "", -1)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	err = json.Unmarshal([]byte(response), &ChatResponse)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	summary := ChatResponse.Data.Content
	fmt.Println(response)
	ChatResponse = types.ChatResponse{}
	//章节速览 总结出一个标题，并返回对应在那句话
	dataAsk = map[string]any{
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": systemPrompt.String(),
			},
			{
				"role":    "user",
				"content": "请对内容归纳章节(尽可能将多句话归纳为一个章节),包含章节标题、每个章节概括性的阐述(每个章节不超过128个字),请返回json格式\"{\"answer\":[{\"title\":\"返回的标题\",\"desc\":\"返回的阐述\"},{\"title\":\"返回的标题\",\"desc\":\"返回的阐述\"}]}\",如果不需要额外回复，则对应字段输出空。",
			},
		},
		"model":     "vivo-BlueLM-TB",
		"sessionId": sessionId,
		"extra": map[string]string{
			"temperature": "0.9",
		},
	}
	response, err = utils.PostBodyRequest(rurl, dataAsk, headers)
	response = strings.Replace(response, " ", "", -1)
	response = strings.Replace(response, "\n", "", -1)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	err = json.Unmarshal([]byte(response), &ChatResponse)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	scanning := ChatResponse.Data.Content
	fmt.Println(response)
	ChatResponse = types.ChatResponse{}
	//要点回顾
	dataAsk = map[string]any{
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": systemPrompt.String(),
			},
			{
				"role":    "user",
				"content": "请从内容中归纳出几个要点问题,并根据要点问题和内容分析出答案,请返回json格式\"{\"answer\":[{\"ques\":\"返回的问题\",\"ans\":\"问题的答案\"},{\"ques\":\"返回的问题\",\"ans\":\"问题的答案\"}]}\",如果不需要额外回复，则对应字段输出空。",
			},
		},
		"model":     "vivo-BlueLM-TB",
		"sessionId": sessionId,
		"extra": map[string]string{
			"temperature": "0.9",
		},
	}
	response, err = utils.PostBodyRequest(rurl, dataAsk, headers)
	response = strings.Replace(response, " ", "", -1)
	response = strings.Replace(response, "\n", "", -1)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	err = json.Unmarshal([]byte(response), &ChatResponse)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	review := ChatResponse.Data.Content
	fmt.Println(response)

	u_id, _ := strconv.Atoi(userId)
	dt, _ := strconv.Atoi(dataType)
	l.svcCtx.VoiceRpc.UploadVoice(l.ctx, &voice.UploadRequest{
		Path:      urlPath,
		UserId:    uint64(u_id),
		DataType:  uint32(dt),
		Keyword:   keyword,
		Summary:   summary,
		Scanning:  scanning,
		Review:    review,
		TransText: sourceContent,
		Innertext: innerHtml,
		Title:     title,
	})

	return &types.SaveResponse{
		Code: cons.ServiceOKCode,
		Msg:  cons.ServiceOKMsg,
	}, nil
}

func (l *SaveLogic) CheckAndCreate(path, userId, hash, dataType string) (string, error) {
	//检查用户文件夹是否存在
	path = strings.Join([]string{path, "voice", userId, hash}, "/")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 目录不存在，创建目录
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	return path, nil
}

func (l *SaveLogic) WriteFile(file *multipart.File, path string) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	//写入文件
	defer out.Close()
	_, err = io.Copy(out, *file)
	if err != nil {
		return err
	}
	return nil
}
