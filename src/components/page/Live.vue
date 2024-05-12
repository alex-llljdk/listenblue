<template>
    <div class="home h-full flex items-center">
        <div id="dplayer" :style="autoHeight"></div>
        <el-button @click="initWebSocket">server</el-button>
        <el-button @click="send">send</el-button>
        <el-button @click="closeWs">send</el-button>
    </div>
</template>

<script>
import flvjs from 'flv.js';
import DPlayer from 'dplayer';
import { useMyStore } from '../../store/store';
const windowHeight = parseInt(window.innerHeight);
const windowWidth = parseInt(window.innerWidth);
export default {
    name: 'PlayVideo',
    props: {
        src: {
            type: String,
            default: 'http://127.0.0.1/live?port=1935&app=myapp&stream=1234'
        }
    },
    data() {
        return {
            windowHeight: windowHeight,
            windowWidth: windowWidth,
            autoHeight: {
                height: '100%',
                width: '60%'
            },
            uri: '/asr/v2',
            domain: 'api-ai.vivo.com.cn',
            wsUrl: 'ws://localhost:8889/'
        };
    },
    created() {
        // window.addEventListener('resize', this.getHeight)
        // this.getHeight()
        this.tempRoute = Object.assign({}, this.$route);
    },
    destroyed() {
        window.removeEventListener('resize', this.getHeight);
    },
    mounted() {
        const flvSrc = this.src;
        console.log('flvSrc: ' + flvSrc);
        const dp = new DPlayer({
            container: document.getElementById('dplayer'),
            // 自动播放
            autoplay: true,
            theme: '#FADFA3',
            // 循环播放
            loop: false,
            lang: 'zh-cn',
            live: true,
            // 允许截图
            hotkey: true,
            preload: 'auto',
            volume: 0.7,
            mutex: true,
            video: {
                url: flvSrc,
                type: 'customFlv',
                customType: {
                    customFlv: function(video, player) {
                        const flvPlayer = flvjs.createPlayer({
                            type: 'flv',
                            url: video.src
                        });
                        flvPlayer.attachMediaElement(video);
                        flvPlayer.load();
                    }
                }
            },
            highlight: [
                {
                    time: 20,
                    text: '这是第 20 秒'
                },
                {
                    time: 120,
                    text: '这是 2 分钟'
                }
            ]
        });
    },
    methods: {
        initWebSocket() {
            console.log('initWesocket==');

            // 建立WebSocket连接
            this.socket = new WebSocket(this.wsUrl);

            // 监听WebSocket的open事件，当连接打开时触发
            this.socket.onopen = event => {
                console.log('socket.onopen==');

                //可删除，只是用来测试后端自动断开连接的时间
                setInterval(() => {
                    this.second++;
                }, 1000);
            };
            // 监听WebSocket的message事件，当收到服务器消息时触发
            this.socket.onmessage = event => {
                console.log('socket.onmessage==', event.data);
                var data = JSON.parse(event.data);
                switch (
                    data.type //根据WebSocket返回值的某个字段，区分改做什么事情
                ) {
                    case 'connect': // // {"type":"connect","data":"7f0000010b540000000a"}
                        //写你自己的逻辑
                        break;
                    case 'xxx':
                        //写你自己的逻辑
                        break;
                }
            };
            // 监听WebSocket的close事件，当连接关闭时触发
            this.socket.onclose = event => {
                console.log('socket.onclose==');
                console.log(`WebSocket连接${this.second}秒后关闭了`);
                //连接关闭，就重新连接
                // this.reconnect();
            };

            // 监听WebSocket的error事件，当发生错误时触发
            this.socket.onerror = error => {
                console.error('socket.onerror:', error);
            };
        },
        reconnect() {
            //断线重新连接
            console.log('断线重新连接==');
            var that = this;
            if (that.lockReconnect) {
                return;
            }
            that.lockReconnect = true;
            //没连接上会一直重连，设置延迟避免请求过多
            setTimeout(function() {
                //重新连接
                that.initWebSocket();
                that.lockReconnect = false;
            }, 5000);
        },
        send(){
          this.socket.send("11111111111")
        },
        closeWs(){
          this.socket.close()
        }
    },
    computed: {
        appid() {
            return useMyStore().appid;
        },
        appkey() {
            return useMyStore().appkey;
        }
    }
};
</script>

<style>
.home {
    text-align: center;
}

.example-title {
    font-size: 1.5em;
    font-weight: bold;
    margin-top: 2em;
    margin-bottom: 1.8em;
}
</style>