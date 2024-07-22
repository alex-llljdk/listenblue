// Code generated by goctl. DO NOT EDIT.
package types

type ChatRequest struct {
	Name   string `form:"name"`
	UserID int    `form:"user_id"`
}

type ClosePublishRequest struct {
	Name string `form:"name"`
}

type ClosePublishResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CreateRoomRequest struct {
	RoomTitle  string `json:"room_title"`
	IsOpen     bool   `json:"is_open"`
	Password   string `json:"password"`
	IsRecord   bool   `json:"is_record"`
	RoomNumber string `json:"room_number"`
	UserId     int    `json:"user_id"`
	UserName   string `json:"user_name"`
	UserAvatar string `json:"avatar"`
}

type CreateRoomResponse struct {
	Url  string `json:"url"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type GetLiveListResponse struct {
	Code     int            `json:"code"`
	Msg      string         `json:"msg"`
	LiveList []LiveListData `json:"live_list"`
}

type JoinData struct {
	UserID     int    `json:"user_id"`
	RoomTitle  string `json:"room_title"`
	IsOpen     bool   `json:"is_open"`
	Password   string `json:"password"`
	IsRecord   bool   `json:"is_record"`
	RoomNumBer string `json:"room_number"`
	Url        string `json:"url"`
}

type JoinRoomRequest struct {
	RoomNumber string `json:"room_number"`
	Password   string `json:"password"`
}

type JoinRoomResponse struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data JoinData `json:"data"`
}

type LiveListData struct {
	UserName   string `json:"user_name"`
	RoomTitle  string `json:"room_title"`
	RoomNumBer string `json:"room_number"`
	CoverUrl   string `json:"cover_url"`
	UserAvatar string `json:"avatar"`
}

type OnPublishRequest struct {
	Name string `form:"name"`
}

type OnPublishResponse struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	IsRecord bool   `json:"is_record"`
}

type SaveRequest struct {
	Path      string `json:"path"`
	UserID    int    `json:"user_id"`
	InnerText string `json:"inner_text"`
}

type SaveResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
