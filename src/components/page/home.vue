<template>
    <div class="w-full h-full overflow-auto">
        <div class="bg"></div>
        <div
            class="w-full h-1/5 bottom-0 absolute bg-gradient-to-b from-transparent via-blue-50 to-blue-100"
        ></div>
        <div class="pl-20 pr-10 pb-10 z-1">
            <div class="w-full h-32 flex px-28 py-4 text-6xl font-semibold mt-20">
                <div class="flex h-full z-10">
                    <div class="text-white select-none">ListenBlue,&nbsp&nbsp</div>
                </div>
                <div class="scroll h-full z-10">
                    <div
                        v-for="(item, index) in ulList"
                        :key="item.id"
                        :class="!index && play ? 'toUp' : ''"
                        class="truncate h-full text-blue-500 select-none"
                    >{{ item.msg }}</div>
                </div>
            </div>
            <div class="flex">
                <div class="ml-20 z-10">
                    <div class="flex">
                        <router-link to="/voice-record">
                            <div
                                class="hcard transform transition duration-300 hover:scale-110 rounded-2xl shadow-xl h-64 hover:shadow-xl flex cursor-pointer relative"
                            >
                                <div
                                    class="w-1/2 h-full meet_bg bg-cover bg-no-repeat rounded-l-2xl"
                                ></div>
                                <div
                                    class="w-1/2 h-full bg-gradient-to-br from-blue-300 via-blue-400 to-blue-500 rounded-r-2xl text-white p-4 bx_sd"
                                >
                                    <div class="font-semibold text-2xl">会议、课堂实时记录</div>
                                    <div class="font-medium text-base leading-7 mt-3">实时语音转文字</div>
                                    <div class="font-medium text-base leading-7">同步翻译，智能总结要点</div>
                                    <div
                                        class="hover_svg bottom-0 right-0 absolute pb-2 pr-2 opacity-0"
                                    >
                                        <svg
                                            t="1714583281873"
                                            class="icon"
                                            viewBox="0 0 1024 1024"
                                            version="1.1"
                                            xmlns="http://www.w3.org/2000/svg"
                                            p-id="13528"
                                            width="30"
                                            height="30"
                                        >
                                            <path
                                                d="M844.15488 512.70656L571.58656 254.95552c-11.84768-11.84768-29.62432-11.84768-41.48224 0-11.84768 11.84768-11.84768 29.62432 0 41.472L749.34272 503.808H201.2672c-17.77664 0-29.62432 11.84768-29.62432 29.63456 0 17.7664 11.84768 29.62432 29.62432 29.62432h548.08576l-219.2384 207.38048c-11.84768 11.84768-11.84768 29.63456 0 41.48224 5.92896 5.92896 14.81728 8.88832 20.736 8.88832 5.92896 0 14.81728-2.95936 20.736-8.88832l272.55808-257.7408c5.92896-5.92896 8.88832-14.81728 8.88832-20.736 0.01024-5.92896-2.94912-14.81728-8.87808-20.74624z m0 0"
                                                fill="#ffffff"
                                                p-id="13529"
                                            />
                                        </svg>
                                    </div>
                                </div>
                                <div></div>
                            </div>
                        </router-link>
                        <router-view />
                        <div
                            class="ml-10 hcard transform transition duration-300 hover:scale-110 rounded-2xl shadow-lg h-64 hover:shadow-xl flex cursor-pointer"
                        >
                            <div class="w-1/2 h-full class_bg bg-cover bg-no-repeat rounded-l-2xl"></div>
                            <div
                                class="bg-gradient-to-br from-purple-300 via-purple-400 to-purple-500 w-1/2 h-full rounded-r-2xl text-white p-4 bx_sd"
                            >
                                <div class="font-semibold text-2xl">音视频上传记录总结</div>
                                <div class="font-medium text-base leading-7 mt-3">音视频转文字</div>
                                <div class="font-medium text-base leading-7">章节速览总结，一键导出</div>
                                <div
                                    class="hover_svg bottom-0 right-0 absolute pb-2 pr-2 opacity-0"
                                >
                                    <svg
                                        t="1714583281873"
                                        class="icon"
                                        viewBox="0 0 1024 1024"
                                        version="1.1"
                                        xmlns="http://www.w3.org/2000/svg"
                                        p-id="13528"
                                        width="30"
                                        height="30"
                                    >
                                        <path
                                            d="M844.15488 512.70656L571.58656 254.95552c-11.84768-11.84768-29.62432-11.84768-41.48224 0-11.84768 11.84768-11.84768 29.62432 0 41.472L749.34272 503.808H201.2672c-17.77664 0-29.62432 11.84768-29.62432 29.63456 0 17.7664 11.84768 29.62432 29.62432 29.62432h548.08576l-219.2384 207.38048c-11.84768 11.84768-11.84768 29.63456 0 41.48224 5.92896 5.92896 14.81728 8.88832 20.736 8.88832 5.92896 0 14.81728-2.95936 20.736-8.88832l272.55808-257.7408c5.92896-5.92896 8.88832-14.81728 8.88832-20.736 0.01024-5.92896-2.94912-14.81728-8.87808-20.74624z m0 0"
                                            fill="#ffffff"
                                            p-id="13529"
                                        />
                                    </svg>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="flex mt-5">
                        <div
                            class="transform transition duration-300 hover:scale-110 rounded-2xl shadow-lg h-64 hover:shadow-xl hcard flex cursor-pointer"
                        >
                            <div class="w-1/2 h-full doc_bg bg-cover bg-no-repeat rounded-l-2xl"></div>
                            <div
                                class="bg-gradient-to-br from-pink-200 via-pink-300 to-pink-400 w-1/2 h-full rounded-r-2xl text-white p-4 bx_sd"
                            >
                                <div class="font-semibold text-2xl">文档、图片一键归纳</div>
                                <div class="font-medium text-base leading-7 mt-3">文档、图片总结归纳</div>
                                <div class="font-medium text-base leading-7">扫描上传，轻松阅读</div>
                                <div
                                    class="hover_svg bottom-0 right-0 absolute pb-2 pr-2 opacity-0"
                                >
                                    <svg
                                        t="1714583281873"
                                        class="icon"
                                        viewBox="0 0 1024 1024"
                                        version="1.1"
                                        xmlns="http://www.w3.org/2000/svg"
                                        p-id="13528"
                                        width="30"
                                        height="30"
                                    >
                                        <path
                                            d="M844.15488 512.70656L571.58656 254.95552c-11.84768-11.84768-29.62432-11.84768-41.48224 0-11.84768 11.84768-11.84768 29.62432 0 41.472L749.34272 503.808H201.2672c-17.77664 0-29.62432 11.84768-29.62432 29.63456 0 17.7664 11.84768 29.62432 29.62432 29.62432h548.08576l-219.2384 207.38048c-11.84768 11.84768-11.84768 29.63456 0 41.48224 5.92896 5.92896 14.81728 8.88832 20.736 8.88832 5.92896 0 14.81728-2.95936 20.736-8.88832l272.55808-257.7408c5.92896-5.92896 8.88832-14.81728 8.88832-20.736 0.01024-5.92896-2.94912-14.81728-8.87808-20.74624z m0 0"
                                            fill="#ffffff"
                                            p-id="13529"
                                        />
                                    </svg>
                                </div>
                            </div>
                        </div>
                        <div
                            class="ml-10 transform transition duration-300 hover:scale-110 rounded-2xl shadow-lg h-64 hover:shadow-xl hcard flex cursor-pointer"
                        >
                            <div class="w-1/2 h-full live_bg bg-cover bg-no-repeat rounded-l-2xl"></div>

                            <div
                                class="bg-gradient-to-br from-red-200 via-red-300 to-red-400 w-1/2 h-full rounded-r-2xl text-white p-4 bx_sd"
                            >
                                <div class="font-semibold text-2xl">实时直播,实时记录</div>
                                <div class="font-medium text-base leading-7 mt-3">直播总结，实时笔记</div>
                                <div class="font-medium text-base leading-7">双语翻译，实时字幕</div>
                                <div
                                    class="hover_svg bottom-0 right-0 absolute pb-2 pr-2 opacity-0"
                                >
                                    <svg
                                        t="1714583281873"
                                        class="icon"
                                        viewBox="0 0 1024 1024"
                                        version="1.1"
                                        xmlns="http://www.w3.org/2000/svg"
                                        p-id="13528"
                                        width="30"
                                        height="30"
                                    >
                                        <path
                                            d="M844.15488 512.70656L571.58656 254.95552c-11.84768-11.84768-29.62432-11.84768-41.48224 0-11.84768 11.84768-11.84768 29.62432 0 41.472L749.34272 503.808H201.2672c-17.77664 0-29.62432 11.84768-29.62432 29.63456 0 17.7664 11.84768 29.62432 29.62432 29.62432h548.08576l-219.2384 207.38048c-11.84768 11.84768-11.84768 29.63456 0 41.48224 5.92896 5.92896 14.81728 8.88832 20.736 8.88832 5.92896 0 14.81728-2.95936 20.736-8.88832l272.55808-257.7408c5.92896-5.92896 8.88832-14.81728 8.88832-20.736 0.01024-5.92896-2.94912-14.81728-8.87808-20.74624z m0 0"
                                            fill="#ffffff"
                                            p-id="13529"
                                        />
                                    </svg>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>

                <div class="ml-28 z-10 px-2">
                    <div class="clock text-center z-10">
                        <p class="date z-10 select-none">{{ date }}</p>
                        <p class="time z-10 select-none">{{ time }}</p>
                    </div>
                    <div class="text-lg font-normal select-none pl-4">最近</div>
                    <div class="grid w-full place-items-center">
                        <div
                            v-for="item in rencent_items"
                            class="relative h-40 w-80 mb-2 rounded-md shadow-lg py-4 px-6 border text-base font-semibold hover_card cursor-pointer"
                        >
                            <div class="flex changetotext">
                                <div
                                    class="rencent_title w-full mr-5 select-none"
                                >{{item}}1211111111111111111111111131232131232131111111111111111111111111111</div>

                                <div class="rec_icon collectIconWrap">
                                    <el-tooltip
                                        class="item"
                                        effect="dark"
                                        content="收藏"
                                        placement="bottom"
                                    >
                                        <svg
                                            t="1655195283445"
                                            class="cardTitleImg cardTitleImgShow collectHover"
                                            viewBox="0 0 1024 1024"
                                            p-id="6924"
                                            width="16"
                                            height="16"
                                            data-spm-anchor-id="5176.28158796.0.i42.68656ac5pdme4W"
                                        >
                                            <path
                                                d="M529.066667 165.034667l104.533333 216.405333 206.378667 20.181333a42.666667 42.666667 0 0 1 23.808 74.666667l-162.176 140.885333 46.037333 224.384a42.666667 42.666667 0 0 1-64.298667 44.8l-192.725333-119.637333-192.597333 119.637333a42.666667 42.666667 0 0 1-64.298667-44.8l45.994667-224.384-162.133334-140.842666a42.666667 42.666667 0 0 1 23.808-74.666667L348.16 381.44l104.106667-216.32a42.666667 42.666667 0 0 1 76.8-0.085333z"
                                                p-id="6925"
                                                fill="currentColor"
                                            />
                                        </svg>
                                    </el-tooltip>
                                </div>
                                <div class="rec_icon bg-gray-50 iconHoverWrap">
                                    <svg
                                        t="1714551777145"
                                        class="icon"
                                        viewBox="0 0 1024 1024"
                                        version="1.1"
                                        xmlns="http://www.w3.org/2000/svg"
                                        p-id="4607"
                                        width="16"
                                        height="16"
                                    >
                                        <path
                                            d="M128 128v768h768V128H128z m853.333333-85.333333v938.666666H42.666667V42.666667h938.666666z"
                                            fill="#17abe3"
                                            p-id="4608"
                                        />
                                        <path
                                            d="M341.333333 341.333333l341.333334 170.666667-341.333334 170.666667z"
                                            fill="#17abe3"
                                            p-id="4609"
                                        />
                                    </svg>
                                </div>
                            </div>
                            <div class="mt-5 overflow-hidden flex-wrap flex select-none h-6">
                                <el-tag v-for="tag in recent_tag" class="flex mr-1">{{tag}}</el-tag>
                            </div>
                            <div
                                class="mt-4 flex justify-between text-xs font-normal text-gray-400 select-none"
                            >
                                <div>{{period}}</div>
                                <div>{{record_date}}</div>
                            </div>
                            <div class="absolute bottom-0 left-0 w-full z-1">
                                <div
                                    class="h-8 flex justify-between px-14 bg-white label_b opacity-0 cursor-default"
                                >
                                    <div class="text-black cursor-pointer" @click="rename">
                                        <el-tooltip
                                            class="item"
                                            effect="dark"
                                            content="重命名"
                                            placement="bottom"
                                        >
                                            <svg
                                                t="1655194314953"
                                                class="HoverImg"
                                                viewBox="0 0 1024 1024"
                                                p-id="5876"
                                                width="20"
                                                height="20"
                                            >
                                                <path
                                                    d="M532.394667 259.84l235.306666 220.245333-292.138666 291.712h341.290666V853.333333H170.666667v-0.042666h-2.858667l1.066667-227.968 363.52-365.482667zM750.933333 201.557333l5.845334 5.504 142.08 142.08-81.109334 80.896-235.52-220.416 2.56-2.56a121.6 121.6 0 0 1 166.144-5.504z"
                                                    p-id="5877"
                                                    fill="currentColor"
                                                />
                                            </svg>
                                        </el-tooltip>
                                    </div>

                                    <div class="text-black cursor-pointer">
                                        <el-tooltip
                                            class="item"
                                            effect="dark"
                                            content="导出"
                                            placement="bottom"
                                        >
                                            <svg
                                                t="1714569161423"
                                                class="icon"
                                                viewBox="0 0 1024 1024"
                                                version="1.1"
                                                xmlns="http://www.w3.org/2000/svg"
                                                p-id="11322"
                                                width="20"
                                                height="20"
                                            >
                                                <path
                                                    d="M512 640.64a42.666667 42.666667 0 0 0 42.666667-42.666667v-341.333333h130.986666a21.333333 21.333333 0 0 0 14.250667-5.461333l2.688-2.901334a21.333333 21.333333 0 0 0-4.010667-29.909333l-165.717333-126.464a32 32 0 0 0-38.912 0.042667L329.472 218.453333a21.333333 21.333333 0 0 0 12.970667 38.229334H469.333333v341.333333a42.666667 42.666667 0 0 0 42.666667 42.666667z m229.674667-298.368a42.666667 42.666667 0 0 0 4.992 85.034667H853.333333v426.666666H170.666667v-426.666666h106.666666a42.666667 42.666667 0 0 0 0-85.333334H170.666667a85.333333 85.333333 0 0 0-85.333334 85.333334v426.666666a85.333333 85.333333 0 0 0 85.333334 85.333334h682.666666a85.333333 85.333333 0 0 0 85.333334-85.333334v-426.666666a85.333333 85.333333 0 0 0-85.333334-85.333334h-106.666666z"
                                                    fill="#000000"
                                                    p-id="11323"
                                                />
                                            </svg>
                                        </el-tooltip>
                                    </div>
                                    <div class="text-black cursor-pointer">
                                        <el-tooltip
                                            class="item"
                                            effect="dark"
                                            content="删除"
                                            placement="bottom"
                                        >
                                            <svg
                                                t="1714569254710"
                                                class="icon"
                                                viewBox="0 0 1024 1024"
                                                version="1.1"
                                                xmlns="http://www.w3.org/2000/svg"
                                                p-id="12340"
                                                width="20"
                                                height="20"
                                            >
                                                <path
                                                    d="M360 184h-8c4.4 0 8-3.6 8-8v8h304v-8c0 4.4 3.6 8 8 8h-8v72h72v-80c0-35.3-28.7-64-64-64H352c-35.3 0-64 28.7-64 64v80h72v-72zM864 256H160c-17.7 0-32 14.3-32 32v32c0 4.4 3.6 8 8 8h60.4l24.7 523c1.6 34.1 29.8 61 63.9 61h454c34.2 0 62.3-26.8 63.9-61l24.7-523H888c4.4 0 8-3.6 8-8v-32c0-17.7-14.3-32-32-32zM731.3 840H292.7l-24.2-512h487l-24.2 512z"
                                                    p-id="12341"
                                                />
                                            </svg>
                                        </el-tooltip>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <!-- <div class="bg_gif">
        </div>-->
    </div>
