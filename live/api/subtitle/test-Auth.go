package subtitle

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

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
func genCanonicalQueryString(params map[string]string) string {
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
func GenSignHeaders(appId, appKey, method, uri string, query map[string]string) map[string]string {
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

	canonicalQueryString := genCanonicalQueryString(query)
	//fmt.Println("canonicalQueryString:"+canonicalQueryString)

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
	}
}
