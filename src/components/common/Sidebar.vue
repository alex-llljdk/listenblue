<template>
    <div class="sidebar">
        <el-menu
            class="sidebar-el-menu py-3"
            :default-active="onRoutes"
            :collapse="collapse"
            background-color="#ffffff"
            text-color="#000000"
            font = 16px
            unique-opened
            router
        >

    <div class="py-2 px-3">
    <template>
     <el-text class="mb-2 select-none px-4 text-xl font-semibold tracking-tight text-primary">探索</el-text>
    </template>
        <!-- active-text-color="#20a0ff" -->
        <div class="space-y-2">
            <template v-for="item in explore_items">
                <template v-if="item.subs">
                    <el-submenu :index="item.index" :key="item.index">
                        <template slot="title">
                                <i :class="item.icon"></i>
                                <span slot="title">{{ item.title }}</span>
                        </template>
                        <template v-for="subItem in item.subs">
                            <el-submenu
                                v-if="subItem.subs"
                                :index="subItem.index"
                                :key="subItem.index"
                            >
                                <template slot="title">{{ subItem.title }}</template>
                                <el-menu-item
                                    v-for="(threeItem,i) in subItem.subs"
                                    :key="i"
                                    :index="threeItem.index"
                                >{{ threeItem.title }}</el-menu-item>
                            </el-submenu>
                            <el-menu-item
                                v-else
                                :index="subItem.index"
                                :key="subItem.index"
                            >{{ subItem.title }}</el-menu-item>
                        </template>
                    </el-submenu>
                </template>
                <template v-else>
                        <el-menu-item :index="item.index" :key="item.index"  class="font-medium h-11 px-8 text-base">
                        <i :class="item.icon"></i>
                        <span slot="title" class="font-medium text-base">{{ item.title }}</span>
                        </el-menu-item>
                </template>   
            </template>
            </div>
        </div>

 <div class="py-2 px-3">
    <template>
     <el-text class="mb-2 select-none px-4 text-xl font-semibold tracking-tight text-primary">我的</el-text>
    </template>
        <!-- active-text-color="#20a0ff" -->
        <div class="space-y-2">
            <template v-for="item in my_items">
                <template v-if="item.subs">
                    <el-submenu :index="item.index" :key="item.index">
                        <template slot="title">
                                <i :class="item.icon"></i>
                                <span slot="title">{{ item.title }}</span>
                        </template>
                        <template v-for="subItem in item.subs">
                            <el-submenu
                                v-if="subItem.subs"
                                :index="subItem.index"
                                :key="subItem.index"
                            >
                                <template slot="title">{{ subItem.title }}</template>
                                <el-menu-item
                                    v-for="(threeItem,i) in subItem.subs"
                                    :key="i"
                                    :index="threeItem.index"
                                >{{ threeItem.title }}</el-menu-item>
                            </el-submenu>
                            <el-menu-item
                                v-else
                                :index="subItem.index"
                                :key="subItem.index"
                            >{{ subItem.title }}</el-menu-item>
                        </template>
                    </el-submenu>
                </template>
                <template v-else>
                        <el-menu-item :index="item.index" :key="item.index" class=" h-11 px-8 ">
                        <i :class="item.icon"></i>
                        <span slot="title" class="font-medium text-base">{{ item.title }}</span>
                        </el-menu-item>
                </template>   
            </template>
            </div>
        </div>
        </el-menu>
    </div>
</template>

<script>
import bus from '../common/bus';
export default {
    data() {
        return {
            collapse: false,
            explore_items: [
                {
                    icon: 'el-icon-lx-home',
                    index: 'dashboard',
                    title: '发现'
                },
                {
                    icon: 'el-icon-lx-cascades',
                    index: 'table',
                    title: '会议纪要'
                },
                {
                    icon: 'el-icon-lx-cascades',
                    index: 'table',
                    title: '学习归纳'
                },
                {
                    icon: 'el-icon-lx-cascades',
                    index: 'table',
                    title: '有声阅读'
                },
                {
                    icon: 'el-icon-lx-copy',
                    index: 'tabs',
                    title: '直播',
                },
               
            ],
            my_items: [ 
                {
                    icon: 'el-icon-lx-calendar',
                    index: 'share',
                    title: '分享',
                },
                {
                    icon: 'el-icon-lx-calendar',
                    index: 'share',
                    title: '笔记',
                },
                {
                    icon: 'el-icon-pie-chart',
                    index: 'delete',
                    title: '历史记录'
                },
                {
                    icon: 'el-icon-lx-redpacket_fill',
                    index: '/donate',
                    title: '收藏'
                },
                {
                    icon: 'el-icon-pie-chart',
                    index: 'delete',
                    title: '通知'
                }]
        };
    },
    computed: {
        onRoutes() {
            return this.$route.path.replace('/', '');
        }
    },
    created() {
        // 通过 Event Bus 进行组件间通信，来折叠侧边栏
        bus.$on('collapse', msg => {
            this.collapse = msg;
            bus.$emit('collapse-content', msg);
        });
    }
};
</script>

<style scoped>
.sidebar {
    display: block;
    position: absolute;
    left: 0;
    top: 97px;
    bottom: 0;
    overflow-y: scroll;
}
.sidebar::-webkit-scrollbar {
    width: 0;
}
.sidebar-el-menu:not(.el-menu--collapse) {
    width: 250px;
}
.sidebar > ul {
    height: 100%;
}
.el-menu-item{
    border-radius: 30px;
    margin: 2px 10px 2px 10px;
    height: 2.75rem;
    display: flex;
    align-items: center; /* 垂直居中 */
}
.el-menu-item:hover{
    background-color: hsl(210deg 37.5% 96.86%) !important;
}
.el-menu-item.is-active{
    background-color: hsl(210deg 37.5% 96.86%) !important;
}
</style>
