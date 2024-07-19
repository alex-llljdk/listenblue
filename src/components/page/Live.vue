<template>
    <div class="home h-full w-full flex">
        <el-aside width="80px" style="background: #e9efff">
            <router-link to="/home">
                <div class="text-center mb-8 mt-8">
                    <div class="rounded-xl bg-white mx-4 py-3 flex justify-center">
                        <svg
                            t="1715682947773"
                            class="icon"
                            viewBox="0 0 1024 1024"
                            version="1.1"
                            xmlns="http://www.w3.org/2000/svg"
                            p-id="14609"
                            width="20"
                            height="20"
                        >
                            <path
                                d="M238.208798 504.759085c114.078175 125.190257 330.334497 341.639984 441.040878 454.927143L785.790178 856.32207c-36.308951-36.902468-343.031681-348.783685-344.618828-358.506117 62.496379-64.083527 246.805082-251.967662 324.579423-339.65784l-95.228875-93.844342L238.208798 504.759085z"
                                fill="#272636"
                                p-id="14610"
                            />
                        </svg>
                    </div>
                </div>
            </router-link>
            <div class="text-center mt-8">
                <div class="aside-icon text-center mb-8 block px-4 cursor-pointer" @click="clickSaveLive">
                    <div class="flex justify-center mb-2">
                        <svg
                            t="1715676163293"
                            class="icon"
                            viewBox="0 0 1024 1024"
                            version="1.1"
                            xmlns="http://www.w3.org/2000/svg"
                            p-id="4267"
                            width="30"
                            height="30"
                        >
                            <path
                                d="M925.248 356.928l-258.176-258.176a64 64 0 0 0-45.248-18.752H144a64 64 0 0 0-64 64v736a64 64 0 0 0 64 64h736a64 64 0 0 0 64-64V402.176a64 64 0 0 0-18.752-45.248zM288 144h192V256H288V144z m448 736H288V736h448v144z m144 0H800V704a32 32 0 0 0-32-32H256a32 32 0 0 0-32 32v176H144v-736H224V288a32 32 0 0 0 32 32h256a32 32 0 0 0 32-32V144h77.824l258.176 258.176V880z"
                                p-id="4268"
                            />
                        </svg>
                    </div>
                    <div class="text-sm">
                        <p>保存</p>
                    </div>
                </div>
            </div>
        </el-aside>
        <div class="flex w-full">
            <div class="w-3/5 bg-blue-50 px-4 py-2">
                <div class="bg-white flex p-2 justify-between rounded-t-2xl">
                    <div class="flex items-center">
                        <el-avatar shape="square" :size="100" :src="AvatarUrl"></el-avatar>
                        <div class="ml-4">
                            <div class="text-black text-4xl font-bold">{{ roomTitle }}</div>
                            <div class="text-base">房间号:{{ roomNumber }}</div>
                            <div class="text-base">用户名:{{ room_user_name }}</div>
                        </div>
                    </div>
                    <div>
                        <div class="Tagcontainer">
                            <div class="radio-tile-group">
                                <div class="input-container">
                                    <input id="walk" class="radio-button" type="radio" name="radio" />
                                    <div class="radio-tile">
                                        <div class="icon walk-icon">
                                            <svg
                                                v-if="is_open"
                                                t="1720899578498"
                                                class="icon"
                                                viewBox="0 0 1024 1024"
                                                version="1.1"
                                                xmlns="http://www.w3.org/2000/svg"
                                                p-id="1491"
                                                width="24"
                                                height="24"
                                            >
                                                <path
                                                    d="M768.25422 0q48.810328 0 94.061569 18.303873t80.333664 50.33565 56.436941 74.740814 21.354518 91.519364l0 150.49851-123.042701 0 0-122.025819q0-64.063555-36.099305-99.654419t-97.112214-35.590864q-54.911619 0-88.468719 35.590864t-33.5571 99.654419l0 124.059583-128.12711 0 0-152.532274q0-48.810328 19.320755-91.519364t53.386296-74.740814 80.333664-50.33565 101.179742-18.303873zM766.220457 693.513406l0 87.451837 0 47.793446q0 27.455809-9.660377 51.860973t-26.438928 41.692155-39.658391 27.455809-50.33565 10.168818l-514.542205 0q-27.455809 0-49.82721-9.660377t-38.641509-26.438928-24.913605-39.14995-8.643496-47.793446l0-323.368421q0-28.472691 19.829196-47.793446t46.268123-19.320755l629.449851 0q28.472691 0 47.793446 19.320755t19.320755 47.793446l0 179.988083z"
                                                    p-id="1492"
                                                ></path>
                                            </svg>
                                            <svg
                                                v-else
                                                t="1720899638233"
                                                class="icon"
                                                viewBox="0 0 1024 1024"
                                                version="1.1"
                                                xmlns="http://www.w3.org/2000/svg"
                                                p-id="1640"
                                                width="24"
                                                height="24"
                                            >
                                                <path
                                                    d="M385.150849 385.662338l-128.895105 0 0-150.377622q0-49.102897 19.436563-91.556444t53.706294-74.677323 80.815185-50.637363 101.786214-18.413586q49.102897 0 94.625375 18.413586t80.815185 50.637363 56.263736 74.677323 20.971029 91.556444l0 150.377622-123.78022 0 0-121.734266q0-64.447552-35.804196-99.74026t-97.182817-35.292707q-55.240759 0-88.999001 35.292707t-33.758242 99.74026l0 121.734266zM826.053946 447.040959q27.62038 0 47.568432 19.948052t19.948052 47.568432l0 317.122877q0 27.62038-9.718282 51.66034t-26.597403 41.942058-39.896104 28.131868-50.637363 10.22977l-516.603397 0q-27.62038 0-50.125874-10.22977t-38.361638-27.108891-24.551449-39.384615-8.695305-48.07992l0-324.283716q0-27.62038 19.436563-47.568432t47.056943-19.948052l61.378621 0 128.895105 0 255.744256 0 123.78022 0 61.378621 0z"
                                                    p-id="1641"
                                                ></path>
                                            </svg>
                                        </div>
                                        <label v-if="is_open" for="walk" class="radio-tile-label">公开房间</label>
                                        <label v-else for="walk" class="radio-tile-label">私密房间</label>
                                    </div>
                                </div>

                                <div class="input-container">
                                    <input id="bike" class="radio-button" type="radio" name="radio" />
                                    <div class="radio-tile">
                                        <div class="icon bike-icon">
                                            <svg
                                                v-if="is_record"
                                                t="1720900218294"
                                                class="icon"
                                                viewBox="0 0 1024 1024"
                                                version="1.1"
                                                xmlns="http://www.w3.org/2000/svg"
                                                p-id="15833"
                                                width="24"
                                                height="24"
                                            >
                                                <path
                                                    d="M746.666667 512a234.666667 234.666667 0 1 1 0 469.333333 234.666667 234.666667 0 0 1 0-469.333333zM896 362.666667l0.042667 150.314666a277.333333 277.333333 0 0 0-383.104 383.104L266.666667 896A138.666667 138.666667 0 0 1 128 757.333333V362.666667h768z m-36.437333 283.562666l-174.592 174.634667L635.733333 755.2a21.333333 21.333333 0 0 0-34.133333 25.6l64 85.333333a21.333333 21.333333 0 0 0 32.170667 2.261334l192-192a21.333333 21.333333 0 0 0-30.165334-30.165334zM757.333333 128A138.666667 138.666667 0 0 1 896 266.666667V298.666667H128v-32A138.666667 138.666667 0 0 1 266.666667 128h490.666666z"
                                                    p-id="15834"
                                                ></path>
                                            </svg>
                                            <svg
                                                v-else
                                                t="1720900122887"
                                                class="icon"
                                                viewBox="0 0 1024 1024"
                                                version="1.1"
                                                xmlns="http://www.w3.org/2000/svg"
                                                p-id="21721"
                                                width="200"
                                                height="200"
                                            >
                                                <path
                                                    d="M746.666667 512a234.666667 234.666667 0 1 1 0 469.333333 234.666667 234.666667 0 0 1 0-469.333333z m-105.642667 129.024a21.333333 21.333333 0 0 0 0 30.165333l75.477333 75.52-75.392 75.349334a21.333333 21.333333 0 1 0 30.165334 30.208l75.392-75.392 75.52 75.477333a21.333333 21.333333 0 1 0 30.165333-30.165333L776.874667 746.666667l75.562666-75.52a21.333333 21.333333 0 0 0-30.165333-30.165334l-75.562667 75.52-75.52-75.52a21.333333 21.333333 0 0 0-30.165333 0zM896 362.666667l0.042667 150.314666a277.333333 277.333333 0 0 0-383.104 383.104L266.666667 896A138.666667 138.666667 0 0 1 128 757.333333V362.666667h768zM757.333333 128A138.666667 138.666667 0 0 1 896 266.666667V298.666667H128v-32A138.666667 138.666667 0 0 1 266.666667 128h490.666666z"
                                                    p-id="21722"
                                                ></path>
                                            </svg>
                                        </div>
                                        <label v-if="is_record" for="walk" class="radio-tile-label">允许保存</label>
                                        <label v-else for="walk" class="radio-tile-label">禁止保存</label>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div id="dplayer" :style="autoHeight"></div>

                <div class="card">
                    <div class="chat-window">
                        <ul class="message-list p-4">
                            <div
                                v-for="(chat, index) in chatList"
                                :key="index"
                                class="message flex items-center"
                                :class="index % 2 === 0 ? 'incoming' : 'outcoming'"
                            >
                                <p class="break-all">
                                    <span class="text-blue-600 text-base">{{ chat.Name }}</span> :{{ chat.msg }}
                                </p>
                            </div>
                        </ul>
                    </div>
                    <div class="chat-input">
                        <input type="text" class="message-input" v-model="input_text" placeholder="请输入你的消息" />
                        <button class="send-button bg-white mr-1 border" v-if="lock" @click="changeLock">解锁</button>
                        <button class="send-button bg-black text-white mr-1" v-else @click="changeLock">锁定</button>
                        <button class="send-button bg-black text-white" @click="sendMsg">发送</button>
                    </div>
                </div>
            </div>
            <div class="h-full w-2/5" style="background: #f2f5fb">
                <div class="h-full editorViewport">
                    <quill-editor
                        class="h-full"
                        v-model="content"
                        :options="editorOption"
                        @blur="onEditorBlur($event)"
                        @focus="onEditorFocus($event)"
                        @change="onEditorChange($event)"
                    ></quill-editor>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
