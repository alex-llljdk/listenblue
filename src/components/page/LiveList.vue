<template>
    <div class="livelist-container h-full w-full">
        <div class="livelist-head flex px-10 pt-5 items-center border-b select-none">
            <div class="text-4xl">
                <p>在线直播</p>
            </div>

            <div class="search">
                <div class="flex items-center justify-center p-5">
                    <div class="rounded-lg border">
                        <div class="flex">
                            <div
                              class="flex w-10 items-center justify-center rounded-tl-lg rounded-bl-lg border-r border-gray-200 bg-white p-5"
                            >
                                <svg
                                    viewBox="0 0 20 20"
                                    aria-hidden="true"
                                    class="pointer-events-none absolute w-5 fill-gray-500 transition"
                                >
                                    <path
                                        d="M16.72 17.78a.75.75 0 1 0 1.06-1.06l-1.06 1.06ZM9 14.5A5.5 5.5 0 0 1 3.5 9H2a7 7 0 0 0 7 7v-1.5ZM3.5 9A5.5 5.5 0 0 1 9 3.5V2a7 7 0 0 0-7 7h1.5ZM9 3.5A5.5 5.5 0 0 1 14.5 9H16a7 7 0 0 0-7-7v1.5Zm3.89 10.45 3.83 3.83 1.06-1.06-3.83-3.83-1.06 1.06ZM14.5 9a5.48 5.48 0 0 1-1.61 3.89l1.06 1.06A6.98 6.98 0 0 0 16 9h-1.5Zm-1.61 3.89A5.48 5.48 0 0 1 9 14.5V16a6.98 6.98 0 0 0 4.95-2.05l-1.06-1.06Z"
                                    />
                                </svg>
                            </div>
                            <input
                                type="text"
                                class="w-full max-w-[160px] bg-white pl-2 text-base font-semibold outline-0"
                                placeholder
                                id
                            />
                            <input
                                type="button"
                                value="Search"
                                class="bg-blue-500 p-2 rounded-tr-lg rounded-br-lg text-white font-semibold hover:bg-blue-800 transition-colors cursor-pointer"
                            />
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <div class="livelist-body flex w-full h-4/5">
            <ul class="w-full px-4 pt-10">
                <li v-for="(item,index) in LiveList" :key="index" class="w-1/5 px-4 mb-5 inline-block cursor-pointer" >
                    <div @click="clickJoinRoom(item.room_number)">
                        <el-card :body-style="{ padding: '0px' }">
                            <img
                                src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                                class="w-full h-48 block"
                                v-if="item.cover_url===''"
                            />
                            <img
                                :src="item.cover_url"
                                class="w-full h-48 block"
                                v-else
                            />
                            <div class="p-4 flex w-full">
                                <div class="avatar w-1/5" v-if="item.avatar===''">
                                    <el-avatar
                                        src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
                                    ></el-avatar>
                                </div>
                                <div class="avatar w-1/5" v-else>
                                    <el-avatar
                                        :src="item.avatar"
                                    ></el-avatar>
                                </div>
                                <div class="title w-3/5">
                                    <div>{{item.room_title}}</div>
                                    <div>{{item.user_name}}</div>
                                </div>
                                <div class="see_num w-1/5 flex justify-end items-center">
                                    <span>{{item.room_number}}</span>
                                </div>
                            </div>
                        </el-card>
                    </div>
                </li>
            </ul>
        </div>

        <div class="livelist-footer w-full flex justify-center">
            <div class="block">
                <el-pagination
                    @size-change="handleSizeChange"
                    @current-change="handleCurrentChange"
                    :current-page.sync="currentPage3"
                    :page-size="100"
                    layout="prev, pager, next, jumper"
                    :total="1000"
                ></el-pagination>
            </div>
        </div>
    </div>
</template>

<script>
import { GetLiveList} from '../../api/live';
import { JoinRoom } from '../../api/live';
import { encodeWithSalt,Salt } from '../../utils/util';

export default {
    mounted(){
        GetLiveList().then((res)=>{
            console.log(res)
            this.LiveList = res.live_list
        })
    },
    data() {
        return {
            LiveList: []
        };
    },
    methods:{
        clickJoinRoom(roomNumber) {
            JoinRoom(roomNumber, "").then((res) => {
                if (res.code === 200) {
                    this.$router.push({
                        path: '/live',
                        query: { rn: roomNumber, pw: encodeWithSalt("", Salt) }
                    });
                }
            });
        },
    }
};
</script>


<style scoped>
.avatar {
    display: block;
    width: 50px;
    height: 50px;
    border-radius: 50%;
}
</style>
