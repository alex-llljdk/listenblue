package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
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

type UploadUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadUrlLogic {
	return &UploadUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadUrlLogic) UploadUrl(req *types.UploadUrlRequest) (resp *types.UploadUrlResponse, err error) {
	// todo: add your logic here and delete this line
	//获取地址的hash码
	hasher := md5.New()
	hasher.Write([]byte(req.Url))
	hashByte := hasher.Sum(nil)
	hash := hex.EncodeToString(hashByte)

	//获取文件名和文件类型
	filename, dataType, err := l.checkFile(req.Url)
	if err != nil {
		return nil, err
	}

	//文件下载并写入
	dirpath, urlPath, err := l.CheckAndCreate(l.svcCtx.Config.FilePath, fmt.Sprintf("%d", req.UserId), hash, dataType)
	if err != nil {
		return nil, err
	}

	filepath, err := l.downloadFile(dirpath, filename, req.Url)
	if err != nil {
		return nil, err
	}
	//检查文件是否需要转mp3
	if dataType == "1" {
		_, err = utils.ExtractMp3(filepath, dirpath)
		if err != nil {
			return nil, err
		}
	}

	filename_split := strings.Split(filename, ".")
	// 参数设置
	//判断提取的音频是否大于5MB，如果是则分片
	slice_num, err := utils.SplitFile(dirpath, filename_split[0])
	if err != nil {
		return nil, err
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
			uploadPath = fmt.Sprintf("%s/%s.mp3", dirpath, filename_split[0])
			fn = fmt.Sprintf("%s.mp3", filename_split[0])
		} else {
			uploadPath = fmt.Sprintf("%s/%s_%d.mp3", dirpath, filename_split[0], i)
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
		return nil, err
	}
	err = json.Unmarshal([]byte(response), &ChatResponse)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
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
	//角色预设
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
	var sourceLanguage, destLanguage string
	if Isch == "中文" {
		sourceLanguage = "zh-CHS"
		destLanguage = "en"
	} else {
		sourceLanguage = "en"
		destLanguage = "zh-CHS"
	}

	innerText := types.InnerText{
		RecordTimeStamp: req.Url,
		SourceLanguage:  sourceLanguage,
		DestLanguage:    destLanguage,
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
		UserId:    uint64(req.UserId),
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

	return &types.UploadUrlResponse{
		Code:     int(res.Code),
		Msg:      res.Msg,
		RecordId: int(res.RecordId),
	}, nil
}

func (l *UploadUrlLogic) downloadFile(path, filename, url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download file, status code: %d", resp.StatusCode)
	}

	filepath := fmt.Sprintf("%s/%s", path, filename)
	f, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return "", err
	}

	return filepath, nil
}

func (l *UploadUrlLogic) checkFile(url string) (string, string, error) {
	// 发送GET请求
	resp, err := http.Head(url)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()
	// 检查HTTP响应状态码
	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("failed to get header: %s", resp.Status)
	}
	// 获取Content-Length
	contentLengthStr := resp.Header.Get("Content-Length")
	if contentLengthStr == "" {
		return "", "", fmt.Errorf("failed to get content length")
	}
	// 将Content-Length转换为整数
	contentLength, err := strconv.ParseInt(contentLengthStr, 10, 64)
	if err != nil {
		return "", "", err
	}
	fmt.Println(4)
	// 检查文件大小是否超过限制
	if contentLength > 500*1024*1024 {
		return "", "", fmt.Errorf("file size exceeds the maximum limit of %dMB", 500)
	}

	// 根据Content-Type确定文件后缀名
	fileExtension, datatype := l.getFileExtensionFromContentType(resp.Header.Get("Content-Type"))
	// 如果无法从Content-Type确定后缀名，尝试从URL中提取
	if fileExtension == "" {
		fileExtension, datatype = l.getFileExtensionFromURL(url)
		if fileExtension == "" {
			return "", "", errors.New("暂不支持该链接")
		}
	}

	// 创建文件并写入内容
	fileName := "downloaded_file" + fileExtension
	return fileName, datatype, nil
}

func (l *UploadUrlLogic) getFileExtensionFromContentType(contentType string) (string, string) {
	// 根据Content-Type获取文件后缀名
	switch {
	case strings.Contains(contentType, "audio/mpeg"):
		return ".mp3", "0"
	case strings.Contains(contentType, "video/mp4"):
		return ".mp4", "1"
	default:
		return "", "" // 如果无法识别，返回空字符串
	}
}

func (l *UploadUrlLogic) getFileExtensionFromURL(url string) (string, string) {
	// 从URL中提取文件后缀名
	if strings.HasSuffix(url, ".mp3") {
		return ".mp3", "0"
	} else if strings.HasSuffix(url, ".mp4") {
		return ".mp4", "1"
	}
	return "", "" // 如果URL中没有后缀名，返回空字符串
}

func (l *UploadUrlLogic) CheckAndCreate(path, userId, hash, dataType string) (string, string, error) {
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