import flvjs from 'flv.js';
import DPlayer from 'dplayer';
import { userStore } from '../../store/store';
import { quillEditor } from 'vue-quill-editor';
import { JoinRoom, LiveChatUrl ,SaveLive} from '../../api/live';
import { UserInfo } from '../../api/user';
import 'quill/dist/quill.core.css';
import 'quill/dist/quill.snow.css';
import 'quill/dist/quill.bubble.css';
import { encodeWithSalt, decodeWithSalt, Salt } from '../../utils/util';
import { Message } from 'element-ui';
const userstore = userStore();

const windowHeight = parseInt(window.innerHeight);
const windowWidth = parseInt(window.innerWidth);
// toolbar标题
const titleConfig = [
    { Choice: '.ql-insertMetric', title: '跳转配置' },
    { Choice: '.ql-bold', title: '加粗' },
    { Choice: '.ql-italic', title: '斜体' },
    { Choice: '.ql-underline', title: '下划线' },
    { Choice: '.ql-header', title: '段落格式' },
    { Choice: '.ql-strike', title: '删除线' },
    { Choice: '.ql-blockquote', title: '块引用' },
    { Choice: '.ql-code', title: '插入代码' },
    { Choice: '.ql-code-block', title: '插入代码段' },
    { Choice: '.ql-font', title: '字体' },
    { Choice: '.ql-size', title: '字体大小' },
    { Choice: '.ql-list[value="ordered"]', title: '编号列表' },
    { Choice: '.ql-list[value="bullet"]', title: '项目列表' },
    { Choice: '.ql-direction', title: '文本方向' },
    { Choice: '.ql-header[value="1"]', title: 'h1' },
    { Choice: '.ql-header[value="2"]', title: 'h2' },
    { Choice: '.ql-align', title: '对齐方式' },
    { Choice: '.ql-color', title: '字体颜色' },
    { Choice: '.ql-background', title: '背景颜色' },
    { Choice: '.ql-image', title: '图像' },
    { Choice: '.ql-video', title: '视频' },
    { Choice: '.ql-link', title: '添加链接' },
    { Choice: '.ql-formula', title: '插入公式' },
    { Choice: '.ql-clean', title: '清除字体格式' },
    { Choice: '.ql-script[value="sub"]', title: '下标' },
    { Choice: '.ql-script[value="super"]', title: '上标' },
    { Choice: '.ql-indent[value="-1"]', title: '向左缩进' },
    { Choice: '.ql-indent[value="+1"]', title: '向右缩进' },
    { Choice: '.ql-header .ql-picker-label', title: '标题大小' },
    { Choice: '.ql-header .ql-picker-item[data-value="1"]', title: '标题一' },
    { Choice: '.ql-header .ql-picker-item[data-value="2"]', title: '标题二' },
    { Choice: '.ql-header .ql-picker-item[data-value="3"]', title: '标题三' },
    { Choice: '.ql-header .ql-picker-item[data-value="4"]', title: '标题四' },
    { Choice: '.ql-header .ql-picker-item[data-value="5"]', title: '标题五' },
    { Choice: '.ql-header .ql-picker-item[data-value="6"]', title: '标题六' },
    { Choice: '.ql-header .ql-picker-item:last-child', title: '标准' },
    { Choice: '.ql-size .ql-picker-item[data-value="small"]', title: '小号' },
    { Choice: '.ql-size .ql-picker-item[data-value="large"]', title: '大号' },
    { Choice: '.ql-size .ql-picker-item[data-value="huge"]', title: '超大号' },
    { Choice: '.ql-size .ql-picker-item:nth-child(2)', title: '标准' },
    { Choice: '.ql-align .ql-picker-item:first-child', title: '居左对齐' },
    { Choice: '.ql-align .ql-picker-item[data-value="center"]', title: '居中对齐' },
    { Choice: '.ql-align .ql-picker-item[data-value="right"]', title: '居右对齐' },
    { Choice: '.ql-align .ql-picker-item[data-value="justify"]', title: '两端对齐' }
];
export default {
    components: { quillEditor },
    name: 'PlayVideo',
    data() {
        return {
            content: null,
            roomNumber: '未知!',
            password: '',
            roomTitle: '未知房间',
            is_open: true,
            is_record: false,
            room_user_id: '',
            liveurl: '',
            input_text: '',
            AvatarUrl: '',
            room_user_name: '未知',

            lock: true,
            IsEnd: false,
            chaturl: '',
            dp: null,

            chatList: [],

            windowHeight: windowHeight,
            windowWidth: windowWidth,
            autoHeight: {
                height: '60%',
                width: '100%'
            },

            editorOption: {
                placeholder: '请在这里记录您的想法', //提示
                modules: {
                    toolbar: [
                        ['bold', 'italic', 'underline', 'strike'], //加粗，斜体，下划线，删除线
                        ['blockquote', 'code-block'], //引用，代码块
                        [{ header: 1 }, { header: 2 }], // 标题，键值对的形式；1、2表示字体大小
                        [{ list: 'ordered' }, { list: 'bullet' }], //列表
                        [{ script: 'sub' }, { script: 'super' }], // 上下标
                        [{ indent: '-1' }, { indent: '+1' }], // 缩进
                        [{ direction: 'rtl' }], // 文本方向
                        [{ size: ['small', false, 'large', 'huge'] }], // 字体大小
                        [{ header: [1, 2, 3, 4, 5, 6, false] }], //几级标题
                        [{ color: [] }, { background: [] }], // 字体颜色，字体背景颜色
                        [{ font: [] }], //字体
                        [{ align: [] }], //对齐方式
                        ['clean'], //清除字体样式
                        ['image', 'video'] //上传图片、上传视频
                    ] //工具菜单栏配置
                },
                readyOnly: false, //是否只读
                theme: 'snow', //主题 snow/bubble
                syntax: true //语法检测
            }
        };
    },
    beforeDestroy(){
        this.socket.close()
    },
    created() {
        // window.addEventListener('resize', this.getHeight)
        // this.getHeight()
        // this.tempRoute = Object.assign({}, this.$route);
        //加入房间后赋值
        this.tempRoute = Object.assign({}, this.$route);
        this.roomNumber = this.$route.query.rn;
        this.password = decodeWithSalt(this.$route.query.pw, Salt);
        JoinRoom(this.roomNumber, this.password).then((res) => {
            if (res.code === 200) {
                this.is_open = res.data.is_open;
                this.is_record = res.data.is_record;
                this.roomTitle = res.data.room_title;
                this.room_user_id = res.data.user_id;
                this.src = res.data.url;
                //获取房主信息
                UserInfo(userstore.user_id, this.room_user_id).then((res) => {
                    console.log('---------------', res);
                    this.room_user_name = res.userInfo.name;
                    this.AvatarUrl = res.userInfo.avatar;
                });
                //初始化直播
                this.chaturl = LiveChatUrl + `?name=${userstore.name}&user_id=${userstore.user_id}`;
                console.log(this.chaturl);
                this.initLive();
                //连接聊天socket服务器
                this.initWebSocket();
            }
        });
    },
    destroyed() {
        window.removeEventListener('resize', this.getHeight);
        //关闭定时器
        if (this.interval) {
            clearInterval(this.interval);
        }
    },
    mounted() {},
    methods: {
        onEditorBlur(editor) {},
        onEditorFocus(editor) {},
        onEditorReady(editor) {},
        onEditorChange(editor) {
            this.content = editor.html;
            console.log(editor);
        },
        initLive() {
            const flvSrc = this.src;
            console.log('flvSrc: ' + flvSrc);
            this.dp = new DPlayer({
                container: document.getElementById('dplayer'),
                // 自动播放
                autoplay: true,
                theme: '#FADFA3',
                // 循环播放
                loop: false,
                lang: 'zh-cn',
                live: true,
                danmaku: true,
                // 允许截图
                hotkey: true,
                preload: 'auto',
                volume: 0.7,
                mutex: true,

                video: {
                    url: this.src,
                    type: 'customFlv',
                    customType: {
                        customFlv: function (video, player) {
                            const flvPlayer = flvjs.createPlayer({
                                type: 'flv',
                                url: video.src
                            });
                            flvPlayer.attachMediaElement(video);
                            flvPlayer.load();
                        }
                    }
                }
            });
        },
        initWebSocket() {
            console.log('initWesocket==');
            // 建立WebSocket连接
            this.socket = new WebSocket(this.chaturl);

            // 监听WebSocket的open事件，当连接打开时触发
            this.socket.onopen = (event) => {
                var data = JSON.parse(event.data);
                console.log('onopen: ', data);
            };

            // 监听WebSocket的message事件，当收到服务器消息时触发
            this.socket.onmessage = (event) => {
                var data = JSON.parse(event.data);
                if (data.type === 'msg') {
                    if (this.chatList.length >= 30) {
                        this.chatList.shift();
                        this.chatList.push(data);
                    } else {
                        this.chatList.push(data);
                    }
                    const danmu = {
                        text: data.msg,
                        color: '#fff',
                        type: 'right'
                    };
                    this.dp.danmaku.draw(danmu);
                    if (this.lock) {
                        this.scrollToBottom();
                    }
                } else if (data.type === 'end') {
                    this.$alert('直播已经结束,点击保存可以保存哦！', '直播结束', {
                        confirmButtonText: '确定'
                    });
                    this.IsEnd = true;
                }
            };

            // 监听WebSocket的close事件，当连接关闭时触发
            this.socket.onclose = (event) => {
                console.log('socket.onclose==');
                console.log(`WebSocket连接${this.second}秒后关闭了`);
                console.log('WebSocket is closed. Code: ' + event.code + ' Reason: ' + event.wasClean);
                if (this.interval) {
                    clearInterval(this.interval);
                }
                //连接关闭，就重新连接
                // this.reconnect();
            };

            // 监听WebSocket的error事件，当发生错误时触发
            this.socket.onerror = (error) => {
                console.error('socket.onerror:', error);
            };

            this.interval = setInterval(() => {
                this.socket.send(
                    JSON.stringify({
                        type: 'ping'
                    })
                );
            }, 60 * 1000);
        },
        sendMsg() {
            if (this.socket) {
                if (this.input_text !== '') {
                    this.socket.send(
                        JSON.stringify({
                            type: 'msg',
                            msg: this.input_text
                        })
                    );
                    this.input_text = '';
                }
            }
        },
        scrollToBottom() {
            this.$nextTick(() => {
                let chat = this.$el.querySelector('.chat-window');
                chat.scrollTop = chat.scrollHeight;
            });
        },
        changeLock() {
            this.lock = !this.lock;
        },
        clickSaveLive() {
            if (this.IsEnd === true) {
                if (this.is_record === true) {
                    let text = {
                        record_timestamp: this.roomTitle,
                        source_language: "zh-CHS",
                        dest_language: "en",
                        content: this.content,
                        savedTime: this.getCurSavedTime()
                    };
                    SaveLive(this.src, JSON.stringify(text), userstore.user_id).then((res) => {
                        if (res.code===200){
                            Message.success("保存成功")
                            this.$router.push('/home');
                        }
                    })
                } else {
                    Message.info('房主没有开放保存功能！');
                }
            } else {
                Message.info('直播还没结束无法保存！');
            }
        },
        getHMSTime() {
            const now = new Date();
            const hours = String(now.getHours()).padStart(2, '0');
            const minutes = String(now.getMinutes()).padStart(2, '0');
            const seconds = String(now.getSeconds()).padStart(2, '0');
            return `${hours}:${minutes}:${seconds}`;
        },
        getCurSavedTime() {
            const hms = this.getHMSTime();
            return `已保存于 ${hms}`;
        },
    },
    computed: {
        isLogin() {
            return userStore().isLogin;
        },
        editor() {
            return this.$refs.myQuillEditor.quill;
        }
    }
};
</script>

