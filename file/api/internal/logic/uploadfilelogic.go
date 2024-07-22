package logic

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"zjhx.com/file/api/internal/svc"
	"zjhx.com/file/api/internal/types"
	"zjhx.com/file/code"
	"zjhx.com/file/rpc/file"
	"zjhx.com/pkg/cons"
	"zjhx.com/pkg/utils"
)

type UploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type OCRText struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type GPTResp struct {
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	KeyPoints string `json:"keypoints"`
	SpeedRead string `json:"speedread"`
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadFileLogic) UploadFile(filename, userId, dataType, hash string, files multipart.File) (resp *types.UploadFileResponse, err error) {
	// todo: add your logic here and delete this line
	//dataType 2是pdf,3是word,4是img,5是html
	//检查创建目录
	dir_path, saveDBDir, err := l.CheckAndCreate(l.svcCtx.Config.FilePath, userId, hash, dataType)
	if err != nil {
		return nil, err
	}

	path := strings.Join([]string{dir_path, filename}, "/")
	saveDB_Path := strings.Join([]string{saveDBDir, filename}, "/")
	// fmt.Println(dir_path)
	// fmt.Println(path)
	// fmt.Println("---------------------------")
	//写入文件
	err = l.WriteFile(&files, path)
	if err != nil {
		return nil, err
	}

	exportimgpath := strings.Join([]string{dir_path, cons.ExportPrefix}, "/")
	//获取文件基本名不包括拓展名
	baseName := strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))
	// dataType=2 pdf dataType=3 word dataType=4 img
	if dataType == "3" {
		fmt.Println(path)
		_, err = utils.ConvertToPDF(path, dir_path)
		if err != nil {
			return nil, err
		}
	} else if dataType == "4" {
		//img 2 pdf
		//将图片转成pdf保存
		utils.Img2pdf(dir_path, strings.Join([]string{dir_path, "img2pdf.pdf"}, "/"))

	}
	var pdf_path string
	var saveDB_pdfPath string //保存到数据库的pdf路径(http路径)
	if dataType == "2" {
		pdf_path = path
		saveDB_pdfPath = saveDB_Path
	} else if dataType == "3" {
		//获取文件名
		pdf_path = strings.Join([]string{dir_path, baseName + ".pdf"}, "/")
		saveDB_pdfPath = strings.Join([]string{saveDBDir, baseName + ".pdf"}, "/")
	} else if dataType == "4" {
		pdf_path = strings.Join([]string{dir_path, "img2pdf.pdf"}, "/")
		saveDB_pdfPath = strings.Join([]string{saveDBDir, "img2pdf.pdf"}, "/")
	}

	fmt.Println("*****************************************")
	fmt.Println(saveDB_pdfPath)
	fmt.Println("*****************************************")

	if dataType != "4" {
		//读取pdf文件，调用pdf2png.exe将pdf导出为png图像
		cmd := exec.Command(cons.PdfExePath, pdf_path, exportimgpath)
		cmderr := cmd.Run()
		if cmderr != nil {
			fmt.Println("pdf转换成图像失败: %w", cmderr)
			return
		}
	}

	// 定义OCR提取的完整文本变量
	var OCRText strings.Builder
	// var OCRText1 string
	OCRText.WriteString("请你作为一名文学方面的专家,结合以下一段文本内容回答我所提出的问题,我将json的格式输入内容,内容结构为{\"page_n\":\"content\"},其中page_n表示为pdf的页码,content为具体内容,每个json表示一页pdf的内容,回答需要遵循指定的格式。")
	OCRText.WriteString("<Context>")
	pageidx := 1
	//循环读取保存好的图片
	// 遍历文件夹中所有导出的图片
	filewalkerr := filepath.Walk(dir_path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // 处理错误
		}
		// 检查文件是否为PNG格式
		if !info.IsDir() && (filepath.Ext(path) == ".png" || filepath.Ext(path) == ".jpg") {
			// fmt.Println(path)
			ocrResp := l.DoOCRRequest(path)
			//拼接文本信息
			OCRText.WriteString(fmt.Sprintf("{\"page_%d\":%s}", pageidx, l.ParseOCRResp(ocrResp)))
			// OCRText1 += ocrResp
			pageidx++
		}
		return nil
	})
	if filewalkerr != nil {
		fmt.Println("Error file walking the path:", err)
		return nil, filewalkerr
	}
	OCRText.WriteString("</Context>")
	// fmt.Println(OCRText.String())
	// 请求大模型
	GPTResp, gpterr := l.DoGPTRequest(OCRText.String())
	if gpterr != nil {
		return nil, gpterr
	}
	// var title map[string]interface{}
	// err = json.Unmarshal([]byte(GPTResp.Title), &tmp)
	// if err != nil {
	// 	return nil, err
	// }
	//调用rpc服务请求保存数据
	// fmt.Println(GPTResp)
	userid, _ := strconv.Atoi(userId)
	datatype, _ := strconv.Atoi(dataType)
	res, err := l.svcCtx.FileRpc.UploadFile(l.ctx, &file.FileRecord{
		Path:     saveDB_pdfPath,
		Title:    GPTResp.Title,
		UserId:   uint64(userid),
		DataType: uint32(datatype),
		Keyword:  GPTResp.KeyPoints,
		Summary:  GPTResp.Summary,
		Scanning: GPTResp.SpeedRead,
	})
	if err != nil {
		return nil, err
	}
	return &types.UploadFileResponse{
		Code:     int(res.Code),
		Msg:      res.Msg,
		RecordId: int(res.RecordId),
	}, nil
}

