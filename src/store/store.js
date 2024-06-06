import { defineStore } from 'pinia';

export const useMyStore = defineStore({
  id: 'myStore',
  state: () => ({
    appkey: 'qoDefbZxeZLUHaLV',
    appid: '3032677197',
    params : {'client_version': parse.quote('unknown'), 'package': parse.quote('unknown'),
              'sdk_version': parse.quote('3.0'), 'user_id': parse.quote('2addc42b7ae689dfdf1c63e220df52a2'),
              'android_version': parse.quote('unknown'), 'system_time': parse.quote(str(t)), 'net_type': 1,
              'nonce_str': parse.quote('xarhtv6afy7n5ime'), 'engineid': "longasrlisten"}
  }),
});

export const userStore = defineStore({
  id: 'userstore',
  state: () => ({
    isLoggedIn:false,
    user_id:-1,
    avatar:"https://img1.baidu.com/it/u=3647744349,2477516282&fm=253&fmt=auto&app=138&f=JPEG?w=380&h=380",
    name:"",
    token:"",
  }),
  getters:{
    getToken:(state)=>{
      return  state.token
    },
  },
  persist: {
    enabled: true,
    strategies: [
      {
        key: 'user',
        storage: localStorage,
      },
    ],
  },
});
