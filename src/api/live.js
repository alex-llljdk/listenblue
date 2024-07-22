import request from "../utils/request"

//创建房间
export function CreateRoom(room_title, is_open, password, is_record, room_number, user_id,user_name,user_avatar){
    var url = request.defaults.baseURL + ":10004"
    return request({
        baseURL: url,
        method:'POST',
        url:"/live/createRoom",
        data: {
            "room_title":room_title,
            "is_open": is_open,
            "password": password,
            "is_record": is_record,
            "room_number":room_number,
            "user_id":user_id,
            "user_name":user_name,
            "avatar": user_avatar
        },
    })
}

//加入房间
export function JoinRoom(room_number, password){
    var url = request.defaults.baseURL + ":10004"
    return request({
        baseURL: url,
        method:'POST',
        url:"/live/joinRoom",
        data: {
            "room_number":room_number,
            "password": password,
        },
    })
}

export function SaveLive(path, inner_text, user_id) {
    var url = request.defaults.baseURL + ":10004"
    return request({
        method: "POST",
        url: "/live/saveLive",
        data: {
            'path': path,
            'inner_text': inner_text,
            'user_id':user_id,
        },
        baseURL: url,
    })
}

export function GetLiveList() {
    var url = request.defaults.baseURL + ":10004"
    return request({
        method: "Get",
        url: "/live/liveList",
        baseURL: url,
    })
}

export const LiveChatUrl = "ws://172.30.230.46:10004/live/chat"
export const LiveSubtitleUrl = "ws://172.30.230.46:10004/live/subtitle"