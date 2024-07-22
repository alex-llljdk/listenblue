package utils

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"mime/multipart"

	"io"
	"log"
	"math"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/lampnick/doctron-client-go"
	"github.com/signintech/gopdf"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/zeromicro/go-zero/core/logx"
)

// 判断所给路径文件/文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	//isnotexist来判断，是不是不存在的错误
	if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		return false, nil
	}
	return false, err //如果有错误了，但是不是不存在的错误，所以把这个错误原封不动的返回
}

func InFormat1(target string, str_array []string) bool {
	for _, element := range str_array {
		if target == element {
			return true
		}
	}
	return false
}

// 提取MP3
func ExtractMp3(inputVideoPath, outputDir string) (string, error) {
	_formatArr := []string{"mp4", "flv"}
	_, _file := filepath.Split(inputVideoPath)
	_tmps := strings.Split(_file, ".")
	_ext := _tmps[len(_tmps)-1]
	if !InFormat1(_ext, _formatArr) {
		return "", errors.New("格式不支持")
	}
	_name := _tmps[0]

	_resultVideoPath := filepath.Join(outputDir, fmt.Sprintf("%s.%s", _name, "mp3"))
	fmt.Println(_resultVideoPath, "------------------------")
	err := ffmpeg.Input(inputVideoPath,
		ffmpeg.KwArgs{"loglevel": "quiet"}).
		Output(_resultVideoPath, ffmpeg.KwArgs{"acodec": "libmp3lame"}).
		OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		return "", err
	}
	return _resultVideoPath, nil
}

// 生成随机字符串
func genNonce(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyz0123456789"
	var builder strings.Builder
	for i := 0; i < length; i++ {
		builder.WriteByte(chars[rand.Intn(len(chars))])
	}
	return builder.String()
}

// 构建规范化的查询字符串
func GenCanonicalQueryString(params map[string]string) string {
	query := url.Values{}
	for k, v := range params {
		query.Set(k, v)
	}
	return query.Encode()
}

// 生成签名
func genSignature(appSecret, signingString string) string {
	mac := hmac.New(sha256.New, []byte(appSecret))
	mac.Write([]byte(signingString))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return signature
}

// GenSignHeaders 生成签名的HTTP请求头
func GenSignHeaders(appId, appKey, method, uri string, query map[string]string) (map[string]string, string) {
	//请求方法
	method = strings.ToUpper(method)
	//请求时间
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	//timestamp := "1715444504"
	//随机字符串
	nonce := genNonce(8)
	//nonce := "jqh5jlc9"
	//请求编码后的结果
	//fmt.Println("nonce:    " + nonce + "   timestamp:  "+ timestamp)

	canonicalQueryString := GenCanonicalQueryString(query)

	signedHeadersString := fmt.Sprintf("x-ai-gateway-app-id:%s\nx-ai-gateway-timestamp:%s\nx-ai-gateway-nonce:%s",
		appId, timestamp, nonce)
	signingString := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
		method, uri, canonicalQueryString, appId, timestamp, signedHeadersString)

	signature := genSignature(appKey, signingString)
	return map[string]string{
		"X-AI-GATEWAY-APP-ID":         appId,
		"X-AI-GATEWAY-TIMESTAMP":      timestamp,
		"X-AI-GATEWAY-NONCE":          nonce,
		"X-AI-GATEWAY-SIGNED-HEADERS": "x-ai-gateway-app-id;x-ai-gateway-timestamp;x-ai-gateway-nonce",
		"X-AI-GATEWAY-SIGNATURE":      signature,
	}, canonicalQueryString
}

func XorEncrypt(plaintext []byte, key []byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	for i, b := range plaintext {
		ciphertext[i] = b ^ key[i%len(key)] // 对每个字节进行XOR操作
	}
	return ciphertext
}

// xorDecrypt 使用XOR操作解密数据
func XorDecrypt(ciphertext []byte, key []byte) []byte {
	return XorEncrypt(ciphertext, key) // XOR加密和解密是相同的操作
}

func SplitFile(filedir string, filesplit string) (int64, error) {
	var chunkSize int64 = 1024 * 1024 * 5
	sourceFilePath := fmt.Sprintf("%s/%s.mp3", filedir, filesplit) // 源文件路径                                 // 5MB，每个分片的大小
	// 打开源文件
	file, err := os.Open(sourceFilePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	// 获取文件大小
	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}
	fileSize := fileInfo.Size()
	sliceNum := int64(math.Ceil(float64(fileSize) / float64(chunkSize)))
	if sliceNum == 1 {
		return sliceNum, nil
	}
	for i := int64(0); i < sliceNum; i++ {
		sliceFile, err := os.Create(fmt.Sprintf("%s/%s_%d.mp3", filedir, filesplit, i))
		if err != nil {
			return 0, err
		}
		sectionReader := io.NewSectionReader(file, i*chunkSize, chunkSize)
		_, err = io.Copy(sliceFile, sectionReader)
		sliceFile.Close()
		if err != nil {
			return 0, err
		}
	}
	return sliceNum, nil
}

