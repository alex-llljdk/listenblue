package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"

	"zjhx.com/pkg/cons"
	"zjhx.com/pkg/utils"
	"zjhx.com/video/api/internal/svc"
	"zjhx.com/video/api/internal/types"
	"zjhx.com/video/code"
	"zjhx.com/video/rpc/video"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadVideoLogic {
	return &UploadVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadVideoLogic) UploadVideo(filename, userId, dataType, hash, source_language, dest_language string, file multipart.File) (resp *types.UploadVideoReponse, err error) {
	// todo: add your logic here and delete this line
	//检查创建目录
	dir_path, urlPath, err := l.CheckAndCreate(l.svcCtx.Config.FilePath, userId, hash, dataType)
	if err != nil {
		return nil, err
	}
	path := strings.Join([]string{dir_path, filename}, "/")
	//写入文件
	err = l.WriteFile(&file, path)
	if err != nil {
		return nil, err
	}

	//提取音频
	if dataType == "1" {
		path, err = utils.ExtractMp3(path, dir_path)
		if err != nil {
			return nil, err
		}
	}

	query := map[string]string{
		"client_version": cons.Client_version,
		"package":        cons.Package,
		"user_id":        cons.User_id,
		"system_time":    strconv.Itoa(int(time.Now().Unix())),
		"engineid":       "fileasrrecorder",
	}
	headers, signParams := utils.GenSignHeaders(cons.AppId, cons.AppKey, "POST", code.CreateUri, query)
	headers["Content-Type"] = "application/json; charset=UTF-8"
	rurl := fmt.Sprintf("%s%s?%s", code.Domain, code.CreateUri, signParams)
	xsid := uuid.New().String()

	//创建音频请求
	query = map[string]string{
		"audio_type":  "auto",
		"x-sessionId": xsid,
		"slice_num":   "1",
	}
	response, err := utils.PostBodyRequest(rurl, query, headers)
	if err != nil {
		return nil, err
	}
	var createResponse types.CreateResponse
	err = json.Unmarshal([]byte(response), &createResponse)
	if err != nil {
		return nil, err
	}

	//文件上传
	query = map[string]string{
		"audio_id":    createResponse.Data.AudioID,
		"x-sessionId": xsid,
		"slice_index": "0",
	}
	uploadQueryParams := utils.GenCanonicalQueryString(query)
	rurl = fmt.Sprintf("%s%s?%s&%s", code.Domain, code.UploadUri, signParams, uploadQueryParams)
	response, err = utils.PostFormDataRequest(path, filename, rurl, headers)
	if err != nil {
		return nil, err
	}
	var uploadResponse types.UploadResponse
	err = json.Unmarshal([]byte(response), &uploadResponse)
	if err != nil {
		return nil, err
	}

	//开始转写
	query = map[string]string{
		"audio_id":    createResponse.Data.AudioID,
		"x-sessionId": xsid,
	}
	rurl = fmt.Sprintf("%s%s?%s", code.Domain, code.TransUri, signParams)
	response, err = utils.PostBodyRequest(rurl, query, headers)
	if err != nil {
		return nil, err
	}
	var transResponse types.TransResponse
	err = json.Unmarshal([]byte(response), &transResponse)
	if err != nil {
		return nil, err
	}

	//查询转写进度
	headers["Content-Type"] = "application/json; charset=UTF-8"
	rurl = fmt.Sprintf("%s%s?%s", code.Domain, code.ProgressUri, signParams)
	var progressResponse types.ProgressResponse
	maxDuration := 15 * time.Minute //最大查询时间
	ticker := time.NewTicker(5 * time.Second)
	timeAfter := time.After(maxDuration)
	//循环查询
	progress := 0

ForEnd:
	for {
		select {
		case <-ticker.C:
			response, err = utils.PostBodyRequest(rurl, map[string]string{
				"task_id":     transResponse.Data.TaskID,
				"x-sessionId": xsid,
			}, headers)
			if err != nil {
				l.Logger.Error(err)
				return
			}
			err = json.Unmarshal([]byte(response), &progressResponse)
			if err != nil {
				l.Logger.Error(err)
				return
			}
			if progressResponse.Data.Progress == 100 {
				progress = 100
				l.Logger.Info("progress 100")
				ticker.Stop()
				break ForEnd
			}
		case <-timeAfter:
			l.Logger.Error("转写超时，后端断开连接，请压缩文件")
			ticker.Stop()
			return
		}
	}

	//当progress为100时查询结果
	var resultResponse types.ResultResponse
	if progress == 100 {
		headers["Content-Type"] = "application/json; charset=UTF-8"
		rurl := fmt.Sprintf("%s%s?%s", code.Domain, code.ResultUri, signParams)
		response, err = utils.PostBodyRequest(rurl, map[string]string{
			"task_id":     transResponse.Data.TaskID,
			"x-sessionId": xsid,
		}, headers)
		if err != nil {
			l.Logger.Error(err)
			return
		}
		err = json.Unmarshal([]byte(response), &resultResponse)
		if err != nil {
			l.Logger.Error(err)
			return
		}
	} else {
		l.Logger.Error("未知错误")
		return
	}

	//查询出来的结果，进行问答
	query = map[string]string{"requestId": uuid.New().String()}
	headers, signParams = utils.GenSignHeaders(cons.AppId, cons.AppKey, "POST", code.GptUri, query)
	headers["Content-Type"] = "application/json"

	var systemPrompt strings.Builder
	systemPrompt.WriteString("请你作为一名文学方面的专家,结合以下一段文本内容回答我所提出的问题,回答需要遵循指定的格式。")
	systemPrompt.WriteString("<Context>")
	for _, v := range resultResponse.Data.Result {
		systemPrompt.WriteString(v.OneBest)
	}
	systemPrompt.WriteString("</Context>")

	sessionId := uuid.New().String()
	rurl = fmt.Sprintf("%s%s?%s", code.Domain, code.GptUri, signParams)
	dataAsk := map[string]any{
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": systemPrompt.String(),
			},
			{
				"role":    "user",
				"content": "请帮我提炼出关键词(不多于20个关键词),回答的格式为\"{\"answer\": [\"关键词1\",\"关键词2\",...\"关键词n\"]}\",如果不需要额外回复，则对应字段输出空。",
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
				"content": "请对内容进行总结得出一个简略的全文概要(不超过512个字),回答的格式为\"{\"answer\": \"给用户的回复\"}\",如果不需要额外回复，则对应字段输出空。",
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
				"content": "请对内容归纳章节(尽可能将多句话归纳为一个章节),包含章节标题、每个章节概括性的阐述(每个章节不超过128个字),回答的格式为\"{\"answer\":[{\"title\":\"返回的标题\",\"desc\":\"返回的阐述\"},{\"title\":\"返回的标题\",\"desc\":\"返回的阐述\"}]}\",如果不需要额外回复，则对应字段输出空。",
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
				"content": "请从内容中归纳出几个要点问题,并根据要点问题和内容分析出答案,返回的格式为\"{\"answer\":[{\"ques\":\"返回的问题\",\"ans\":\"问题的答案\"},{\"ques\":\"返回的问题\",\"ans\":\"问题的答案\"}]}\",如果不需要额外回复，则对应字段输出空。",
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
	d_type, _ := strconv.Atoi(dataType)
	var transText []map[string]string
	for _, v := range resultResponse.Data.Result {
		if v.OneBest != "" {
			transText = append(transText, map[string]string{
				"timestamp": utils.IntToMSS(v.Bg),
				"content":   v.OneBest,
			})
		}
	}
	transTextbyte, err := json.Marshal(transText)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	innerText := types.InnerText{
		RecordTimeStamp: filename,
		SourceLanguage:  source_language,
		DestLanguage:    dest_language,
		Content:         "",
		SaveTime:        "已保存于 " + time.Now().Format("15:04:05"),
	}
	innerTextbyte, err := json.Marshal(innerText)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	saveRecord := video.UploadRequest{
		Path:      fmt.Sprintf("%s/%s", urlPath, filename),
		UserId:    uint64(u_id),
		DataType:  uint32(d_type),
		Keyword:   keyword,
		Summary:   summary,
		Scanning:  scanning,
		Review:    review,
		TransText: string(transTextbyte),
		Innertext: string(innerTextbyte),
		Title:     filename,
	}
	res, err := l.svcCtx.VideoRpc.UploadVideo(l.ctx, &saveRecord)
	if err != nil {
		logx.Error(err)
		return
	}

	return &types.UploadVideoReponse{
		Code:     int(res.Code),
		Msg:      res.Msg,
		RecordId: int(res.RecordId),
	}, nil
}

func (l *UploadVideoLogic) CheckAndCreate(path, userId, hash, dataType string) (string, string, error) {
	//检查用户文件夹是否存在
	var urlPath string
	switch dataType {
	case "0":
		path = strings.Join([]string{path, "voice", userId, hash}, "/")
		urlPath = strings.Join([]string{cons.RequestIp, "voice", userId, hash}, "/")
	case "1":
		path = strings.Join([]string{path, "video", userId, hash}, "/")
		urlPath = strings.Join([]string{cons.RequestIp, "video", userId, hash}, "/")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 目录不存在，创建目录
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return "", "", err
		}
	}
	return path, urlPath, nil
}

func (l *UploadVideoLogic) WriteFile(file *multipart.File, path string) error {
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
