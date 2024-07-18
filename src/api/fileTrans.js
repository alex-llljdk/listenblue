import request from "../utils/request"
export function uploadFile(data) {
    var url = request.defaults.baseURL + ":10006"
    return request({
        method: "POST",
        url: "/file/upload",
        headers: {
            "Content-Type": "multipart/form-data"
        },
        baseURL: url,
        data: data,
        timeout: 1000 *60*15,
    })
}
