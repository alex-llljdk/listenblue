import axios from 'axios';
import {userStore} from '../store/store'
import { Message } from 'element-ui';
import { error } from 'autoprefixer/lib/utils';
import router from '../router';
const userstore = userStore()

const service = axios.create({
    // process.env.NODE_ENV === 'development' 来判断是否开发环境
    // easy-mock服务挂了，暂时不使用了
    // baseURL: 'https://www.easy-mock.com/mock/592501a391470c0ac1fab128',
    baseURL: 'http://127.0.0.1:10001',
    timeout: 15000,
});

//请求拦截器
service.interceptors.request.use(
    config => {
        console.log('config:',config)
        const token = userstore.token
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
        console.log("rsp",response )
        if (response.data.code === 200) {
            return response.data;
        } else if (response.status ===401){
            Message.error(response.data.message + ",请重新登录")
            userstore.isLoggedIn = false
            userstore.user_id=-1
            userstore.avatar=""
            userstore.name=""
            userstore.token=""
            router.push("/home")
            window.location.reload()
            return Promise.reject();
        }
        else {
            Message.error(response.data.message)
            return Promise.reject();
        }
    },
    error => {
        Message.error(error);
        return Promise.reject();
    }
);

export default service;
