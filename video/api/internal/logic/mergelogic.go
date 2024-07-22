package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

type MergeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMergeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MergeLogic {
	return &MergeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MergeLogic) Merge(req types.MergeRequest) (resp *types.MergeReponse, err error) {
	// todo: add your logic here and delete this line
	//检查文件夹是否存在
	var path string
	var urlPath string
	switch req.DataType {
	case "0":
		path = strings.Join([]string{l.svcCtx.Config.FilePath, "voice", req.UserId, req.Md5 + "_" + req.Total}, "/")
		urlPath = strings.Join([]string{cons.RequestIp, "voice", req.UserId, req.Md5 + "_" + req.Total}, "/")
	case "1":
		path = strings.Join([]string{l.svcCtx.Config.FilePath, "video", req.UserId, req.Md5 + "_" + req.Total}, "/")
		urlPath = strings.Join([]string{cons.RequestIp, "video", req.UserId, req.Md5 + "_" + req.Total}, "/")
	}
	isExits, err := utils.PathExists(path)
	if err != nil {
		return nil, err
	}
	//如果文件夹不存在，则返回不存在
	if !isExits {
		return nil, errors.New("文件夹不存在，请重新上传")
	}
	//存在则检查是否上传完毕，如果文件不完整则返回最小下标，完整则合并
	t, err := strconv.Atoi(req.Total)
	if err != nil {
		return nil, err
	}

	filename_split := strings.Split(req.FileName, ".")
	for i := 0; i < t; i++ {
		isExits, err = utils.PathExists(fmt.Sprintf("%s/%s_%d.%s", path, filename_split[0], i, filename_split[1]))
		if err != nil {
			return nil, err
		}
		if !isExits {
			return &types.MergeReponse{
				Code: cons.ServiceOKCode,
				Msg:  cons.ServiceOKMsg,
				Idx:  i,
			}, nil
		}
	}

	//合并视频
	mergeFile, err := os.Create(fmt.Sprintf("%s/%s", path, req.FileName))
	if err != nil {
		return nil, err
	}

	multiWriter := io.MultiWriter(mergeFile)
	for i := 0; i < t; i++ {
		err = l.MergeFile(fmt.Sprintf("%s/%s_%d.%s", path, filename_split[0], i, filename_split[1]), &multiWriter)
		if err != nil {
			return nil, err
		}
	}
	mergeFile.Close()

	//提取音频
	if filename_split[1] == "mp4" {
		_, err = utils.ExtractMp3(fmt.Sprintf("%s/%s", path, req.FileName), path)
		if err != nil {
			return nil, err
		}
	}

	//判断提取的音频是否大于5MB，如果是则分片
	slice_num, err := utils.SplitFile(path, filename_split[0])
	if err != nil {
		return nil, err
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
		return nil, err
	}
	var createResponse types.CreateResponse
	err = json.Unmarshal([]byte(response), &createResponse)
	if err != nil {
		return nil, err
	}

	for i := int64(0); i < slice_num; i++ {
		var uploadPath, fn string
		if slice_num == 1 {
			uploadPath = fmt.Sprintf("%s/%s.mp3", path, filename_split[0])
			fn = fmt.Sprintf("%s.mp3", filename_split[0])
		} else {
			uploadPath = fmt.Sprintf("%s/%s_%d.mp3", path, filename_split[0], i)
			fn = fmt.Sprintf("%s_%d.mp3", filename_split[0], i)
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
			return nil, err
		}
		var uploadResponse types.UploadResponse
		err = json.Unmarshal([]byte(response), &uploadResponse)

		if err != nil {
			return nil, err
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
			fmt.Println("progereee-------------", response)

			if err != nil {
				l.Logger.Error(err)
				return nil, err
			}
			err = json.Unmarshal([]byte(response), &progressResponse)

			if err != nil {
				l.Logger.Error(err)
				return nil, err
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
			return nil, err
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
			l.Logger.Error(err)
			return nil, err
		}
		err = json.Unmarshal([]byte(response), &resultResponse)
		if err != nil {
			l.Logger.Error(err)
			return nil, err
		}
	} else {
		l.Logger.Error("未知错误")
		return nil, err
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

	u_id, _ := strconv.Atoi(req.UserId)
	d_type, _ := strconv.Atoi(req.DataType)

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
		RecordTimeStamp: req.FileName,
		SourceLanguage:  req.SourceLanguage,
		DestLanguage:    req.DestLanguage,
		Content:         "",
		SaveTime:        "已保存于 " + time.Now().Format("15:04:05"),
	}
	innerTextbyte, err := json.Marshal(innerText)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	saveRecord := video.UploadRequest{
		Path:      fmt.Sprintf("%s/%s", urlPath, req.FileName),
		UserId:    uint64(u_id),
		DataType:  uint32(d_type),
		Keyword:   keyword,
		Summary:   summary,
		Scanning:  scanning,
		Review:    review,
		TransText: string(transTextbyte),
		Innertext: string(innerTextbyte),
		Title:     req.FileName,
	}
	res, err := l.svcCtx.VideoRpc.UploadVideo(l.ctx, &saveRecord)
	if err != nil {
		logx.Error(err)
		return
	}
	return &types.MergeReponse{
		Code:     int(res.Code),
		Msg:      res.Msg,
		Idx:      -1,
		RecordId: int(res.RecordId),
	}, nil
}

func (l *MergeLogic) MergeFile(path string, writer *io.Writer) (err error) {
	sliceFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer sliceFile.Close()
	_, err = io.Copy(*writer, sliceFile)
	if err != nil {
		return err
	}
	return nil
}