// ConvertToPDF
// @Description: 转换文件为pdf
// @param filePath 需要转换的文件
// @param outPath 转换后的PDF文件存放目录
// @return string
func ConvertToPDF(filePath string, outPath string) (string, error) {
	// fmt.Println(filePath, outPath)
	// // 1、拼接执行转换的命令
	// commandName := ""
	// var params []string
	// if runtime.GOOS == "windows" {
	// 	commandName = "cmd"
	// 	params = []string{"/c", "soffice", "--headless", "--invisible", "--convert-to", "pdf", filePath, "--outdir", outPath}
	// } else if runtime.GOOS == "linux" {
	// 	commandName = "libreoffice"
	// 	params = []string{"--invisible", "--headless", "--convert-to", "pdf", filePath, "--outdir", outPath}
	// }
	// // 开始执行转换
	// cmd := exec.Command(commandName, params...)
	// _, err := cmd.Output()
	// if err != nil {
	// 	log.Println("Error: <", err, "> when exec command read out buffer")
	// 	return "", err
	// } else {
	// 	return outPath, nil
	// }
	spfilename := strings.Split(filePath, ".")
	url := "http://localhost:3000/forms/libreoffice/convert"

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		logx.Error(err)
	}
	defer file.Close()

	// 创建一个缓冲区
	var buffer bytes.Buffer

	// 创建一个新的 multipart.Writer
	writer := multipart.NewWriter(&buffer)

	// 添加文件到表单数据
	part, err := writer.CreateFormFile("files", "file."+spfilename[len(spfilename)-1])
	if err != nil {
		logx.Error(err)
	}
	if _, err := io.Copy(part, file); err != nil {
		logx.Error(err)
	}

	// 关闭 multipart.Writer
	if err := writer.Close(); err != nil {
		logx.Error(err)
	}

	// 设置 Content-Type 头部
	contentType := writer.FormDataContentType()

	// 创建请求
	req, err := http.NewRequest("POST", url, &buffer)
	if err != nil {
		logx.Error(err)
	}
	req.Header.Set("Content-Type", contentType)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logx.Error(err)
	}
	defer resp.Body.Close()

	// 读取响应体并写入文件
	output, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logx.Error(err)
	}

	// 保存文件
	outputFile, err := os.Create(spfilename[0] + ".pdf")
	if err != nil {
		logx.Error(err)
	}
	defer outputFile.Close()

	if _, err := outputFile.Write(output); err != nil {
		logx.Error(err)
	}

	log.Println("File saved successfully.")
	return outPath, nil
}

func Img2pdf(src, as string) {
	//	src = flag.String("src", "./", "the images directory path (png, jpg & jpeg) files")
	//  as  = flag.String("as", "./result.pdf", "the result filename")

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	jpgs, _ := filepath.Glob(src + "/*.jpg")
	pngs, _ := filepath.Glob(src + "/*.png")
	jpegs, _ := filepath.Glob(src + "/*.jpeg")
	files := append(jpgs, append(jpegs, pngs...)...)
	for i := 0; i < len(files); i++ {
		// fmt.Println(i+1, ")- adding ", files[i])
		img, err := loadImage(files[i])
		if err != nil {
			fmt.Println("Error loading image:", err)
			continue
		}
		width, height := float64(img.Bounds().Dx()), float64(img.Bounds().Dy())

		// 计算缩放比例
		scale := minScale(width, height, float64(gopdf.PageSizeA4.W), float64(gopdf.PageSizeA4.H))

		// 计算缩放后的尺寸
		scaledWidth, scaledHeight := width*scale, height*scale
		// 计算图片在页面中的位置，使其水平居中
		x := (float64(gopdf.PageSizeA4.W) - scaledWidth) / 2
		y := (float64(gopdf.PageSizeA4.H) - scaledHeight) / 2
		if x < 0 {
			continue
		}
		pdf.AddPage()
		pdf.Image(files[i], x, y, &gopdf.Rect{W: scaledWidth, H: scaledHeight})
	}
	fmt.Println("saving to ", as, " ...")
	pdf.WritePdf(as)
}

func Html2Img(url, savePath string) {
	const domain = "http://localhost:25600/"
	const defaultUsername = "doctron"
	const defaultPassword = "lampnick"
	client := doctron.NewClient(context.Background(), domain, defaultUsername, defaultPassword)
	req := doctron.NewDefaultHTML2ImageRequestDTO()
	req.ConvertURL = url
	response, err := client.HTML2Image(req)
	if err != nil {
		logx.Error(err)
	}
	// log.Println(len(response.Data))

	// 保存图片
	if err := SaveImageFromBytes(response.Data, savePath); err != nil {
		logx.Error(err)
	}
}

func SaveImageFromBytes(data []byte, filename string) error {
	// 使用image.Decode来解码数据流
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return err
	}
	// 创建输出文件
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()
	// 使用jpeg.Encode来保存图片
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 100})
}

func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		// 检查ipnet结构体是否nil，如果不是nil，则转换为ipNet类型
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", err
}

func loadImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func minScale(width, height, maxWidth, maxHeight float64) float64 {
	if width > maxWidth || height > maxHeight {
		scaleX := maxWidth / width
		scaleY := maxHeight / height
		return min(scaleX, scaleY)
	}
	return 1
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func IntToMSS(mileseconds int) string {
	minutes := mileseconds / 1000 / 60
	secs := mileseconds / 1000 % 60
	return fmt.Sprintf("%02d:%02d", minutes, secs)
}
