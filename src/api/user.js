import request from "../utils/request"

//登录
export function Login(data){
    return request({
        method:'POST',
        url:"/user/login",
        data: data,
    })
}

//注册
export function Register(data){
    return request({
        method:'POST',
        url:"/user/register",
        data: data,
    })
}