</template>
 
<script>
var week = ['SUN', 'MON', 'TUE', 'WED', 'THU', 'FRI', 'SAT'];
export default {
    mounted() {
        setInterval(this.startPlay, 7000);
        setInterval(this.updateTime, 1000);
    },
    data() {
        return {
            ulList: [{ msg: '高效开会！' }, { msg: '轻松学习！' }, { msg: '随手总结！' }],
            play: false,
            rencent_items: ['你问我为何时常沉默，有的人无话可说，有的话无人可说.你问我为何时常沉默，有的人无话可说，有的话无人可说.', '2'],
            time: '',
            date: '',
            recent_tag: ['音乐', '音乐', 'dasdsa', '音乐dasd', '音乐dsa', '音乐ds'],
            period: '02:14',
            record_date: '04-25 16:11'
        };
    },
    methods: {
        startPlay() {
            let that = this;
            that.play = true; //开始播放
            setTimeout(() => {
                that.ulList.push(that.ulList[0]); //将第一条数据塞到最后一个
                that.ulList.shift(); //删除第一条数据
                that.play = false; //暂停播放,此处修改，保证每一次都会有动画显示
            }, 500);
        },
        updateTime() {
            let that = this;
            var cd = new Date();
            that.time =
                that.zeroPadding(cd.getHours(), 2) +
                ':' +
                that.zeroPadding(cd.getMinutes(), 2) +
                ':' +
                that.zeroPadding(cd.getSeconds(), 2);
            that.date =
                that.zeroPadding(cd.getFullYear(), 4) +
                '-' +
                that.zeroPadding(cd.getMonth() + 1, 2) +
                '-' +
                that.zeroPadding(cd.getDate(), 2) +
                ' ' +
                week[cd.getDay()];
        },
        zeroPadding(num, digit) {
            var zero = '';
            for (var i = 0; i < digit; i++) {
                zero += '0';
            }
            return (zero + num).slice(-digit);
        },
        rename() {
            console.log('111111');
            this.$prompt('重命名', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                inputPattern: /^[^\n]{1,150}$/,
                inputErrorMessage: '命名不能为空,不能带有回车，最大长度为150字符',
                inputType: 'textarea',
                inputValue: this.rencent_items[0]
            })
                .then(({ value }) => {
                    this.$message({
                        type: 'success',
                        message: '你的邮箱是: ' + value
                    });
                })
                .catch(() => {
                    this.$message({
                        type: 'info',
                        message: '取消输入'
                    });
                });
        }
    }
};
</script>
 
