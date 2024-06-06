import request from "../utils/request"

//登录
export function Login(email, pwd){
    return request({
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
    return request({
        method:'POST',
        url:"/user/register",
        data: {
            'name': name,
            'email': email,
            'password': pwd,
            'avatar': "111",
        },
    })
}

