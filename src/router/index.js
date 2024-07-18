import Vue from 'vue';
import Router from 'vue-router';

Vue.use(Router);
export default new Router({
    routes: [
        {
            path: '/',
            redirect: '/home'
        },
        {
            path: '/',
            component: () => import(/* webpackChunkName: "home" */ '../components/common/Home.vue'),
            meta: { title: '自述文件' },
            children: [
                {
                    path: '/home',
                    component: () => import(/* webpackChunkName: "home" */ '../components/page/home.vue'),
                    meta: { title: '系统首页' }
                },
                // {
                //     path: '/livelist',
                //     component: () => import(/* webpackChunkName: "livelist" */ '../components/page/LiveList.vue'),
                //     meta: { title: '直播列表' }
                // },
                {
                    path: '/404',
                    component: () => import(/* webpackChunkName: "404" */ '../components/page/404.vue'),
                    meta: { title: '404' }
                },
                {
                    path: '/403',
                    component: () => import(/* webpackChunkName: "403" */ '../components/page/403.vue'),
                    meta: { title: '403' }
                },
            ]
        },
        {
            path: '/live',
            component: () => import(/* webpackChunkName: "live" */ '../components/page/Live.vue'),
            meta: { title: '直播'}
        },
        {
            path: '/voice-record',
            component: () => import(/* webpackChunkName: "voice-record" */ '../components/page/VoiceRecord.vue'),
            meta: { title: '语音记录'}
        },
        {
            path: '/transcript',
            component: () => import(/* webpackChunkName: "transcript" */ '../components/page/Transcript.vue'),
            meta: { title: '记录'}
        },
        {
            path: '/subtitle',
            component: () => import(/* webpackChunkName: "transcript" */ '../components/page/Subtitle.vue'),
            meta: { title: '记录'}
        },
        {
            path: '*',
            redirect: '/404'
        },
    ]
});
