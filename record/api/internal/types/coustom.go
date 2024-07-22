package types

//对得到的结果进行解析
type AnserRespone struct {
	Answer string `json:"answer"`
}

type InnerTextRespone struct {
	RecordInfo     string `json:"record_timestamp"`
	SourceLanguage string `json:"source_language"`
	DestLanguage   string `json:"dest_language"`
	Content        string `json:"content"`
	SaveTime       string `json:"savedTime"`
}

type TransResponse struct {
	TimeStamp string `json:"timestamp"`
	Content   string `json:"content"`
}

type ScanningRespone struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type ReviewResponse struct {
	Question string `json:"ques"`
	Answer   string `json:"ans"`
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
