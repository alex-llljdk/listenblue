package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"zjhx.com/live/api/internal/svc"
	"zjhx.com/live/api/internal/types"
	"zjhx.com/live/code"
	"zjhx.com/live/gmodel"
	"zjhx.com/live/rpc/live"
	"zjhx.com/pkg/cons"
	"zjhx.com/pkg/utils"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type ClosePublishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClosePublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClosePublishLogic {
	return &ClosePublishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClosePublishLogic) ClosePublish(req *types.ClosePublishRequest) (resp *types.ClosePublishResponse, err error) {
	// todo: add your logic here and delete this line

	//防止obs提前中断处理逻辑
	l.ctx = context.Background()
	//关闭房间，删除redis数据，中断推流
	res, err := l.svcCtx.LiveRpc.ClosePublish(l.ctx, &live.ClosePublishRequest{
		EncRoomNumber: req.Name,
	})
	if err != nil {
		return nil, err
	}

	//判断是否需要保存
	var roomInfo gmodel.LiveRoom
	err = json.Unmarshal([]byte(res.Res), &roomInfo)
	if err != nil {
		return nil, err
	}

	if roomInfo.IsRecord {
		l.SaveToDB(req, roomInfo)
	}
	l.ctx.Done()
	return &types.ClosePublishResponse{
		Code: int(res.Code),
		Msg:  res.Msg,
	}, nil
}

