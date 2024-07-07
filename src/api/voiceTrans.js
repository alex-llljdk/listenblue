import request from "../utils/request"
export function voiceTrans(text, from, to) {
    var url = request.defaults.baseURL + ":10003"
    
    return request({
        method: "POST",
        url: "/voice/trans",
        baseURL: url,
        data: {
            'text': text,
            'from': from,
            'to': to,
        },
    })
}

export function saveVoice(data) {
    var url = request.defaults.baseURL + ":10003"
    return request({
        method: "POST",
        url: "/voice/save",
        headers: {
            "Content-Type": "multipart/form-data"
        },
        baseURL: url,
        data: data,
    })
}