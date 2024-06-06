<template>
    <div class="home h-full w-full flex">
        <el-aside width="80px" style="background: #e9efff;">
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
                <div class="aside-icon text-center mb-8 block px-4 cursor-pointer">
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
            <div class="w-3/5 bg-blue-50 px-4" >
                <div id="dplayer" :style="autoHeight"></div>
                <div class="contentContainer w-full h-1/2">
                    <div class="contentHeader w-full" style="height:12%">
                        <div class="flex justify-between w-full items-center">
                            <div class="w-24">
                                <img src="../../assets/img/zhuanxielogo.png" alt="">
                            </div>
                            <div>
                                <div class="hover_input focus_input">
                                    <input
                                        type="text"
                                        class="hip fip pl-2 font-semibold text-lg bg-transparent border-2 border-blue-50 border-opacity-30 focus:border-2 focus:outline-none rounded-md max-w-lg"
                                        v-model="recordInfo"
                                        :style="{ width: inputWidth + 'px' }"
                                        @input="handleInput"
                                    />
                                </div>
                                <div>
                                    <span class="text-xs text-gray-400 select-none ml-2">{{ savedTime }}</span>
                                </div>
                            </div>

                            <div class="flex text-lg items-center">
                                <el-popover placement="bottom-end" width="280" trigger="hover" visible-arrow="false">
                                    <div class="p-2 select-none">
                                        <div class="font-semibold text-black text-sm leading-6">翻译</div>
                                        <div class="flex items-center mb-2">
                                            <div class="text-xs whitespace-nowrap mr-3">选择语言</div>
                                            <el-select v-model="transType" placeholder="请选择">
                                                <el-option
                                                    v-for="item in transOptions"
                                                    :key="item.value"
                                                    :label="item.label"
                                                    :value="item.value"
                                                >
                                                </el-option>
                                            </el-select>
                                        </div>
                                        <div class="flex items-center">
                                            <div class="text-xs mr-3">显示语言</div>
                                            <el-radio-group v-model="showType">
                                                <el-radio-button label="0">双语显示</el-radio-button>
                                                <el-radio-button label="1">纯译文显示</el-radio-button>
                                            </el-radio-group>
                                        </div>
                                    </div>
                                    <div
                                        class="tools-icon cursor-pointer w-10 h-10 rounded-lg flex items-center justify-center mr-3"
                                        slot="reference"
                                    >
                                        <svg
                                            t="1715703178965"
                                            class="icon"
                                            viewBox="0 0 1024 1024"
                                            version="1.1"
                                            xmlns="http://www.w3.org/2000/svg"
                                            p-id="35183"
                                            width="30"
                                            height="30"
                                        >
                                            <path
                                                d="M889.362963 259.792593c-11.377778-64.474074-60.681481-113.777778-125.155556-125.155556C680.77037 121.362963 595.437037 113.777778 512 113.777778s-168.77037 7.585185-252.207407 20.859259c-64.474074 11.377778-113.777778 60.681481-125.155556 125.155556C121.362963 343.22963 113.777778 428.562963 113.777778 512s7.585185 168.77037 20.859259 252.207407c11.377778 64.474074 60.681481 113.777778 125.155556 125.155556 83.437037 13.274074 166.874074 20.859259 252.207407 20.859259 83.437037 0 168.77037-7.585185 254.103704-20.859259 64.474074-11.377778 113.777778-60.681481 125.155555-125.155556 11.377778-83.437037 18.962963-166.874074 18.962963-252.207407 0-83.437037-7.585185-166.874074-20.859259-252.207407z m-551.822222 30.34074c34.133333 24.651852 60.681481 51.2 83.437037 75.851852l-24.651852 24.651852c-20.859259-24.651852-49.303704-49.303704-83.437037-75.851852l24.651852-24.651852z m15.17037 422.874074l-13.274074-30.34074c7.585185-7.585185 11.377778-15.17037 11.377778-22.755556V468.385185h-70.162963v-34.133333h102.4V644.740741c15.17037-13.274074 30.340741-30.340741 47.407407-49.303704l9.481482 36.02963c-26.548148 30.340741-56.888889 56.888889-87.22963 81.54074z m365.985185-62.577777h-123.259259v83.437037h-34.133333v-83.437037h-113.777778v-34.133334h113.777778v-54.992592h-92.918519v-32.237037h92.918519v-45.511111h34.133333v45.511111h98.607407v32.237037h-98.607407v54.992592h123.259259v34.133334z m-3.792592-163.081482c-54.992593-7.585185-100.503704-20.859259-136.533334-41.718518-37.925926 18.962963-83.437037 36.02963-138.429629 47.407407l-17.066667-28.444444c49.303704-11.377778 91.022222-24.651852 125.155556-39.822223-28.444444-22.755556-49.303704-53.096296-62.577778-87.229629H455.111111v-30.340741h233.244445v28.444444c-15.17037 34.133333-39.822222 64.474074-77.748149 91.022223 32.237037 15.17037 70.162963 24.651852 115.674074 30.34074l-11.377777 30.340741z"
                                                p-id="35184"
                                                fill="#2c2c2c"
                                            />
                                            <path
                                                d="M650.42963 337.540741h-132.740741c13.274074 28.444444 34.133333 51.2 60.681481 70.162963 32.237037-20.859259 56.888889-43.614815 72.05926-70.162963z"
                                                p-id="35185"
                                                fill="#2c2c2c"
                                            />
                                        </svg>
                                    </div>
                                </el-popover>
                            </div>
                        </div>
                    </div>
                    <div class="contentBody overflow-auto pb-2" style="height:88%">
                        <div class="w-full h-96 px-2" v-if="isVideoOpen">
                            <d-player id="dplayer" class="h-full" :options="playerOptions"></d-player>
                        </div>

                        <div class="scrollContainer overflow-auto" :style="{ height: isVideoOpen ? 'calc(100% - 384px)' : '100%' }">
                            <div class="dSJFEj" v-for="(item_recording, index) in recording_items" :key="index">
                                <div class="jgzocF">
                                    <div class="jVabAE">
                                        <svg t="1714715933780" class="icon" viewBox="0 0 1024 1024" width="25" height="25">
                                            <path
                                                d="M111.8 596c-18.5 30.8-20.8 71.1-6.9 119.7 13.7 48.1 37.2 81.2 69.7 98.3 17.3 9.1 35 12.6 50.6 13.4-16.5-40.5-47.5-134.4-55.7-279.2-16.8 7.1-41.7 21.2-57.7 47.8zM854.8 548.4C846.3 683.9 814 785 798.2 827.3c15.9-0.7 34-4.2 51.5-13.5 32.3-17.2 55.7-50.2 69.4-98.2 13.9-48.7 11.6-88.9-6.9-119.7-15.9-26.3-40.6-40.4-57.4-47.5z"
                                                fill="#A0D3F8"
                                                p-id="9612"
                                            />
                                            <path
                                                d="M939.4 579.1c-24.1-39.8-62.8-57.6-82.9-64.4 0.1-1.9 0.1-3.9 0.1-5.8 0-190-154.6-344.6-344.6-344.6S167.4 318.9 167.4 508.9c0 1.9 0 3.9 0.1 5.8-20.2 6.9-58.8 24.6-82.9 64.4-23.4 38.7-27 87.7-10.5 145.4 16.2 56.8 45 96.5 85.5 117.8 58.8 31 120.1 11.2 122.7 10.3 8.2-2.7 12.8-11.5 10.3-19.8l-6.5-21.8c-5.8-20.3-16.6-62-24.5-112.6-16.3-104.3-20.4-246.5 53-332.2 43.2-50.4 109.6-76 197.5-76s154.3 25.6 197.5 76c73.3 85.7 69.3 228 53 332.2-7.9 50.6-18.8 92.3-24.4 112.4l-6.6 22c-2.5 8.3 2.1 17 10.3 19.8 1.5 0.5 21.7 7 49.6 7 21.5 0 47.5-3.9 73.1-17.4 40.5-21.4 69.3-61 85.5-117.8 16.2-57.6 12.7-106.5-10.7-145.3zM174.6 814c-32.5-17.1-55.9-50.2-69.7-98.3-13.9-48.6-11.5-88.9 6.9-119.7 15.9-26.5 40.9-40.7 57.7-47.7 8.2 144.8 39.2 238.7 55.7 279.2-15.6-0.9-33.3-4.3-50.6-13.5z m600.1-16.5c6.3-24.7 13.5-57 19.3-94.1 17.4-110.9 21-263.1-60.3-358-49.5-57.9-124.1-87.2-221.8-87.2s-172.3 29.3-221.8 87.2c-81.2 95-77.6 247.1-60.2 358 6.9 44.1 15.9 81.5 22.2 105-16.6-43.2-46-137.8-51.8-280.8v-0.1-0.2l-0.4-5.4c-0.3-4.3-0.6-8.6-0.6-13 0-172.4 140.2-312.6 312.6-312.6s312.6 140.2 312.6 312.6c0 4.4-0.3 8.7-0.6 13l-0.4 5.4c0 0.8 0.2 1.5 0.3 2.2-5.5 121.4-32.1 217.5-49.1 268z m144.4-81.8c-13.7 48-37 81-69.4 98.2-17.6 9.3-35.7 12.7-51.5 13.5 15.9-42.4 48.2-143.5 56.6-278.9 16.8 7.1 41.5 21.2 57.3 47.6 18.5 30.7 20.9 71 7 119.6z"
                                                fill="#108EE9"
                                                p-id="9613"
                                            />
                                            <path
                                                d="M368 605.5c-4.4 0-8 3.6-8 8v131.2c0 4.4 3.6 8 8 8s8-3.6 8-8V613.5c0-4.5-3.6-8-8-8zM320 653.5c-4.4 0-8 3.6-8 8v35.2c0 4.4 3.6 8 8 8s8-3.6 8-8v-35.2c0-4.5-3.6-8-8-8zM416 541.5c-4.4 0-8 3.6-8 8v259.2c0 4.4 3.6 8 8 8s8-3.6 8-8V549.5c0-4.5-3.6-8-8-8zM464 589.5c-4.4 0-8 3.6-8 8v163.2c0 4.4 3.6 8 8 8s8-3.6 8-8V597.5c0-4.5-3.6-8-8-8zM512 557.5c-4.4 0-8 3.6-8 8v227.2c0 4.4 3.6 8 8 8s8-3.6 8-8V565.5c0-4.5-3.6-8-8-8zM560 589.5c-4.4 0-8 3.6-8 8v163.2c0 4.4 3.6 8 8 8s8-3.6 8-8V597.5c0-4.5-3.6-8-8-8zM608 541.5c-4.4 0-8 3.6-8 8v259.2c0 4.4 3.6 8 8 8s8-3.6 8-8V549.5c0-4.5-3.6-8-8-8zM656 605.5c-4.4 0-8 3.6-8 8v131.2c0 4.4 3.6 8 8 8s8-3.6 8-8V613.5c0-4.5-3.6-8-8-8zM704 653.5c-4.4 0-8 3.6-8 8v35.2c0 4.4 3.6 8 8 8s8-3.6 8-8v-35.2c0-4.5-3.6-8-8-8z"
                                                fill="#108EE9"
                                                p-id="9614"
                                            />
                                        </svg>
                                    </div>
                                    <div class="iHbjbC">
                                        <span>语音识别</span>
                                    </div>
                                    <div class="emHNMM">
                                        <span>{{ recording_items[index].time_stamp }}</span>
                                    </div>
                                </div>

                                <div class="iwvckw">
                                    <div class="gnYZJU" :class="item_recording.isActive ? 'gnYZJU_focus ring-1 ring-purple-300' : 'gnYZJU'">
                                        <div class="jItVgd">
                                            <div class="kCofWf">
                                                <div
                                                    class="kFJVLr w-full"
                                                    contenteditable
                                                    @blur="updateContent(index, $event)"
                                                    @focusin="inDiv(index, $event)"
                                                    @focusout="outDiv(index, $event)"
                                                >{{item_recording.content}}</div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
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
                    >
                    </quill-editor>
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
import 'quill/dist/quill.core.css';
import 'quill/dist/quill.snow.css';
import 'quill/dist/quill.bubble.css';
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
    props: {
        src: {
            type: String,
            default: 'http://127.0.0.1/live?port=1935&app=myapp&stream=1234'
        }
    },
    data() {
        return {
            content: null,
            recordInfo: '',
            inputWidth: 204,
            savedTime: '',
            windowHeight: windowHeight,
            windowWidth: windowWidth,
            autoHeight: {
                height: '50%',
                width: '100%'
            },
            uri: '/asr/v2',
            domain: 'api-ai.vivo.com.cn',
            wsUrl: 'ws://localhost:8889/',
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
            },
            recording_items: [
                { time_stamp: '12:05', content: '', isActive: false },
            ]
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
                    customFlv: function (video, player) {
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
            this.socket.onopen = (event) => {
                console.log('socket.onopen==');

                //可删除，只是用来测试后端自动断开连接的时间
                setInterval(() => {
                    this.second++;
                }, 1000);
            };
            // 监听WebSocket的message事件，当收到服务器消息时触发
            this.socket.onmessage = (event) => {
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
            this.socket.onclose = (event) => {
                console.log('socket.onclose==');
                console.log(`WebSocket连接${this.second}秒后关闭了`);
                //连接关闭，就重新连接
                // this.reconnect();
            };

            // 监听WebSocket的error事件，当发生错误时触发
            this.socket.onerror = (error) => {
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
            setTimeout(function () {
                //重新连接
                that.initWebSocket();
                that.lockReconnect = false;
            }, 5000);
        },
        send() {
            this.socket.send('11111111111');
        },
        closeWs() {
            this.socket.close();
        },
        onEditorBlur(editor) {},
        onEditorFocus(editor) {},
        onEditorReady(editor) {},
        onEditorChange(editor) {
            this.content = editor.html;
            console.log(editor);
        },
         updateContent(index, event) {
            this.recording_items[index].content = event.target.innerText;
        },
        inDiv(index, event) {
            this.recording_items[index].isActive = true;
        },
        outDiv(index, event) {
            this.recording_items[index].isActive = false;
        },
        handleInput(event) {
            const padding = 8;
            let englishWidthPerChar = 10;
            let chineseWidthPerChar = 18;
            let spaceWidthPerChar = 5;
            let englishLength = 0;
            let chineseLength = 0;
            let spaceLength = 0;

            for (let i = 0; i < event.target.value.length; i++) {
                const char = event.target.value.charAt(i);
                if (char === ' ') {
                    spaceLength++;
                } else if (/[\u4E00-\u9FA5]/.test(char)) {
                    chineseLength++;
                } else {
                    englishLength++;
                }
            }

            const totalWidth =
                spaceLength * spaceWidthPerChar + englishLength * englishWidthPerChar + chineseLength * chineseWidthPerChar + padding;
            this.inputWidth = totalWidth;
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
    text-align: center;
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

.jItVgd {
    display: flex;
    align-items: flex-start;
    width: 100%;
    padding: 16px 24px;
    background-color: rgb(255, 255, 255);
    border-radius: 6px;
    box-sizing: border-box;
    color: rgba(39, 38, 77, 0.85);
}

.kCofWf {
    display: flex;
    flex-grow: 1;
    flex-shrink: 1;
    flex-basis: 0%;
}

.kFJVLr {
    display: block;
    min-width: 20px;
    overflow-wrap: anywhere;
    outline: none;
    font-size: 14px;
    line-height: 24px;
    min-height: 24px;
    font-family: var(--font-family);
    user-select: text;
    cursor: text;
    white-space: pre-wrap;
}
</style>