import request from "../utils/request"

//登录
export function Login(email, pwd){
    var url = request.defaults.baseURL + ":10001"
    return request({
        baseURL: url,
        method:'POST',
        url:"/user/login",
        data: {
            'email':email,
            'password': pwd,
        },
    })
}

//注册
export function Register( email, pwd,name,avtar){
    var url = request.defaults.baseURL + ":10001"
    return request({
        method:'POST',
        baseURL: url,
        url:"/user/register",
        data: {
            'name': name,
            'email': email,
            'password': pwd,
            'avatar': avtar,
        },
    })
}

//上传头像
export function UploadAvatar(data){
    var url = request.defaults.baseURL + ":10001"
    return request({
        method:'POST',
        baseURL: url,
        url:"/user/uploadImg",
        data: data,
    })
}

//上传头像
export function UserInfo(user_id, action_id){
    var url = request.defaults.baseURL + ":10001"
    return request({
        method:'GET',
        baseURL: url,
        headers: {
            "Content-Type": "multipart/form-data"
        },
        url:"/user/userinfo?user_id="+user_id+"&action_id="+action_id,
    })
}