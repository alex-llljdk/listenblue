import request from "../utils/request"
export function uploadVideo(data) {
    var url = request.defaults.baseURL + ":10002"
    return request({
        method: "POST",
        url: "/video/upload",
        headers: {
            "Content-Type": "multipart/form-data"
        },
        baseURL: url,
        data: data,
        timeout: 1000 *60*15,
    })
}

export function uploadLargeVideo(data) {
    var url = request.defaults.baseURL + ":10002"
    return request({
        method: "POST",
        url: "/video/uploadLarge",
        headers: {
            "Content-Type": "multipart/form-data"
        },
        baseURL: url,
        data: data,
        timeout: 1000 *60*15,
    })
}


export function mergeVideo(md5, total, filename, user_id, type) {
    var url = request.defaults.baseURL + ":10002"
    return request({
        method: "POST",
        baseURL: url,
        url: "/video/merge",
        data: {
            'md5': md5,
            'total': total,
            'filename': filename,
            'user_id': user_id,
            'type': type,
        },
        timeout: 1000 *60*15,
    })
}

export function uploadVideoUrl(VideoUrl,user_id) {
    var url = request.defaults.baseURL + ":10002"
    return request({
        method: "POST",
        baseURL: url,
        url: "/video/uploadUrl",
        data: {
            'url': VideoUrl,
            'user_id': user_id,
        },
        timeout: 1000 *60*15,
    })
}