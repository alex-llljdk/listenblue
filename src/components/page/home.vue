<template>
    <div class="w-full h-full bg-gradient-to-br from-white via-gray-100 to-gray-200 overflow-auto">
        <div class="bg"></div>
        <div class="pl-10 pr-10 pb-10 z-1">
            <div class="w-full h-32 flex px-5 py-4 text-6xl font-semibold mt-20">
                <div class="flex h-full z-10">
                    <el-text>ListenBlue,&nbsp&nbsp</el-text>
                </div>
                <div class="scroll h-full z-10">
                    <div
                        v-for="(item, index) in ulList"
                        :key="item.id"
                        :class="!index && play ? 'toUp' : ''"
                        class="truncate h-full text-blue-500"
                    >{{ item.msg }}</div>
                </div>
            </div>
            <div class="flex">
                <div>
                    <div class="flex">
                        <div
                            class="hcard transform transition duration-300 hover:scale-110 rounded-2xl shadow-lg h-64 w-96 hover:shadow-xl opacity-50"
                        >
                            <div
                                class="bg-gradient-to-br flex from-blue-100 via-blue-300 to-blue-500 h-full rounded-2xl"
                            >
                                <div
                                    class="w-1/2 h-full meet_bg bg-cover bg-no-repeat rounded-l-2xl"
                                ></div>
                            </div>
                        </div>
                        <div
                            class="ml-10 hcard transform transition duration-300 hover:scale-110 rounded-2xl shadow-lg h-64 w-96 hover:shadow-xl bg-white opacity-30"
                        >
                            <div
                                class="bg-gradient-to-br from-purple-100 via-purple-300 to-purple-500 h-full rounded-2xl"
                            >
                                <div
                                    class="w-1/2 h-full class_bg bg-cover bg-no-repeat rounded-l-2xl"
                                ></div>
                            </div>
                        </div>
                    </div>
                    <div class="flex mt-5">
                        <div
                            class="transform transition duration-300 hover:scale-110 rounded-2xl shadow-lg h-64 w-96 hover:shadow-xl bg-white opacity-30 hcard"
                        >
                            <div
                                class="bg-gradient-to-br from-pink-100 via-pink-200 to-pink-300 h-full rounded-2xl"
                            >
                                <div
                                    class="w-1/2 h-full doc_bg bg-cover bg-no-repeat rounded-l-2xl"
                                ></div>
                            </div>
                        </div>
                        <div
                            class="ml-10 transform transition duration-300 hover:scale-110 rounded-2xl shadow-lg h-64 w-96 hover:shadow-xl bg-white opacity-30 hcard"
                        >
                            <div
                                class="bg-gradient-to-br from-red-100 via-red-300 to-red-500 h-full rounded-2xl"
                            >
                                <div
                                    class="w-1/2 h-full live_bg bg-cover bg-no-repeat rounded-l-2xl"
                                ></div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="flex mt-5">
                <div class="ml-10 text-base">
                    最近
                    <div class="mt-5">
                        <div
                            v-for="item in rencent_items"
                            class="h-32 w-72 bg-white mb-6 rounded-2xl"
                        >{{item}}</div>
                    </div>
                </div>
            </div>
            </div>
            
        </div>
    </div>
</template>
 
<script>
export default {
    mounted() {
        setInterval(this.startPlay, 7000);
    },
    data() {
        return {
            ulList: [{ msg: '高效开会！' }, { msg: '轻松学习！' }, { msg: '随手总结！' }],
            play: false,
            rencent_items: ['1', '2', '3']
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
    background: url(../../assets/img/homepage_bg.png) 0px 0px no-repeat;
    background-size: 395px;
    position: absolute;
    top: 0px;
    right: 0px;
    width: 395px;
    height: 512px;
    user-select: none;
}

.hcard:hover {
    opacity: 1;
}

.meet_bg {
    background-image: url(../../assets/img/meeting.jpg);
    background-size: 100% 100%;
    opacity: 1;
    user-select: none;
}

.class_bg {
    background-image: url(../../assets/img/class.png);
    background-size: 100% 100%;
    opacity: 1;
    user-select: none;
}

.doc_bg {
    background-image: url(../../assets/img/doc.png);
    background-size: 100% 100%;
    opacity: 1;
    user-select: none;
}

.live_bg {
    background-image: url(../../assets/img/live.png);
    background-size: 100% 100%;
    opacity: 1;
    user-select: none;
}
</style>