<style scoped>
.toUp {
    margin-top: -96px;
    transition: all 0.5s;
    z-index: 1;
}

.scroll {
    overflow: hidden;
}

.bg {
    /* background: url(//g.alicdn.com/idst-fe/mind-meeting-assistant2/0.0.204/static/mask.4122db2d.png) 0px 0px no-repeat; */
    background: url(../../assets/img/homepage.webp) 0px 0px no-repeat;
    height: 1200px;
    width: 100%;
    background-size: cover !important;
    position: absolute;
    user-select: none;
    filter: blur(2px);
}

.bg_gif {
    /* background: url(//g.alicdn.com/idst-fe/mind-meeting-assistant2/0.0.204/static/mask.4122db2d.png) 0px 0px no-repeat; */
    background: url(../../assets/img/home_gif.gif) 0px 0px no-repeat;
    background-size: 500px;
    position: absolute;
    top: 0px;
    right: 400px;
    width: 500px;
    height: 278px;
    user-select: none;
}

.hcard {
    width: 424px;
}

.hcard:hover .hover_svg {
    opacity: 1 !important;
    transition: all 0.2s ease-in 0s;
}

.meet_bg {
    background-image: url(../../assets/img/meeting.jpg);
    background-size: 100% 100%;
    user-select: none;
}

