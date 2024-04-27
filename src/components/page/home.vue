<template>
    <div>
        <div class="bg"></div>
        <div class="pl-10 pr-10 pb-10 z-1">
        <div class="w-full h-32 flex px-5 py-4 text-6xl font-semibold mt-20">
            <div class="flex h-full">
                <el-text>ListenBlue,&nbsp&nbsp</el-text>
            </div>
            <div class="scroll h-full">
                <div
                    v-for="(item, index) in ulList"
                    :key="item.id"
                    :class="!index && play ? 'toUp' : ''"
                    class="truncate h-full text-blue-500"
                >{{ item.msg }}</div>
            </div>
        </div>

        <el-row class="mt-20">
            <el-col :span="8" v-for="(o, index) in 2" :key="o" :offset="index > 0 ? 2 : 0">
                <div
                    class="transform transition duration-300 hover:scale-110 rounded-lg shadow-lg h-52 w-112 hover:shadow-xl bg-white"
                >
                    <div
                        class="bg-gradient-to-br from-pink-100 via-purple-200 to-purple-300 m-2 h-3/6 rounded-lg"
                    ></div>

                    <div class="px-5 pt-2 flex flex-col">
                        <h2 class="font-semibold">Title</h2>
                    </div>
                </div>
            </el-col>
        </el-row>
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
            play: false
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
}

.scroll {
    overflow: hidden;
}

.bg{
    background: url(//g.alicdn.com/idst-fe/mind-meeting-assistant2/0.0.204/static/mask.4122db2d.png) 0px 0px no-repeat;
    position: absolute;
    top: 0px;
    right: 0px;
    width: 1440px;
    height: 810px;
}

/* 隐藏滚动条 */
::-webkit-scrollbar {
    display: none;
}
</style>