<style>
.home {
    overflow: auto;
}

.example-title {
    font-size: 1.5em;
    font-weight: bold;
    margin-top: 2em;
    margin-bottom: 1.8em;
}

.editorViewport {
    height: 100%;
    background-color: #fff;
    border-radius: 20px !important;
    padding: 50px 50px 100px;
    position: relative;
    z-index: 1;
}

.dSJFEj {
    position: relative;
}

.jgzocF {
    display: flex;
    -webkit-box-align: center;
    align-items: center;
    padding-top: 8px;
    padding-bottom: 8px;
    height: 32px;
}

.jVabAE {
    border-radius: 24px;
    overflow: hidden;
    display: flex;
    -webkit-box-pack: center;
    justify-content: center;
    -webkit-box-align: center;
    align-items: center;
    color: rgb(255, 255, 255);
    margin-right: 12px;
    user-select: none;
    flex-shrink: 0;
}

.iHbjbC {
    font-size: 12px;
    line-height: 20px;
    color: rgba(39, 38, 77, 0.65);
    margin-right: 8px;
    transition: all 0.2s ease 0s;
    cursor: text;
    display: flex;
    -webkit-box-align: center;
    align-items: center;
    user-select: none;
    flex-shrink: 1;
    overflow: hidden;
}

.emHNMM {
    font-size: 12px;
    line-height: 20px;
    color: rgba(39, 38, 77, 0.65);
    opacity: 1;
    transition: all 0.2s ease 0s;
    cursor: text;
    user-select: none;
    flex-shrink: 0;
}

