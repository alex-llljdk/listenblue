<template>
    <div class="livelist-container h-full w-full">
        <div class="livelist-head flex px-10 pt-5 items-center border-b select-none">
            <div class="text-4xl">
                <p>我的记录</p>
            </div>

            <div class="search">
                <div class="flex items-center justify-center p-5">
                    <div class="rounded-lg border">
                        <div class="flex">
                            <div
                                class="flex items-center justify-center rounded-tl-lg rounded-bl-lg border-r p-2"
                            >
                                <svg t="1721373336081" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="6258" width="16" height="16"><path d="M953.474215 908.234504l-152.576516-163.241391c61.92508-74.48211 95.81186-167.36973 95.81186-265.073744 0-229.294809-186.63531-415.930119-416.102133-415.930119-229.294809 0-415.930119 186.63531-415.930119 415.930119s186.63531 415.930119 415.930119 415.930119c60.032925 0 118.00168-12.55703 172.186125-37.327062 16.169326-7.396607 23.221905-26.318159 15.825298-42.315471-7.396607-16.169326-26.318159-23.221905-42.315471-15.825298-45.927768 20.813707-94.951789 31.478582-145.695952 31.478582-194.031917 0-351.94087-157.908953-351.94087-351.94087 0-194.031917 157.908953-351.94087 351.94087-351.94087 194.031917 0 351.94087 157.908953 351.94087 351.94087 0 91.339493-34.918864 177.86259-98.048043 243.743995-12.213002 12.729044-11.868974 33.026709 0.860071 45.239711 1.032085 0.860071 2.236183 1.204099 3.268268 2.064169 0.860071 1.204099 1.376113 2.752226 2.408198 3.956325l165.477574 177.00252c6.192508 6.70855 14.793214 10.148833 23.393919 10.148833 7.912649 0 15.653284-2.92424 21.845792-8.600706C964.827146 941.433227 965.515202 921.135562 953.474215 908.234504z" fill="#575B66" p-id="6259"></path></svg>
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

        <div class="flex justify-center w-full place-items-center">
            <div class="flex flex-wrap px-12">
                <div
                    v-for="(item, index) in recordList"
                    :key="index"
                    class="w-80 mb-2 rounded-md shadow-lg mx-4 my-4 py-4 px-6 border text-base font-semibold hover_card cursor-pointer"
                >
                    <div @click="clickToRecord(item.id)">
                        <div class="flex changetotext">
                            <div class="rencent_title w-full mr-5 select-none" v-if="item.dataType > 1">
                                {{ item.title.answer }}
                            </div>
                            <div class="rencent_title w-full mr-5 select-none" v-else>{{ item.title }}</div>
                            <div class="rec_icon collectIconWrap">
                                <el-tooltip class="item" effect="dark" content="收藏" placement="bottom">
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
                            <div class="rec_icon bg-gray-50 iconHoverWrap" v-if="item.dataType === 0">
                                <svg
                                    t="1721300751329"
                                    class="icon"
                                    viewBox="0 0 1024 1024"
                                    version="1.1"
                                    xmlns="http://www.w3.org/2000/svg"
                                    p-id="4260"
                                    width="16"
                                    height="16"
                                >
                                    <path
                                        d="M849.408 6.656L411.648 140.8c-53.248 15.36-95.744 70.656-95.744 123.392v461.312S284.16 704 213.504 714.24C109.568 729.088 25.6 808.448 25.6 891.904s83.968 134.656 187.904 119.808c103.936-14.848 179.712-91.648 179.712-175.104v-445.44c0-36.864 44.544-52.736 44.544-52.736l387.072-121.344s43.008-14.336 43.008 25.088v367.616s-39.424-22.528-110.08-14.336c-103.936 12.8-187.904 90.624-187.904 174.08S653.824 905.728 757.76 893.44c103.936-12.8 187.904-90.624 187.904-174.08V74.752c-0.512-52.224-43.52-82.944-96.256-68.096z"
                                        fill="#1296db"
                                        p-id="4261"
                                    ></path>
                                </svg>
                            </div>
                            <div class="rec_icon bg-gray-50 iconHoverWrap" v-if="item.dataType === 1">
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
                            <div class="rec_icon bg-gray-50 iconHoverWrap" v-if="item.dataType > 1">
                                <svg
                                    t="1721216794948"
                                    class="icon"
                                    viewBox="0 0 1024 1024"
                                    version="1.1"
                                    xmlns="http://www.w3.org/2000/svg"
                                    p-id="4319"
                                    width="16"
                                    height="16"
                                >
                                    <path
                                        d="M284.434286 971.428571h455.131428c88.722286 0 132.864-45.001143 132.864-134.144V448.146286c0-55.296-5.997714-79.286857-40.283428-114.432L595.565714 93.275429c-32.548571-33.426286-59.136-40.704-107.574857-40.704H284.452571c-88.283429 0-132.864 45.421714-132.864 134.582857v650.130285c0 89.563429 44.580571 134.144 132.864 134.144z m3.419428-68.992c-44.141714 0-67.291429-23.588571-67.291428-66.432V188.434286c0-42.422857 23.149714-66.870857 67.730285-66.870857h190.281143v249.014857c0 53.997714 27.428571 80.566857 80.566857 80.566857h244.297143v384.859428c0 42.843429-22.729143 66.432-67.291428 66.432zM566.857143 386.413714c-16.713143 0-23.588571-6.838857-23.588572-23.990857V134.857143l246.857143 251.574857z"
                                        p-id="4320"
                                        fill="#1296db"
                                    ></path>
                                </svg>
                            </div>
                        </div>
                        <div class="mt-5 overflow-hidden flex-wrap flex select-none h-6" v-if="item.dataType > 1">
                            <el-tag v-for="tag in item.keyword.answer" :key="tag.id" class="flex mr-1">{{ tag.desc }}</el-tag>
                        </div>
                        <div class="mt-5 overflow-hidden flex-wrap flex select-none h-6" v-else>
                            <el-tag v-for="tag in item.keyword.answer" :key="tag.id" class="flex mr-1">{{ tag }}</el-tag>
                        </div>
                    </div>
                    <div class="mt-4 flex justify-between text-xs font-normal text-gray-400 select-none">
                        <div></div>
                        <div>{{ item.updatetime }}</div>
                    </div>
                    <div class="bottom-0 left-0 w-full z-1">
                        <div class="h-8 flex justify-between px-14 bg-white label_b opacity-0 cursor-default">
                            <div class="text-black cursor-pointer" @click="rename">
                                <el-tooltip class="item" effect="dark" content="重命名" placement="bottom">
                                    <svg t="1655194314953" class="HoverImg" viewBox="0 0 1024 1024" p-id="5876" width="20" height="20">
                                        <path
                                            d="M532.394667 259.84l235.306666 220.245333-292.138666 291.712h341.290666V853.333333H170.666667v-0.042666h-2.858667l1.066667-227.968 363.52-365.482667zM750.933333 201.557333l5.845334 5.504 142.08 142.08-81.109334 80.896-235.52-220.416 2.56-2.56a121.6 121.6 0 0 1 166.144-5.504z"
                                            p-id="5877"
                                            fill="currentColor"
                                        />
                                    </svg>
                                </el-tooltip>
                            </div>

                            <div class="text-black cursor-pointer">
                                <el-tooltip class="item" effect="dark" content="导出" placement="bottom">
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
                                <el-tooltip class="item" effect="dark" content="删除" placement="bottom">
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
import { userStore } from '../../store/store';
import { getRecordList } from '../../api/record';
import { fileToBuffer, encodeWithSalt, decodeWithSalt, Salt } from '../../utils/util';
const userstore = userStore();
export default {
    mounted() {
        //获取我的所有记录
        getRecordList(userstore.user_id).then((res) => {
            console.log(res);
            this.recordList = res.singlerecords;
            for (var i = 0; i < this.recordList.length; i++) {
                if (this.recordList[i].dataType > 1) {
                    this.recordList[i].title = JSON.parse(this.recordList[i].title);
                    this.recordList[i].keyword = JSON.parse(this.recordList[i].keyword);
                    this.recordList[i].updatetime = new Date(this.recordList[i].updatetime).toISOString().slice(0, 19).replace('T', ' ');
                } else {
                    this.recordList[i].title = this.recordList[i].title;
                    this.recordList[i].keyword = JSON.parse(this.recordList[i].keyword);
                    this.recordList[i].updatetime = new Date(this.recordList[i].updatetime).toISOString().slice(0, 19).replace('T', ' ');
                }
            }
            console.log(this.recordList);
        });
    },
    methods: {
        goBack() {
            this.$router.go(-1);
        },
        rename() {
            if (!userstore.isLogin) {
                this.loginDialogVisible = true;
                return;
            }
            this.$prompt('重命名', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                inputPattern: /^[^\n]{1,150}$/,
                inputErrorMessage: '命名不能为空,不能带有回车，最大长度为150字符',
                inputType: 'textarea',
                inputValue: this.recent_items[0]
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
        },
        clickToRecord(record_id) {
            if (!userstore.isLoggedIn) {
                this.loginDialogVisible = true;
                return;
            }
            this.$router.push({
                path: '/transcript',
                query: { id: encodeWithSalt(record_id, Salt) }
            });
        }
    },
    data() {
        return {
            recordList: []
        };
    }
};
</script>


<style scoped>
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
.rencent_title {
    display: -webkit-box;
    line-height: 1.5;
    overflow: hidden;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 2;
    word-break: break-all;
    font-family: PingFangSC-Medium;
}
</style>
