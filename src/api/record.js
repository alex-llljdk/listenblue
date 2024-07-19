import request from "../utils/request"
export function getRecord(data) {
    var url = request.defaults.baseURL + ":10005"
    return request({
        method: "GET",
        url: "/record/recordById?id=" + data,
        headers: {
            "Content-Type": "multipart/form-data"
        },
        baseURL: url,
    })
}

export function getPDFRecord(data) {
    var url = request.defaults.baseURL + ":10005"
    return request({
        method: "GET",
        url: "/record/recordById?id=" + data,
        headers: {
            "Content-Type": "multipart/form-data"
        },
        baseURL: url,
    })
}

export function getRecentRecord(data) {
    var url = request.defaults.baseURL + ":10005"
    return request({
        method: "GET",
        url: "/record/recentRecordById?id=" + data,
        headers: {
            "Content-Type": "multipart/form-data"
        },
        baseURL: url,
    })
}

export function getAiRewriting(data) {
    var url = request.defaults.baseURL + ":10005"
    return request({
        method: "POST",
        url: "/record/getAiRewriting",
        data: {
            'text': data,
        },
        baseURL: url,
        timeout: 1000 *60*15,
    })
}

export function getRecordList(data) {
    var url = request.defaults.baseURL + ":10005"
    return request({
        method: "GET",
        url: "/record/recordByUserid?user_id=" + data,
        headers: {
            "Content-Type": "multipart/form-data"
        },
        baseURL: url,
    })
}