.iwvckw {
    position: relative;
    user-select: none;
}

.gnYZJU {
    box-sizing: border-box;
    display: flex;
    flex-direction: column;
    -webkit-box-align: center;
    align-items: center;
    min-height: 56px;
    border-radius: 8px;
    border: 2px solid transparent;
    background-color: rgb(255, 255, 255);
    overflow: hidden;
    position: relative;
    z-index: 10;
    word-break: break-word;
    padding: 1px;
    margin: 0px 0px 6px;
    box-shadow: 8 8 8px #9491ee !important;
}

.gnYZJU:hover {
    border-color: #b4b2e8; /* 悬停时的边框颜色 */
}

.gnYZJU_focus {
    border-color: #9491ee !important;
}

.container {
    display: flex;
    justify-content: center;
    align-items: center;
}

.radio-tile-group {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
}

.radio-tile-group .input-container {
    position: relative;
    height: 90px;
    width: 90px;
    margin: 0.5rem;
}

.radio-tile-group .input-container .radio-button {
    opacity: 0;
    position: absolute;
    top: 0;
    left: 0;
    height: 100%;
    width: 100%;
    margin: 0;
    cursor: pointer;
}

.radio-tile-group .input-container .radio-tile {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    width: 100%;
    height: 100%;
    border: 2px solid #079ad9;
    border-radius: 5px;
    padding: 1rem;
    transition: transform 300ms ease;
}

