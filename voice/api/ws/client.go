package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	// The websocket connection.
	conn *websocket.Conn
	// Buffered channel of outbound messages.
	send chan []byte
}

// StartData 定义start_data结构体
type StartData struct {
	Type         string  `json:"type"`
	RequestID    string  `json:"request_id"`
	AsrInfo      AsrInfo `json:"asr_info"`
	BusinessInfo string  `json:"business_info"`
}

// AsrInfo 定义asr_info字段的结构体
type AsrInfo struct {
	FrontVadTime    int    `json:"front_vad_time"`
	EndVadTime      int    `json:"end_vad_time"`
	AudioType       string `json:"audio_type"`
	Chinese2Digital int    `json:"chinese2digital"`
	Punctuation     int    `json:"punctuation"`
}

func (c *Client) clientReadPump(s *Server) {
	defer func() {
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		fmt.Println("client read", string(message))
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		s.send <- message
	}
}

func (c *Client) clientWritePump() {
	ticker := time.NewTicker(pingPeriod)
	opcode := 1
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(opcode)
			if opcode == 1 {
				opcode = 2
			}
			if err != nil {
				return
			}
			// fmt.Println("client write", message)
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func ClientWs(s *Server) *Client {

	appId := "3032677197"
	appKey := "qoDefbZxeZLUHaLV"
	domain := "api-ai.vivo.com.cn"
	uri := "/asr/v2"
	params := url.Values{}
	params.Set("android_version", "unknown")
	params.Set("client_version", "unknown")
	params.Set("engineid", "longasrlisten")
	params.Set("net_type", "0")
	params.Set("package", "unknown")
	params.Set("sdk_version", "unknown")
	params.Set("system_time", strconv.Itoa(int(time.Now().Unix())))
	params.Set("user_id", "2addc42b7ae689dfdf1c63e220df52a2")

	var pairs []string // 用于存储键值对的切片
	paramsMap := make(map[string]string)
	for key, value := range params {
		// 假设每个键只有一个值，取第一个值作为键对应的值
		paramsMap[key] = value[0]
		pair := fmt.Sprintf("%s=%s", key, value[0])
		pairs = append(pairs, pair)
	}

	params.Encode()
	paramStr := strings.Join(pairs, "&")
	//fmt.Println(paramStr)

	fullurl := fmt.Sprintf("ws://%s%s?%s", domain, uri, paramStr)
	headers := GenSignHeaders(appId, appKey, "GET", uri, paramsMap)

	header := http.Header{}
	for key, value := range headers {
		header.Set(key, value)
	}

	//创建dialer实例
	conn, resp, err := websocket.DefaultDialer.Dial(fullurl, header)

	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("dial server error:%v\n", err)
		fmt.Println(resp.StatusCode)
	}

	client := &Client{
		conn: conn,
		send: make(chan []byte, bufSize),
	}

	go client.clientReadPump(s)
	go client.clientWritePump()

	// 创建start_data的实例，opcode报文
	startData := StartData{
		Type:      "started",
		RequestID: uuid.New().String(),
		AsrInfo: AsrInfo{
			FrontVadTime:    6000,
			EndVadTime:      2500,
			AudioType:       "pcm",
			Chinese2Digital: 1,
			Punctuation:     1,
		},
		BusinessInfo: "",
	}
	jsonData, err := json.Marshal(startData)
	if err != nil {
		fmt.Printf("Error marshalling to JSON: %v\n", err)
		return nil
	}
	client.send <- jsonData

	return client
}
