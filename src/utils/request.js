import axios from 'axios';
import {loginStore} from '../store/store'
import { Message } from 'element-ui';
import { error } from 'autoprefixer/lib/utils';
import router from '../router';
const loginstore = loginStore()

const service = axios.create({
    // process.env.NODE_ENV === 'development' 来判断是否开发环境
    // easy-mock服务挂了，暂时不使用了
    // baseURL: 'https://www.easy-mock.com/mock/592501a391470c0ac1fab128',
    baseURL: '127.0.0.1:8080',
    timeout: 15000,
    headers: {
        "content-type": "application/json",
    },
});

//请求拦截器
service.interceptors.request.use(
    config => {
        const token = loginstore.token
        if (token) {
            config.headers.Authorization = token
        }
        return config;
    },
    error => {
        Message.error(error);
        return Promise.reject();
    }
);

//响应拦截器
service.interceptors.response.use(
    response => {
        if (response.status === 200) {
            return response.data;
        } else if (response.status ===401){
            Message.error(response.data.status_message + ",请重新登录")
            loginstore.user_id=-1
            loginstore.avatar=""
            loginstore.name=""
            router.push("/home")
            window.location.reload()
            return Promise.reject();
        }
        else {
            Message.error(response.data.status_message)
            Promise.reject();
        }
    },
    error => {
        Message.error(error);
        return Promise.reject();
    }
);

export default service;
