package types

// Response 定义了整个响应的结构
type TransTextResponse struct {
	Code int `json:"code"`
	Data struct {
		Text        string `json:"text"`
		From        string `json:"from"`
		To          string `json:"to"`
		Translation string `json:"translation"`
	} `json:"data"`
	Msg       string `json:"msg"`
	RequestId string `json:"requestId"`
}

// Response 结构体用于映射 JSON 对象
type SourceContent struct {
	Timestamp string `json:"timestamp"`
	Content   string `json:"content"`
}

type ChatResponse struct {
	Code int `json:"code"`
	Data struct {
		SessionId string `json:"sessionId"`
		RequestId string `json:"requestId"`
		Content   string `json:"content"`
		Provider  string `json:"provider"`
		Model     string `json:"model"`
	} `json:"data"`
	Msg string `json:"msg"`
}