.class_bg {
    background-image: url(../../assets/img/class.png);
    background-size: 100% 100%;
    user-select: none;
}

.doc_bg {
    background-image: url(../../assets/img/doc.png);
    background-size: 100% 100%;
    user-select: none;
}

.live_bg {
    background-image: url(../../assets/img/live.png);
    background-size: 100% 100%;
    user-select: none;
}

.rencent_title {
    display: -webkit-box;
    line-height: 1.5;
    overflow: hidden;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 2;
    word-break: break-all;
    font-family: PingFangSC-Medium;
}

/* .recent_box{
background: #f1f1f1;
border-radius: 10px;
  background-image: 
    linear-gradient(#e1e1e1 0.1em, transparent 0.1em);
  background-size: 100% 30px;
} */

.clock {
    /* transform: translate(-50%, -50%); */
    color: #18aff5;
    text-shadow: 0 0 20px rgb(23, 170, 219), 0 0 20px rgba(169, 222, 240, 0.89);
}

.time {
    letter-spacing: 0.05em;
    font-size: 80px;
    padding: 5px 0;
    user-select: none;
}
.date {
    letter-spacing: 0.1em;
    font-size: 24px;
    user-select: none;
}

.rec_icon {
    display: flex;
    -webkit-box-align: center;
    align-items: center;
    -webkit-box-pack: justify;
    justify-content: space-between;
    height: 24px;
    padding: 0 5px;
    min-width: 24px;
    max-width: 24px;
    min-height: 24px;
    max-height: 24px;
    border-radius: 4px;
}

.label_b {
    transition: all 0.1s ease-in 0s;
}

.hover_card:hover .label_b {
    opacity: 1;
}
.collectIconWrap {
    display: none;
}

.hover_card:hover .collectIconWrap {
    display: flex;
}

.hover_card:hover .iconHoverWrap {
    display: none;
}
</style>
