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
export function Register(name, email, pwd){
    var url = request.defaults.baseURL + ":10001"
    return request({
        method:'POST',
        baseURL: url,
        url:"/user/register",
        data: {
            'name': name,
            'email': email,
            'password': pwd,
            'avatar': "111",
        },
    })
}