func (l *UploadFileLogic) UploadHtml(url, userId, dataType string) (resp *types.UploadFileResponse, err error) {
	hasher := md5.New()
	hasher.Write([]byte(url))
	hashByte := hasher.Sum(nil)
	hash := hex.EncodeToString(hashByte)

	dir_path, saveDBDir, err := l.CheckAndCreate(l.svcCtx.Config.FilePath, userId, hash, dataType)
	if err != nil {
		return nil, err
	}
	pdf_path := strings.Join([]string{dir_path, "img2pdf.pdf"}, "/")
	saveDB_pdfPath := strings.Join([]string{saveDBDir, "img2pdf.pdf"}, "/")
	//将 url的内容转换成图像并保存
	utils.Html2Img(url, strings.Join([]string{dir_path, "html2img.png"}, "/"))
	utils.Img2pdf(dir_path, pdf_path)
	// 定义OCR提取的完整文本变量
	var OCRText strings.Builder
	// var OCRText1 string
	OCRText.WriteString("请你作为一名文学方面的专家,结合以下一段文本内容回答我所提出的问题,我将json的格式输入内容,内容结构为{\"page_n\":\"content\"},其中page_n表示为pdf的页码,content为具体内容,每个json表示一页pdf的内容,回答需要遵循指定的格式。")
	OCRText.WriteString("<Context>")
	pageidx := 1
	//循环读取保存好的图片
	// 遍历文件夹中所有导出的图片
	filewalkerr := filepath.Walk(dir_path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // 处理错误
		}
		// 检查文件是否为PNG格式
		if !info.IsDir() && (filepath.Ext(path) == ".png" || filepath.Ext(path) == ".jpg") {
			// fmt.Println(path)
			ocrResp := l.DoOCRRequest(path)
			//拼接文本信息
			OCRText.WriteString(fmt.Sprintf("{\"page_%d\":%s}", pageidx, l.ParseOCRResp(ocrResp)))
			// OCRText1 += ocrResp
			pageidx++
		}
		return nil
	})
	if filewalkerr != nil {
		fmt.Println("Error file walking the path:", err)
		return nil, filewalkerr
	}
	OCRText.WriteString("</Context>")
	fmt.Println("*****************************************")
	fmt.Println(saveDB_pdfPath)
	fmt.Println("*****************************************")
	// 请求大模型
	GPTResp, gpterr := l.DoGPTRequest(OCRText.String())
	if gpterr != nil {
		return nil, gpterr
	}
	//调用rpc服务请求保存数据
	userid, _ := strconv.Atoi(userId)
	datatype, _ := strconv.Atoi(dataType)
	res, err := l.svcCtx.FileRpc.UploadFile(l.ctx, &file.FileRecord{
		Path:     saveDB_pdfPath,
		Title:    GPTResp.Title,
		UserId:   uint64(userid),
		DataType: uint32(datatype),
		Keyword:  GPTResp.KeyPoints,
		Summary:  GPTResp.Summary,
		Scanning: GPTResp.SpeedRead,
	})
	if err != nil {
		return nil, err
	}

	return &types.UploadFileResponse{
		Code:     int(res.Code),
		Msg:      res.Msg,
		RecordId: int(res.RecordId),
	}, nil

}