func (l *ClosePublishLogic) SaveToDB(req *types.ClosePublishRequest, roomInfo gmodel.LiveRoom) {
	//将flv视频转为生成MP3视频
	path := fmt.Sprintf("%s/%s.flv", code.LiveDataUrl, req.Name)
	urlPath := fmt.Sprintf("%s/%s.flv", code.LiveNetUrl, req.Name)
	_, err := utils.ExtractMp3(path, code.LiveDataUrl)
	if err != nil {
		logx.Error(err)
		return
	}

	//检查文件是否大于5MB，大于则进行切片
	slice_num, err := utils.SplitFile(code.LiveDataUrl, req.Name)
	if err != nil {
		logx.Error(err)
		return
	}

	// 参数设置
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

	// 创建音频请求
	query = map[string]string{
		"audio_type":  "auto",
		"x-sessionId": xsid,
		"slice_num":   fmt.Sprintf("%d", slice_num),
	}
	response, err := utils.PostBodyRequest(rurl, query, headers)
	if err != nil {
		logx.Error(err)
		return
	}
	var createResponse types.CreateResponse
	err = json.Unmarshal([]byte(response), &createResponse)
	if err != nil {
		logx.Error(err)
		return
	}

	for i := int64(0); i < slice_num; i++ {
		var uploadPath, fn string
		if slice_num == 1 {
			uploadPath = fmt.Sprintf("%s/%s.mp3", code.LiveDataUrl, req.Name)
			fn = fmt.Sprintf("%s.mp3", req.Name)
		} else {
			uploadPath = fmt.Sprintf("%s/%s_%d.mp3", code.LiveDataUrl, req.Name, i)
			fn = fmt.Sprintf("%s_%d.mp3", req.Name, i)
		}
		query = map[string]string{
			"audio_id":    createResponse.Data.AudioID,
			"x-sessionId": xsid,
			"slice_index": fmt.Sprintf("%d", i),
		}
		uploadQueryParams := utils.GenCanonicalQueryString(query)
		rurl = fmt.Sprintf("%s%s?%s&%s", code.Domain, code.UploadUri, signParams, uploadQueryParams)
		response, err = utils.PostFormDataRequest(uploadPath, fn, rurl, headers)
		if err != nil {
			logx.Error(err)
			return
		}
		var uploadResponse types.UploadResponse
		err = json.Unmarshal([]byte(response), &uploadResponse)

		if err != nil {
			logx.Error(err)
			return
		}
	}

	// 开始转写
	query = map[string]string{
		"audio_id":    createResponse.Data.AudioID,
		"x-sessionId": xsid,
	}
	headers["Content-Type"] = "application/json; charset=UTF-8"
	rurl = fmt.Sprintf("%s%s?%s", code.Domain, code.TransUri, signParams)
	response, err = utils.PostBodyRequest(rurl, query, headers)
	if err != nil {
		logx.Error(err)
		return
	}
	var transResponse types.TransResponse
	err = json.Unmarshal([]byte(response), &transResponse)
	if err != nil {
		logx.Error(err)
		return
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
			fmt.Println("progereee-------------", response)

			if err != nil {
				logx.Error(err)
				return
			}
			err = json.Unmarshal([]byte(response), &progressResponse)

			if err != nil {
				logx.Error(err)
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

	//当progress为100时查询结果，并写入数据库
	var resultResponse types.ResultResponse
	if progress == 100 {
		headers["Content-Type"] = "application/json; charset=UTF-8"
		rurl := fmt.Sprintf("%s%s?%s", code.Domain, code.ResultUri, signParams)
		response, err = utils.PostBodyRequest(rurl, map[string]string{
			"task_id":     transResponse.Data.TaskID,
			"x-sessionId": xsid,
		}, headers)
		if err != nil {
			logx.Error(err)
			return
		}
		err = json.Unmarshal([]byte(response), &resultResponse)
		if err != nil {
			logx.Error(err)
			return
		}
	} else {
		logx.Error(err)
		return
	}

	// 打印解码后的结构体内容
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
				"content": "请问内容是中文还是英文,你只需要回答\"中文\"或\"英文\",不需要额外回复。",
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
		return
	}
	err = json.Unmarshal([]byte(response), &ChatResponse)
	if err != nil {
		l.Logger.Error(err)
		return
	}
	Isch := ChatResponse.Data.Content
	fmt.Println(response)

	ChatResponse = types.ChatResponse{}
	dataAsk = map[string]any{
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
	// 角色预设
	response, err = utils.PostBodyRequest(rurl, dataAsk, headers)
	response = strings.Replace(response, " ", "", -1)
	response = strings.Replace(response, "\n", "", -1)
	if err != nil {
		logx.Error(err)
		return
	}
	err = json.Unmarshal([]byte(response), &ChatResponse)
	if err != nil {
		logx.Error(err)
		return
	}
	keyword := ChatResponse.Data.Content
	fmt.Println(response)

	ChatResponse = types.ChatResponse{}
	// 全文概要
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
		logx.Error(err)
		return
	}
	err = json.Unmarshal([]byte(response), &ChatResponse)
	if err != nil {
		logx.Error(err)
		return
	}
	summary := ChatResponse.Data.Content
	fmt.Println(response)
	ChatResponse = types.ChatResponse{}
	// 章节速览 总结出一个标题，并返回对应在那句话
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
		logx.Error(err)
		return
	}
	err = json.Unmarshal([]byte(response), &ChatResponse)
	if err != nil {
		logx.Error(err)
		return
	}
	scanning := ChatResponse.Data.Content
	fmt.Println(response)
	ChatResponse = types.ChatResponse{}
	// 要点回顾
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
		logx.Error(err)
		return
	}
	err = json.Unmarshal([]byte(response), &ChatResponse)
	if err != nil {
		logx.Error(err)
		return
	}
	review := ChatResponse.Data.Content
	fmt.Println(response)

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
		return
	}
	var sourceLanguage, destLanguage string
	if Isch == "中文" {
		sourceLanguage = "zh-CHS"
		destLanguage = "en"
	} else {
		sourceLanguage = "en"
		destLanguage = "zh-CHS"
	}
	innerText := types.InnerText{
		RecordTimeStamp: req.Name,
		SourceLanguage:  sourceLanguage,
		DestLanguage:    destLanguage,
		Content:         "",
		SaveTime:        "已保存于 " + time.Now().Format("15:04:05"),
	}
	innerTextbyte, err := json.Marshal(innerText)
	if err != nil {
		logx.Error(err)
		return
	}

	//最后将所有内容写入数据库中
	saveRecord := &live.SaveRecordRequest{
		Path:      urlPath,
		Title:     roomInfo.RoomTitle,
		UserId:    uint64(roomInfo.UserID),
		DataType:  uint32(1),
		Keyword:   keyword,
		Summary:   summary,
		Scanning:  scanning,
		Review:    review,
		TransText: string(transTextbyte),
		InnerText: string(innerTextbyte),
	}
	_, err = l.svcCtx.LiveRpc.SaveToRecord(l.ctx, saveRecord)

	if err != nil {
		logx.Error(err)
		return
	}
}
