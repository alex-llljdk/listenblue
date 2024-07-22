package types

// Response 结构体用于映射 JSON 对象
type CreateResponse struct {
	Sid    string `json:"sid"`
	Action string `json:"action"`
	Data   struct {
		AudioID string `json:"audio_id"` // 注意 JSON 字段名是 audio_id，Go 结构体字段名是 AudioID
	} `json:"data"`
	Code int    `json:"code"`
	Desc string `json:"desc"`
	Type string `json:"type"`
}

type UploadResponse struct {
	Sid    string `json:"sid"`
	Action string `json:"action"`
	Data   struct {
		URL     string `json:"url"` // 假设 URL 字段是空字符串，这里使用非导出字段名来避免与 URL 包冲突
		Total   int    `json:"total"`
		Slices  int    `json:"slices"`
		AudioID string `json:"audio_id"` // 注意 JSON 字段名是 audio_id，Go 结构体字段名是 AudioID
	} `json:"data"`
	Code int    `json:"code"`
	Desc string `json:"desc"`
	Type string `json:"type"`
}

type TransResponse struct {
	Sid    string `json:"sid"`
	Action string `json:"action"`
	Data   struct {
		TaskID string `json:"task_id"`
	} `json:"data"`
	Code int    `json:"code"`
	Desc string `json:"desc"`
	Type string `json:"type"`
}

type ProgressResponse struct {
	Sid    string `json:"sid"`
	Action string `json:"action"`
	Data   struct {
		Progress int `json:"progress"` // 假设 progress 是一个整数
	} `json:"data"`
	Code int    `json:"code"`
	Desc string `json:"desc"`
	Type string `json:"type"`
}

type ResultResponse struct {
	Sid    string `json:"sid"`
	Action string `json:"action"`
	Data   struct {
		Result []struct {
			OneBest string `json:"onebest"`
			Bg      int    `json:"bg"`
			Ed      int    `json:"ed"`
			Speaker int    `json:"speaker"`
		} `json:"result"`
	} `json:"data"`
	Code int    `json:"code"`
	Desc string `json:"desc"`
	Type string `json:"type"`
}

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

type InnerText struct {
	RecordTimeStamp string `json:"record_timestamp"`
	SourceLanguage  string `json:"source_language"`
	DestLanguage    string `json:"dest_language"`
	Content         string `json:"content"`
	SaveTime        string `json:"savedTime"`
}