.radio-tile-group .input-container .icon svg {
    fill: #079ad9;
    width: 2rem;
    height: 2rem;
}

.radio-tile-group .input-container .radio-tile-label {
    text-align: center;
    font-size: 0.75rem;
    font-weight: 600;
    text-transform: uppercase;
    letter-spacing: 1px;
    color: #079ad9;
}

.card {
    width: 100%;
    background-color: #fff;
    border: 1px solid #ccc;
    border-radius: 5px;
    box-shadow: 2px 2px 5px rgba(0, 0, 0, 0.1);
}
.chat-window {
    height: 200px;
    overflow-y: scroll;
}

.message-list {
    list-style: none;
}

.message {
    margin-bottom: 10px;
    padding: 10px;
    border-radius: 5px;
}

.incoming {
    background-color: #e1e1e1;
}

.message p {
    font-size: 14px;
    color: #333;
    margin: 0;
}

.chat-input {
    display: flex;
    align-items: center;
    padding: 10px;
    border-top: 1px solid #ccc;
}

.message-input {
    flex: 1;
    border: none;
    outline: none;
    padding: 5px;
    font-size: 14px;
}

.send-button {
    outline: none;
    font-size: 14px;
    padding: 5px 10px;
    cursor: pointer;
}

.send-button:hover {
    background-color: rgb(255, 255, 255);
    color: rgb(0, 0, 0);
    box-shadow: 0 4px 18px 0 rgba(0, 0, 0, 0.25);
}
</style>