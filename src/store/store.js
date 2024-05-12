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