func (l *UploadFileLogic) CheckAndCreate(path, userId, hash, dataType string) (string, string, error) {
	saveDBPath := cons.RequestIp
	//检查用户文件夹是否存在
	switch dataType {
	case "2":
		path = strings.Join([]string{path, "pdf", userId, hash}, "/")
		saveDBPath = strings.Join([]string{saveDBPath, "pdf", userId, hash}, "/")
	case "3":
		path = strings.Join([]string{path, "word", userId, hash}, "/")
		saveDBPath = strings.Join([]string{saveDBPath, "word", userId, hash}, "/")
	case "4":
		path = strings.Join([]string{path, "img", userId, hash}, "/")
		saveDBPath = strings.Join([]string{saveDBPath, "img", userId, hash}, "/")
	case "5":
		path = strings.Join([]string{path, "html", userId, hash}, "/")
		saveDBPath = strings.Join([]string{saveDBPath, "html", userId, hash}, "/")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 目录不存在，创建目录
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return "", "", err
		}
	}
	return path, saveDBPath, nil
}

func (l *UploadFileLogic) WriteFile(file *multipart.File, path string) error {
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

// 请求OCR接口
func (l *UploadFileLogic) DoOCRRequest(imgPath string) string {
	// 读取图片文件
	imgFile, err := os.Open(imgPath)
	if err != nil {
		fmt.Println("读取pdf图片失败")
		return ""
	}
	defer imgFile.Close()
	// 读取文件内容
	bImage, err := ioutil.ReadAll(imgFile)
	if err != nil {
		fmt.Println("读取pdf图片内容失败")
		return ""
	}
	// 将图片内容编码为base64字符串
	image := base64.StdEncoding.EncodeToString(bImage)
	// 构建ocr接口请求体
	post_data := map[string]string{
		"image":      image,
		"pos":        "0",
		"businessid": "1990173156ceb8a09eee80c293135279",
	}
	params := make(map[string]string)

	// 循环发起请求
	// 生成签名
	headers, _ := utils.GenSignHeaders(cons.AppId, cons.AppKey, "POST", code.OCR_Uri, params)
	full_url := "http://api-ai.vivo.com.cn" + code.OCR_Uri

	ocrResp, ocrerr := utils.PostUrlencodedRequest(full_url, post_data, headers)
	if ocrerr != nil {
		fmt.Println("请求失败----------" + string(ocrerr.Error()))
		return ""
	}
	return ocrResp
}

func (l *UploadFileLogic) DoGPTRequest(OCRText string) (gptresp *GPTResp, err error) {
	params := make(map[string]string)
	params["requestId"] = uuid.New().String()
	headers, signParams := utils.GenSignHeaders(cons.AppId, cons.AppKey, "POST", code.GptUri, params)
	headers["Content-Type"] = "application/json"
	rurl := fmt.Sprintf("%s%s?%s", code.Domain, code.GptUri, signParams)
	var data map[string]interface{}

	sessionId := uuid.New().String()
	//对内容起一个标题
	dataAsk := map[string]any{
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": OCRText,
			},
			{
				"role":    "user", //
				"content": "请对内容进行总结得出一个标题,回答的格式为\"{\"answer\":\"给用户的回复\"}\",如果不需要额外回复或者没有合适工具，则对应字段输出空。",
			},
		},
		"model":     "vivo-BlueLM-TB",
		"sessionId": sessionId,
		"extra": map[string]string{
			"temperature": "0.9",
		},
	}
	response, err := utils.PostBodyRequest(rurl, dataAsk, headers)
	response = strings.Replace(response, " ", "", -1)
	response = strings.Replace(response, "\n", "", -1)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	reserr := json.Unmarshal([]byte(response), &data)
	if reserr != nil {
		fmt.Println("JSON文本解析失败")
		return nil, reserr
	}
	TitleText := data["data"].(map[string]interface{})["content"]

	//全文概要
	dataAsk = map[string]any{
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": OCRText,
			},
			{
				"role":    "user",
				"content": "请对内容进行总结得出一个简略的全文概要(不超过512个字),回答的格式为\"{\"answer\":\"给用户的回复\"}\",如果不需要额外回复或者没有合适工具，则对应字段输出空。",
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
	// fmt.Println(response)
	reserr = json.Unmarshal([]byte(response), &data)
	if reserr != nil {
		fmt.Println("JSON文本解析失败")
	}
	SummaryText := data["data"].(map[string]interface{})["content"]

	//关键要点
	dataAsk = map[string]any{
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": OCRText,
			},
			{
				"role":    "user",
				"content": "请对内容进行总结其中的关键要点,返回每个关键要点的阐述(每个关键要点不超过30个字),回答的格式为\"{\"answer\":[{\"index\":\"返回的编号\",\"desc\":\"返回的关键要点\"},{\"index\":\"返回的编号\",\"desc\":\"返回的关键要点\"}]}\",如果不需要额外回复或者没有合适工具，则对应字段输出空。",
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
	// fmt.Println(response)
	reserr = json.Unmarshal([]byte(response), &data)
	if reserr != nil {
		fmt.Println("JSON文本解析失败")
		return nil, reserr
	}
	KeyPointsText := data["data"].(map[string]interface{})["content"]

	//文档速读
	dataAsk = map[string]any{
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": OCRText,
			},
			{
				"role":    "user",
				"content": "请对每个json中的内容进行归纳总结,返回page_n、总结的标题、总结的内容(每一页总结的内容不超过128个字),回答的格式为\"{\"answer\":[{\"index\":\"返回的编号\",\"title\":\"返回的标题\",\"desc\":\"返回的内容\"},{\"index\":\"返回的编号\",\"title\":\"返回的标题\",\"desc\":\"返回的内容\"}]}\",如果不需要额外回复或者没有合适工具，则对应字段输出空。",
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
	// fmt.Println(response)
	reserr = json.Unmarshal([]byte(response), &data)
	if reserr != nil {
		fmt.Println("JSON文本解析失败")
	}
	SpeedReading := data["data"].(map[string]interface{})["content"]

	// fmt.Println("文章标题如下所示：    ")
	// fmt.Println(TitleText)
	// fmt.Println("全文概要如下所示：    ")
	// fmt.Println(SummaryText)
	// fmt.Println("关键要点如下所示：    ")
	// fmt.Println(KeyPointsText)
	// fmt.Println("文档速读如下所示：    ")
	// fmt.Println(SpeedReading)
	// var GPTResp strings.Builder
	// GPTResp.WriteString(TitleText.(string))
	// fmt.Println(data["data"])
	return &GPTResp{
		Title:     TitleText.(string),
		Summary:   SummaryText.(string),
		KeyPoints: KeyPointsText.(string),
		SpeedRead: SpeedReading.(string),
	}, nil
}

// 将每次请求的OCR结果中提取到的字符串进行解析出来返回
func (l *UploadFileLogic) ParseOCRResp(jsonOCRRespStr string) string {
	// 将JSON字符串解析到结构体中
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonOCRRespStr), &data)
	if err != nil {
		fmt.Println("解析JSON时发生错误:", err)
		return ""
	}

	wordsArray, ok := data["result"].(map[string]interface{})["words"].([]interface{})
	if !ok {
		fmt.Println("无法获取words数组")
		return ""
	}

	// 拼接words数组中的所有内容
	var fullString string
	for _, wordItem := range wordsArray {
		if word, ok := wordItem.(map[string]interface{})["words"].(string); ok {
			fullString += word
		}
	}
	return fullString
}
