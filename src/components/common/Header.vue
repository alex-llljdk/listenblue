<template>
    <div class="flex items-center justify-between gap-2 border-b border-gray-400 pl-8 pr-2 col-span-2 h-24">
        <!-- 折叠按钮 -->
        <div class="flex items-center px-2">
            <!-- <div class="px-4" @click="collapseChage">
                <i v-if="!collapse" class="el-icon-s-fold"></i>
                <i v-else class="el-icon-s-unfold"></i>
            </div> -->
            <img src="../../assets/logo.png" />
        </div>

        <div class="h-full">
            <div class="flex items-center h-full">
                <!-- 消息中心 -->
                <!-- <div class="btn-bell">
                    <el-tooltip
                        effect="dark"
                        :content="message?`有${message}条未读消息`:`消息中心`"
                        placement="bottom"
                    >
                        <router-link to="/tabs">
                            <i class="el-icon-bell"></i>
                        </router-link>
                    </el-tooltip>
                    <span class="btn-bell-badge" v-if="message"></span>
                </div> -->
                <!-- 用户头像 -->
                <div class="flex items-center mr-6 select-none" v-if="userLogin">
                    <div class="user-avator">
                        <img :src="imageUrl" />
                    </div>
                    <!-- 用户名下拉菜单 -->
                    <el-dropdown class="user-name" trigger="click" @command="handleCommand">
                        <span class="el-dropdown-link text-black">
                            {{ username }}
                            <i class="el-icon-caret-bottom"></i>
                        </span>
                        <el-dropdown-menu slot="dropdown">
                            <el-dropdown-item divided command="updateAvatar">修改头像</el-dropdown-item>
                            <el-dropdown-item divided command="loginout">退出登录</el-dropdown-item>
                        </el-dropdown-menu>
                    </el-dropdown>
                </div>
                <div class="flex mr-6 cursor-pointer border-l-2 border-r-2 h-1/3 items-center text-blue-400 px-5 select-none" @click="login" v-else>
                    登录
                </div>
                <!-- 全屏显示 -->
                <div class="btn-fullscreen" @click="handleFullScreen">
                    <el-tooltip effect="dark" :content="fullscreen ? `取消全屏` : `全屏`" placement="bottom">
                        <i class="el-icon-rank"></i>
                    </el-tooltip>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
import bus from '../common/bus';
import Logo from '../../assets/logo.png';
import { userStore } from '../../store/store';
import { Message } from 'element-ui';
import home from "../page/home.vue"
const userstore = userStore();
export default {
    data() {
        return {
            collapse: false,
            fullscreen: false,
            name: 'linxin',
            message: 2
        };
    },
    computed: {
        userLogin() {
            return userstore.isLoggedIn;
        },
        username() {
            return userstore.name;
        },
        imageUrl() {
            return userstore.avatar;
        }
    },
    methods: {
        // 用户名下拉菜单选择事件
        login(){
            bus.$emit('openLogin', true);
        },
        handleCommand(command) {
            if (command == 'loginout') {
                Message.success('退出登录');
                userstore.isLoggedIn = false;
                userstore.user_id = -1;
                userstore.avatar = '';
                userstore.name = '';
                userstore.token = '';
                window.location.reload();
            }
        },
        // 侧边栏折叠
        collapseChage() {
            this.collapse = !this.collapse;
            bus.$emit('collapse', this.collapse);
        },
        // 全屏事件
        handleFullScreen() {
            let element = document.documentElement;
            if (this.fullscreen) {
                if (document.exitFullscreen) {
                    document.exitFullscreen();
                } else if (document.webkitCancelFullScreen) {
                    document.webkitCancelFullScreen();
                } else if (document.mozCancelFullScreen) {
                    document.mozCancelFullScreen();
                } else if (document.msExitFullscreen) {
                    document.msExitFullscreen();
                }
            } else {
                if (element.requestFullscreen) {
                    element.requestFullscreen();
                } else if (element.webkitRequestFullScreen) {
                    element.webkitRequestFullScreen();
                } else if (element.mozRequestFullScreen) {
                    element.mozRequestFullScreen();
                } else if (element.msRequestFullscreen) {
                    // IE11
                    element.msRequestFullscreen();
                }
            }
            this.fullscreen = !this.fullscreen;
        }
    },
    mounted() {
        if (document.body.clientWidth < 1500) {
            this.collapseChage();
        }
    }
};
</script>
<style scoped>
.header {
    position: relative;
    box-sizing: border-box;
    width: 100%;
    height: 70px;
    font-size: 22px;
    color: #fff;
    min-width: 1920px !important;
    max-width: 1920px !important;
}
.collapse-btn {
    float: left;
    padding: 0 21px;
    cursor: pointer;
}
/* .header .logo {
    float: left;
    width: 250px;
    line-height: 70px;
}
.header-right {
    float: right;
    padding-right: 50px;
} */
.header-user-con {
    display: flex;
    height: 70px;
    align-items: center;
}
.btn-fullscreen {
    transform: rotate(45deg);
    margin-right: 5px;
    font-size: 24px;
}
.btn-bell,
.btn-fullscreen {
    position: relative;
    width: 30px;
    height: 30px;
    text-align: center;
    border-radius: 15px;
    cursor: pointer;
}
.btn-bell-badge {
    position: absolute;
    right: 0;
    top: -2px;
    width: 8px;
    height: 8px;
    border-radius: 4px;
    background: #f56c6c;
    color: #fff;
}
.btn-bell .el-icon-bell {
    color: #fff;
}
.user-name {
    margin-left: 10px;
}
.user-avator {
    margin-left: 20px;
}
.user-avator img {
    display: block;
    width: 50px;
    height: 50px;
    border-radius: 50%;
}
.el-dropdown-link {
    color: #161515;
    cursor: pointer;
}
.el-dropdown-menu__item {
    text-align: center;
}
</style>
