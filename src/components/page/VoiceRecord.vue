<template>
    <div class="h-full w-full flex">
        <div class="h-full" style="background: #e9efff;">
            <el-aside width="80px">
                <router-link to="/home">
                    <div class="text-center mt-12">
                        <div class="rounded-lg bg-white mx-3 mt-7" style="font-size: 32px;">
                            <i class="el-icon-arrow-left"></i>
                        </div>
                    </div>
                </router-link>
                <div class="text-center mt-8">
                    <div>
                        <i class="el-icon-folder-checked text-4xl"></i>
                    </div>
                    <div class="text-base">保存</div>
                </div>
            </el-aside>
        </div>
        <div class="flex h-full w-full" style="background: #f2f5fb;">
            <div class="jphwIe">
                <div class="WFeMQ">
                    <div>
                        <div class="hover_input focus_input">
                            <input
                                type="text"
                                class="hip fip -ml-2 pl-2 font-semibold text-lg bg-transparent border-2 border-blue-50 border-opacity-30 focus:border-2 focus:outline-none rounded-md max-w-lg"
                                v-model="recordInfo"
                                :style="{ width: inputWidth + 'px' }"
                                @input="handleInput"
                            />
                        </div>

                        <div>
                            <span
                                class="text-xs text-gray-400"
                                style="user-select: none;"
                            >{{ savedTime }}</span>
                        </div>
                    </div>
                </div>
                <div class="h-full w-full">
                    <div v-if="showRecording" class="h-full w-full">
                        <div class="dDPXpc" style="overflow: auto;">
                            <div
                                class="dSJFEj"
                                v-for="(item_recording, index) in recording_items"
                                :key="index"
                            >
                                <div class="jgzocF">
                                    <div class="jVabAE">
                                        <svg
                                            t="1714715933780"
                                            class="icon"
                                            viewBox="0 0 1024 1024"
                                            width="25"
                                            height="25"
                                        >
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
                                    <div
                                        class="gnYZJU"
                                        :class="item_recording.isActive ? 'gnYZJU_focus ring-1 ring-purple-300' : 'gnYZJU'"
                                    >
                                        <div class="jItVgd">
                                            <div class="kCofWf">
                                                <div
                                                    class="kFJVLr w-full"
                                                    contenteditable
                                                    @blur="updateContent(index, $event)"
                                                    @focusin="inDiv(index, $event)"
                                                    @focusout="outDiv(index, $event)"
                                                ></div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <div class="h-1 w-full"></div>
                        </div>
                        <div class="kGnJQu">
                            <div class="bottomContain">
                                <div class="content flex">
                                    <div class="h-28" v-if="recording">
                                        <img src="../../assets/img/recording.gif" class="h-28 w-16" />
                                    </div>
                                    <div class="h-28" v-else>
                                        <img src="../../assets/img/recording.jpg" class="h-28 w-16" />
                                    </div>

                                    <div class="running-content">
                                        <div class="running-text">
                                            <span>{{ recording ? '录音中...' : '录音已暂停' }}</span>
                                        </div>
                                        <div class="running-time">
                                            <span>{{ formatTime(currentTime) }}</span>
                                            <span style="opacity: 0.45; margin: 0px 5px;">/</span>
                                            <span style="opacity: 0.45;">01:00:00</span>
                                        </div>
                                    </div>
                                </div>
                                <div class="right-btn">
                                    <div class="red_container" title="结束录音">
                                        <svg
                                            t="1714747107927"
                                            class="icon"
                                            viewBox="0 0 1024 1024"
                                            width="25"
                                            height="25"
                                        >
                                            <path
                                                d="M718.9248 239.328c24.1984 17.2096 45.7152 36.704 64.5376 58.4896s34.9568 45.312 48.4032 70.592c13.4464 25.28 23.6608 52.032 30.6496 80.2688 6.9952 28.2368 10.4896 56.8704 10.4896 85.9136 0 50.016-9.5424 96.9408-28.64 140.7744-19.0848 43.8336-44.9088 82.016-77.4464 114.5536s-70.7264 58.3552-114.5536 77.4464C608.5376 886.4576 561.6192 896 511.5968 896c-49.472 0-96.1344-9.5424-139.9616-28.64-43.8336-19.0848-82.1504-44.9088-114.9568-77.4464s-58.6176-70.7264-77.4464-114.5536-28.2368-90.7584-28.2368-140.7744c0-28.4992 3.36-56.4736 10.0864-83.8976 6.72-27.4304 16.2624-53.5104 28.64-78.2464 12.3712-24.7424 27.6928-47.872 45.984-69.376 18.2848-21.5168 38.72-40.8768 61.3056-58.0928 11.84-8.6016 24.608-11.8272 38.3232-9.6768 13.7152 2.1504 24.8704 8.8768 33.472 20.1664 8.608 11.296 11.84 23.936 9.6896 37.92-2.1568 13.984-8.8768 25.28-20.1728 33.8816C324.4352 352 298.4896 382.3872 280.4736 418.4192c-18.016 36.032-27.0336 74.7584-27.0336 116.1664 0 35.5008 6.7264 68.9728 20.1728 100.4352 13.4464 31.4624 31.8656 58.8928 55.2576 82.2848 23.3984 23.392 50.8224 41.952 82.2848 55.6608 31.4624 13.7216 64.9344 20.576 100.4288 20.576 35.5008 0 68.9792-6.8544 100.4352-20.576 31.4624-13.7152 58.8928-32.2688 82.2848-55.6608 23.3984-23.392 41.952-50.8224 55.6608-82.2848 13.7216-31.4624 20.576-64.9344 20.576-100.4352 0-41.952-9.6832-81.6128-29.0432-118.9888s-46.5216-68.1664-81.472-92.3712c-11.84-8.064-18.9632-19.0912-21.3824-33.0688-2.4192-13.9904 0.4032-26.8928 8.4672-38.7264 8.0704-11.296 19.0912-18.1504 33.0752-20.5696C694.1824 228.4352 707.0912 231.264 718.9248 239.328L718.9248 239.328zM511.5968 537.0048c-13.984 0-25.9456-4.9728-35.8912-14.9184-9.952-9.952-14.9248-21.92-14.9248-35.904L460.7808 179.6288c0-13.984 4.9728-26.08 14.9248-36.3008S497.6128 128 511.5968 128c14.528 0 26.7648 5.1072 36.704 15.328 9.9584 10.2208 14.9312 22.3168 14.9312 36.3008l0 306.5536c0 13.984-4.9728 25.952-14.9312 35.904C538.3552 532.032 526.1184 537.0048 511.5968 537.0048L511.5968 537.0048zM511.5968 537.0048"
                                                fill="#d1685a"
                                                p-id="13331"
                                            />
                                        </svg>
                                    </div>

                                    <div
                                        class="blue_container"
                                        v-if="recording"
                                        @click="toggleDivs"
                                        title="暂停录音"
                                    >
                                        <svg
                                            t="1714749774206"
                                            class="icon"
                                            viewBox="0 0 1024 1024"
                                            width="25"
                                            height="25"
                                        >
                                            <path
                                                d="M640 832V192h128v640h-128zM256 192h128v640H256V192z"
                                                fill="#0590DF"
                                                p-id="15295"
                                            />
                                        </svg>
                                    </div>

                                    <div
                                        class="blue_container"
                                        v-else
                                        @click="toggleDivs"
                                        title="继续录音"
                                    >
                                        <svg
                                            t="1714746776172"
                                            class="icon"
                                            viewBox="0 0 1024 1024"
                                            width="25"
                                            height="25"
                                        >
                                            <path
                                                d="M512 648c92.8 0 168-75.2 168-168V304c0-92.8-75.2-168-168-168S344 211.2 344 304v176c0 92.8 75.2 168 168 168zM392 304c0-65.6 54.4-120 120-120s120 54.4 120 120v176c0 65.6-54.4 120-120 120s-120-54.4-120-120V304z"
                                                fill="#2090E0"
                                                p-id="3768"
                                            />
                                            <path
                                                d="M768 408c-12.8 0-24 11.2-24 24v56C744 616 640 720 512 720s-232-104-232-232V432c0-12.8-11.2-24-24-24s-24 11.2-24 24v56c0 145.6 113.6 267.2 256 278.4v73.6h-64c-12.8 0-24 11.2-24 24s11.2 24 24 24h176c12.8 0 24-11.2 24-24s-11.2-24-24-24h-64v-73.6c142.4-12.8 256-132.8 256-278.4V432c0-12.8-11.2-24-24-24z"
                                                fill="#2090E0"
                                                p-id="3769"
                                            />
                                        </svg>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div v-else class="ixtgUW">
                        <div class="kGfivz">
                            <div class="epacMt">
                                <img src="../../assets/img/voice_record.jpg" />
                            </div>

                            <div class="w-2/5">
                                <div class="hObNWG">
                                    <div>音频语言</div>
                                </div>
                                <div class="pb-3 pt-1 w-full">
                                    <el-radio-group v-model="language" class="w-full">
                                        <el-radio-button label="中文" class="w-1/2"></el-radio-button>
                                        <el-radio-button label="英文" class="w-1/2"></el-radio-button>
                                    </el-radio-group>
                                </div>

                                <div class="hObNWG">
                                    <div>翻译</div>
                                </div>
                                <div class="pb-7 pt-1 w-full">
                                    <el-radio-group v-model="translation" class="w-full">
                                        <el-radio-button label="翻译" class="w-1/2"></el-radio-button>
                                        <el-radio-button label="不翻译" class="w-1/2"></el-radio-button>
                                    </el-radio-group>
                                </div>
                            </div>

                            <div class="flex justify-center w-2/5 h-16">
                                <button class="custom-button" @click="startRecording">开始录音</button>
                            </div>
                            <div class="justify-center w-2/5 h-48"></div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="h-full w-1/2" style="background: #f2f5fb;">
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
import { quillEditor } from 'vue-quill-editor';
import 'quill/dist/quill.core.css';
import 'quill/dist/quill.snow.css';
import 'quill/dist/quill.bubble.css';
import Recorder from 'js-audio-recorder';

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
    name: 'VoiceRecord',
    data() {
        return {
            // recording_items 数组用于存储要动态生成的 div 元素的数据
            recording_items: [],
            item_index: -1, // recording_items当前最大进行时item的下标

            currentTime: 0, // 当前时间，单位：秒
            maxTime: 3600, // 最大时间，单位：秒，即 01:00:00
            timer: null, // 定时器

            showRecording: false, // 控制是否显示录音时的内容
            recording: true, // 是否正在录音
            inputWidth: 204,
            recordInfo: '',
            savedTime: '',
            language: '中文',
            translation: '不翻译',
            content: null,
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
            wsUrl: 'ws://localhost:8765',
            recorder: null //多媒体对象，用来处理音频
        };
    },
    mounted() {
        document.title = 'ListenBlue-会议记录';
        this.updateInfo();
        this.initTitle();
    },
    created() {
        this.initRecorder();
    },
    computed: {
        getYMDHMSTime() {
            const now = new Date();
            const year = now.getFullYear();
            const month = String(now.getMonth() + 1).padStart(2, '0');
            const date = String(now.getDate()).padStart(2, '0');
            const hours = String(now.getHours()).padStart(2, '0');
            const minutes = String(now.getMinutes()).padStart(2, '0');
            return `${year}-${month}-${date} ${hours}:${minutes}`;
        },
        editor() {
            return this.$refs.myQuillEditor.quill;
        }
    },
    methods: {
        getHMSTime() {
            const now = new Date();
            const hours = String(now.getHours()).padStart(2, '0');
            const minutes = String(now.getMinutes()).padStart(2, '0');
            const seconds = String(now.getSeconds()).padStart(2, '0');
            return `${hours}:${minutes}:${seconds}`;
        },
        updateInfo() {
            const ymd_hms = this.getYMDHMSTime;
            const hms = this.getHMSTime();
            this.recordInfo = `${ymd_hms} 记录`;
            this.savedTime = `已保存于 ${hms}`;
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
        onEditorBlur(editor) {},
        onEditorFocus(editor) {},
        onEditorReady(editor) {},
        onEditorChange(editor) {
            this.content = editor.html;
            console.log(editor);
        },
        initTitle() {
            document.getElementsByClassName('ql-editor')[0].dataset.placeholder = '请在这里记录您的想法';
            for (let item of titleConfig) {
                let tip = document.querySelector('.quill-editor ' + item.Choice);
                if (!tip) continue;
                tip.setAttribute('title', item.title);
            }
        },
        addRecordingItem() {
            // 添加一个新的item到数组中
            this.recording_items.push({ time_stamp: this.formatTime(this.currentTime), content: '', isActive: false });
            this.item_index++;
        },
        deleteItem() {
            // 检查数组长度是否大于0
            if (this.recording_items.length > 0) {
                // 检查最后一个item的content是否为空, 删除最后一个item
                if (this.recording_items[this.recording_items.length - 1].content.length === 0) {
                    this.recording_items.pop();
                    this.item_index--;
                }
            }
        },
        toggleDivs() {
            this.recording = !this.recording; // 录音与暂停状态切换
            if (this.recording) {
                this.startTimer();
                this.addRecordingItem();
                console.log(this.item_index);
            } else {
                this.stopTimer();
                this.deleteItem();
            }
        },
        startTimer() {
            this.timer = setInterval(() => {
                this.currentTime++;
            }, 1000); // Increment time every second
        },
        stopTimer() {
            clearInterval(this.timer);
        },
        formatTime(seconds) {
            // Convert seconds to HH:MM:SS format
            const hours = Math.floor(seconds / 3600);
            const minutes = Math.floor((seconds % 3600) / 60);
            const remainingSeconds = seconds % 60;
            if (hours === 0) {
                // 如果小时数为0，则只显示分钟和秒数，格式为MM:SS
                return `${this.pad(minutes)}:${this.pad(remainingSeconds)}`;
            } else {
                // 否则，显示小时、分钟和秒数，格式为HH:MM:SS
                return `${this.pad(hours)}:${this.pad(minutes)}:${this.pad(remainingSeconds)}`;
            }
        },
        pad(value) {
            return value < 10 ? '0' + value : value;
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

        //启动浏览器麦克风，开始录制
        initRecorder() {
            this.recorder = new Recorder({
                sampleBits: 16, // 采样位数，支持 8 或 16，默认是16
                sampleRate: 16000, // 采样率，支持 11025、16000、22050、24000、44100、48000，根据浏览器默认值，我的chrome是48000
                numChannels: 1, // 声道，支持 1 或 2， 默认是1
                compiling: true //(0.x版本中生效)  // 是否边录边转换，默认是false
            });
        },
        //启动录音
        startRecording() {
            // 获取录音权限
            Recorder.getPermission().then(
                () => {
                    // 麦克风可用，执行成功时的操作
                    this.initWebSocketAndRecording();
                },
                error => {
                    // 处理获取权限失败的情况
                    console.error('无法访问麦克风:', error);
                    alert('无法访问麦克风，请检查权限设置或硬件连接');
                }
            );
        },
        inRecording() {
            try {
                let That = this;
                //获取录音数据
                const blob = this.recorder.getWAVBlob();
                //blob转为arrayBuffer
                let reader = new FileReader();
                reader.readAsArrayBuffer(blob);
                reader.onload = function() {
                    var outbuffer = this.result;
                    var arr = new Int8Array(outbuffer);
                    // //调用webSocket发送服务端
                    That.send(arr);
                };
            } catch (e) {
                console.log(e);
            }
        },
        endRecording() {
            console.log('this.recorder.getWAVBlob(): ', this.recorder.getWAVBlob());
            this.recorder.stop();
        },
        initWebSocketAndRecording() {
            console.log('initWesocket==');
            // 建立WebSocket连接
            this.socket = new WebSocket(this.wsUrl);

            // 监听WebSocket的open事件，当连接打开时触发
            this.socket.onopen = event => {
                console.log('socket.onopen成功, 开始录音');
                
                // 保证socket连接成功再开始录音
                this.showRecording = true;
                if (this.recording) {
                    this.addRecordingItem();
                    this.item_index = 0;
                    this.startTimer();
                }
                this.recorder.start(); // 开始录音
                this.inRecording();

                //可删除，只是用来测试后端自动断开连接的时间
                setInterval(() => {
                    this.second++;
                }, 1000);
            };
            // 监听WebSocket的message事件，当收到服务器消息时触发
            this.socket.onmessage = event => {
                console.log('socket.onmessage==', event.data);
                // var data = JSON.parse(event.data);
                // switch (
                //     data.type //根据WebSocket返回值的某个字段，区分改做什么事情
                // ) {
                //     case 'connect': // // {"type":"connect","data":"7f0000010b540000000a"}
                //         //写你自己的逻辑
                //         break;
                //     case 'xxx':
                //         //写你自己的逻辑
                //         break;
                // }
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
        send(msg) {
            console.log('send msg: ', msg);
            this.socket.send(msg);
        },
        closeWs() {
            this.socket.close();
        }
    }
};
</script>

<style scoped>
.hip {
    overflow: hidden;
    text-overflow: ellipsis;
    word-break: break-all;
}
.hover_input .hip:hover {
    border-color: #409eff;
    border-width: 2px;
}

.focus_input .fip:focus {
    background-color: white;
    border-color: #409eff;
    border-width: 2px;
    box-shadow: 0 0 4px rgba(0, 0, 0, 0.5);
}

.jphwIe {
    width: 50%;
    height: 100%;
    overflow: clip;
    display: flex;
    flex-direction: column;
    position: relative;
    margin: 0px auto;
    padding-left: 28px;
    padding-right: 18px;
}

.WFeMQ {
    display: flex;
    -webkit-box-pack: justify;
    justify-content: space-between;
    top: 0px;
    left: 0px;
    -webkit-box-align: center;
    align-items: flex-end;
    width: 100%;
    height: 96px;
    background-color: transparent;
    border-bottom: 1px solid transparent;
    position: relative;
}

.ixtgUW {
    flex-grow: 7;
    flex-shrink: 1;
    flex-basis: 0%;
    height: 100%;
    display: flex;
    -webkit-box-pack: center;
    justify-content: center;
    -webkit-box-align: center;
    align-items: center;
    flex-direction: column;
}

.kGfivz {
    height: 100%;
    width: 100%;
    display: flex;
    justify-content: center;
    flex-direction: column;
    -webkit-box-align: center;
    align-items: center;
    user-select: none;
    position: relative;
}

.epacMt {
    width: 40%;
    position: relative;
    margin-bottom: 20px;
}

.liWXFA {
    display: flex;
    -webkit-box-pack: center;
    justify-content: center;
    -webkit-box-align: center;
    align-items: center;
    flex-direction: column;
    width: 100%;
}

.hObNWG {
    font-family: PingFangSC-Regular;
    font-size: 12px;
    color: rgba(39, 38, 77, 0.65);
    line-height: 20px;
    margin-bottom: 8px;
    display: flex;
    -webkit-box-align: center;
    align-items: center;
}

.custom-button {
    line-height: 20%;
    width: 100%;
    height: 100%;
    font-size: 20px;
    background-color: #007bff;
    color: #fff;
    border: none;
    cursor: pointer;
    border-radius: 5px;
    transition: border-color 0.3s ease;
}

.custom-button:hover {
    background-color: #2f91fa;
}

::v-deep .el-radio-button__inner {
    width: 100%;
    height: 100%;
    border: 1 !important;
    border-radius: 10px;
    font-size: 15px !important;
    font-weight: 530;
    color: #696969;
    line-height: 24px;
    outline: none;
    box-shadow: none;
}
/*激活样式*/
::v-deep .el-radio-button__orig-radio:checked + .el-radio-button__inner {
    background: #29c8d6;
    border: 1 !important;
    color: #696969;
    font-size: 16px !important;
    font-weight: 500;
    line-height: 24px;
    outline: none;
    box-shadow: none;
}

.editorViewport {
    height: 100%;
    background-color: #fff;
    border-radius: 20px !important;
    padding: 50px 50px 100px;
    position: relative;
    z-index: 1;
}

.dDPXpc {
    height: calc(100vh - 88px - 112px);
}

.kGnJQu {
    user-select: none;
    display: flex;
    -webkit-box-align: center;
    align-items: center;
}

.kGnJQu .bottomContain {
    display: flex;
    -webkit-box-align: center;
    align-items: center;
    width: 100%;
    margin: 0px auto;
}

.kGnJQu .content {
    -webkit-box-flex: 1;
    flex-grow: 1;
}

.kGnJQu .right-btn {
    display: flex;
    padding-right: 25px;
    -webkit-box-align: center;
    align-items: center;
}

.kGnJQu .content .running-text {
    font-weight: 500;
    font-size: 17px;
    color: rgb(43, 40, 103);
    letter-spacing: 0px;
    line-height: 28px;
    display: flex;
    -webkit-box-align: center;
    align-items: center;
}

.kGnJQu .content .running-time {
    display: flex;
    justify-content: flex-start;
    font-size: 12px;
    color: rgb(16, 19, 22);
    line-height: 18px;
    padding-top: 1px;
}
.running-content {
    display: flex;
    -webkit-box-pack: center;
    justify-content: center;
    -webkit-box-align: center;
    flex-direction: column;
    padding-bottom: 12px;
    padding-left: 24px;
}
.red_container {
    border: 2px solid #fff;
    background-color: #fff;
    border-radius: 12px;
    display: inline-block;
    padding: 8px;
    margin: 8px;
    cursor: pointer; /* 添加鼠标指针样式，表明是可以点击的 */
}

.red_container:hover {
    border: 2px solid #d1685a;
    background-color: #d1685a; /* 鼠标悬停时背景颜色变为红色 */
}

.red_container:hover svg path {
    fill: white; /* 鼠标悬停时 SVG 路径填充颜色变为白色 */
}

.blue_container {
    border: 2px solid #fff;
    background-color: #fff;
    border-radius: 12px;
    display: inline-block;
    padding: 8px;
    margin: 8px;
    cursor: pointer; /* 添加鼠标指针样式，表明是可以点击的 */
}

.blue_container:hover {
    border: 2px solid #2090e0;
    background-color: #2090e0; /* 鼠标悬停时背景颜色变为蓝色 */
}

.blue_container:hover svg path {
    fill: white; /* 鼠标悬停时 SVG 路径填充颜色变为白色 */
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