<template>
    <div class="w-full h-full overflow-auto">
        <div class="bg"></div>

        <div class="right-0 top-60 z-10 fixed flex flex-col select-none" v-if="isUpload===true">
            <div class="relative flex flex-col">
                <div
                    class="h-14 w-64 flex px-4 items-center bg-white justify-between rounded-lg cursor-pointer mr-0 shadow-2xl border"
                    :class="[{'hoverTrans': (openTrans===false)},
                    {'closeTrans': (openTrans===false)}]"
                    @click="clickTrans"
                >
                    <div class="flex items-center">
                        <div class="mr-2">
                            <svg
                                t="1714943972617"
                                class="icon"
                                viewBox="0 0 1024 1024"
                                version="1.1"
                                xmlns="http://www.w3.org/2000/svg"
                                p-id="6806"
                                width="20"
                                height="20"
                            >
                                <path
                                    d="M512.005813 0C229.229613 0 0 229.229613 0 512.06394s229.229613 511.93606 512.005813 511.93606 511.982562-229.229613 511.982562-511.93606S794.770387 0 512.005813 0z m190.65654 475.804321a24.715598 24.715598 0 0 1-34.876197 0L536.663283 344.716324v393.577876a24.669096 24.669096 0 0 1-49.326566 0V344.716324L356.24872 475.804321a24.660958 24.660958 0 0 1-34.876196-34.876196l172.695298-172.672048a26.192023 26.192023 0 0 1 35.864356 0l172.683673 172.683673a24.715598 24.715598 0 0 1 0 34.864571z"
                                    p-id="6807"
                                    fill="#1296db"
                                />
                            </svg>
                        </div>
                        <div class>转写中</div>
                    </div>
                    <div>
                        <svg
                            t="1714907031563"
                            class="icon"
                            viewBox="0 0 1024 1024"
                            version="1.1"
                            xmlns="http://www.w3.org/2000/svg"
                            p-id="4545"
                            width="20"
                            height="20"
                        >
                            <path
                                d="M510.182608 762.39791 71.456449 261.260306 952.544575 261.260306Z"
                                fill="#8a8a8a"
                                p-id="4546"
                            />
                        </svg>
                    </div>
                </div>

                <div
                    class="top-full right-0 mt-2.5 ml-1.5 max-h-96 overflow-y-scroll w-80 absolute flex bg-white rounded-lg border"
                    v-if="openTrans===true"
                >
                    <div class="transListContainer relative w-full">
                        <div class="pt-5 pl-5 text-xs h-11">
                            <el-checkbox
                                :indeterminate="isIndeterminate"
                                v-model="checkAll"
                                @change="handleCheckAllChange"
                            >全选</el-checkbox>
                        </div>
                        <ul class="mb-0 pb-4 list-none">
                            <div
                                class="w-full inline-block m-0 p-0 leading-normal text-sm text-black text-opacity-80"
                            >
                                <li class="flex list-none m-0 pr-6 pl-5">
                                    <div class="flex pt-3.5">
                                        <div class="list-none mr-2">
                                            <el-checkbox></el-checkbox>
                                        </div>
                                        <div>
                                            <svg
                                                t="1714906781285"
                                                class="icon"
                                                viewBox="0 0 1024 1024"
                                                version="1.1"
                                                xmlns="http://www.w3.org/2000/svg"
                                                p-id="3240"
                                                width="20"
                                                height="20"
                                            >
                                                <path
                                                    d="M0 0h1024v1024H0V0z"
                                                    fill="#202425"
                                                    opacity=".01"
                                                    p-id="3241"
                                                />
                                                <path
                                                    d="M955.733333 512c0 245.077333-198.656 443.733333-443.733333 443.733333S68.266667 757.077333 68.266667 512 266.922667 68.266667 512 68.266667s443.733333 198.656 443.733333 443.733333z"
                                                    fill="#11AA66"
                                                    p-id="3242"
                                                />
                                                <path
                                                    d="M512 102.4C285.7984 102.4 102.4 285.7984 102.4 512s183.3984 409.6 409.6 409.6 409.6-183.3984 409.6-409.6S738.2016 102.4 512 102.4zM34.133333 512C34.133333 248.081067 248.081067 34.133333 512 34.133333s477.866667 213.947733 477.866667 477.866667-213.947733 477.866667-477.866667 477.866667S34.133333 775.918933 34.133333 512z"
                                                    fill="#11AA66"
                                                    p-id="3243"
                                                />
                                                <path
                                                    d="M787.114667 339.285333a51.2 51.2 0 0 1 0 72.362667l-307.2 307.2a51.2 51.2 0 0 1-72.362667 0l-170.666667-170.666667a51.2 51.2 0 0 1 72.362667-72.362666L443.733333 610.235733l271.018667-271.018666a51.2 51.2 0 0 1 72.362667 0z"
                                                    fill="#FFFFFF"
                                                    p-id="3244"
                                                />
                                            </svg>
                                        </div>
                                    </div>
                                    <div class="ml-2 py-3.5 flex-1 w-0">
                                        <div class="w-full flex justify-between items-center">
                                            <span class="overflow-hidden overflow-ellipsis">乱世巨星</span>
                                            <span class="cursor-pointer">
                                                <svg
                                                    t="1714943421170"
                                                    class="icon"
                                                    viewBox="0 0 1024 1024"
                                                    version="1.1"
                                                    xmlns="http://www.w3.org/2000/svg"
                                                    p-id="4042"
                                                    width="11"
                                                    height="11"
                                                >
                                                    <path
                                                        d="M551.358403 519.66047l402.416186 402.416186L914.421302 961.429942 512.002047 559.013757 109.578698 961.437106 70.224388 922.083819l402.423349-402.423349L70.221318 117.229958l39.35738-39.353286 402.423349 402.426419L914.425396 77.882811l39.353286 39.35738L551.358403 519.66047z"
                                                        fill="#3A2E2E"
                                                        p-id="4043"
                                                    />
                                                </svg>
                                            </span>
                                        </div>
                                        <div class="mt-1 text-xs text-gray-500">转写中，预计剩余...</div>
                                    </div>
                                </li>
                            </div>
                        </ul>
                    </div>
                </div>
            </div>
        </div>

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
            <div class="flex select-none">
                <div class="ml-20 z-10">
                    <div class="flex">
                        <router-link to="/voice-record">
                            <div
                                class="hcard transform transition duration-300 hover:scale-110 rounded-2xl shadow-xl h-64 hover:shadow-xl flex cursor-pointer bg-gradient-to-br bg-blue-100 hover:from-blue-300 hover:via-blue-400 hover:to-blue-500 text-black hover:text-white"
                            >
                                <div
                                    class="w-1/2 h-full meet_bg bg-cover bg-no-repeat rounded-l-2xl"
                                ></div>
                                <div class="w-1/2 h-full rounded-r-2xl p-4 bx_sd">
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
                            class="ml-10 hcard transform transition duration-300 hover:scale-110 rounded-2xl shadow-lg h-64 hover:shadow-xl flex cursor-pointer bg-purple-200 bg-gradient-to-br hover:from-purple-300 hover:via-purple-400 hover:to-purple-500 text-black hover:text-white"
                            @click="videoDialogVisible = true"
                        >
                            <div class="w-1/2 h-full class_bg bg-cover bg-no-repeat rounded-l-2xl"></div>
                            <div class="w-1/2 h-full rounded-r-2xl p-4 bx_sd">
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
                            class="transform transition duration-300 hover:scale-110 rounded-2xl shadow-lg h-64 hover:shadow-xl hcard flex cursor-pointer bg-pink-200 bg-gradient-to-br hover:from-pink-200 hover:via-pink-300 hover:to-pink-400 text-black hover:text-white"
                            @click="docUploadDialogVisible = true"
                        >
                            <div class="w-1/2 h-full doc_bg bg-cover bg-no-repeat rounded-l-2xl"></div>
                            <div class="w-1/2 h-full rounded-r-2xl p-4 bx_sd">
                                <div class="font-semibold text-2xl">文档、图片一键提取</div>
                                <div class="font-medium text-base leading-7 mt-3">文档、图片总结归纳</div>
                                <div class="font-medium text-base leading-7">拍照上传，轻松阅读</div>
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
                            class="ml-10 transform transition duration-300 hover:scale-110 rounded-2xl shadow-lg h-64 hover:shadow-xl hcard flex cursor-pointer bg-red-100 bg-gradient-to-br hover:from-red-200 hover:via-red-300 hover:to-red-400 text-black hover:text-white"
                        >
                            <div class="w-1/2 h-full live_bg bg-cover bg-no-repeat rounded-l-2xl"></div>

                            <div class="w-1/2 h-full rounded-r-2xl p-4 bx_sd">
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
                            :key="item.id"
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
                                <el-tag
                                    v-for="tag in recent_tag"
                                    :key="tag.id"
                                    class="flex mr-1"
                                >{{tag}}</el-tag>
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
        <el-dialog
            :visible="videoDialogVisible"
            title="上传音视频"
            width="30%"
            top="25vh"
            :close-on-click-modal="false"
            @close="closeVideoDialog"
            :destroy-on-close="true"
            class="select-none"
        >
            <div class="items-center">
                <div
                    class="videodialogcard flex items-center bg-purple-100 bg-opacity-95 hover:bg-purple-200 cursor-pointer"
                    @click="videoDialogVisible = false;videoUploadDialogVisible = true"
                >
                    <svg
                        t="1714675464716"
                        class="icon ml-5 mr-10"
                        viewBox="0 0 1024 1024"
                        version="1.1"
                        xmlns="http://www.w3.org/2000/svg"
                        p-id="6860"
                        width="60"
                        height="60"
                    >
                        <path
                            d="M960.711582 372.46606l0 523.083429c0 36.684504-29.738301 66.423829-66.422805 66.423829L130.420375 961.973318c-36.684504 0-66.422805-29.738301-66.422805-66.423829L63.99757 131.682112c0-36.684504 29.738301-66.422805 66.422805-66.422805l523.083429 0 0-0.261966 0.261966 0.261966 30.127158 0 1.151219 1.152243 1.671059-1.670036L956.784132 334.808392l-1.671059 1.671059 5.598508 5.597485 0 30.127158 0.261966 0.261966L960.711582 372.46606zM695.018314 131.753743l0 116.169813c0 45.855374 37.173644 83.029018 83.029018 83.029018l116.169813 0L695.018314 131.753743zM919.197072 372.46606 736.532823 372.46606c-45.855374 0-83.029018-37.173644-83.029018-83.029018L653.503804 106.772792 155.329694 106.772792c-27.513634 0-49.817616 22.303982-49.817616 49.817616l0 714.050786c0 27.513634 22.303982 49.817616 49.817616 49.817616l714.050786 0c27.513634 0 49.817616-22.303982 49.817616-49.817616L919.198096 372.46606zM861.077374 845.73085l-49.817616 0c-4.585435 0-8.303106-3.714602-8.303106-8.30106l0-49.819662c0-4.582365 3.717671-8.30106 8.303106-8.30106l49.817616 0c4.585435 0 8.303106 3.718695 8.303106 8.30106l0 49.819662C869.38048 842.016249 865.662809 845.73085 861.077374 845.73085zM861.077374 704.58418l-49.817616 0c-4.585435 0-8.303106-3.718695-8.303106-8.305153l0-49.814546c0-4.586458 3.717671-8.305153 8.303106-8.305153l49.817616 0c4.585435 0 8.303106 3.718695 8.303106 8.305153l0 49.814546C869.38048 700.865485 865.662809 704.58418 861.077374 704.58418zM861.077374 563.434439l-49.817616 0c-4.585435 0-8.303106-3.718695-8.303106-8.305153l0-49.814546c0-4.586458 3.717671-8.305153 8.303106-8.305153l49.817616 0c4.585435 0 8.303106 3.718695 8.303106 8.305153l0 49.814546C869.38048 559.715745 865.662809 563.434439 861.077374 563.434439zM654.699026 535.981181 436.936397 670.177556c-7.795547 4.804423-17.503653 7.468089-27.588335 8.499581-2.730181 0.60989-5.634324 0.9967-8.705266 0.9967-3.207042 0-6.245238-0.406253-9.075703-1.067308-1.24741-0.140193-2.524496-0.101307-3.756557-0.296759L387.810536 348.922854c1.231037-0.195451 2.509147-0.156566 3.756557-0.296759 2.830465-0.662079 5.868661-1.067308 9.075703-1.067308 3.070942 0 5.976108 0.387833 8.705266 0.9967 10.084683 1.032516 19.792788 3.695159 27.588335 8.499581l217.761606 134.195352C674.742525 503.602749 674.742525 523.628852 654.699026 535.981181zM600.258113 502.43311l-173.951821-107.197464 0 236.761332 173.951821-107.197464C610.280374 518.622838 610.280374 508.609787 600.258113 502.43311zM213.449393 845.73085l-49.817616 0c-4.585435 0-8.303106-3.714602-8.303106-8.30106l0-49.819662c0-4.582365 3.717671-8.30106 8.303106-8.30106l49.817616 0c4.585435 0 8.303106 3.718695 8.303106 8.30106l0 49.819662C221.7525 842.016249 218.035852 845.73085 213.449393 845.73085zM213.449393 704.58418l-49.817616 0c-4.585435 0-8.303106-3.718695-8.303106-8.305153l0-49.814546c0-4.586458 3.717671-8.305153 8.303106-8.305153l49.817616 0c4.585435 0 8.303106 3.718695 8.303106 8.305153l0 49.814546C221.7525 700.865485 218.035852 704.58418 213.449393 704.58418zM213.449393 563.434439l-49.817616 0c-4.585435 0-8.303106-3.718695-8.303106-8.305153l0-49.814546c0-4.586458 3.717671-8.305153 8.303106-8.305153l49.817616 0c4.585435 0 8.303106 3.718695 8.303106 8.305153l0 49.814546C221.7525 559.715745 218.035852 563.434439 213.449393 563.434439zM213.449393 413.982616l-49.817616 0c-4.585435 0-8.303106-3.719718-8.303106-8.306176l0-49.814546c0-4.586458 3.717671-8.306176 8.303106-8.306176l49.817616 0c4.585435 0 8.303106 3.719718 8.303106 8.306176l0 49.814546C221.7525 410.262898 218.035852 413.982616 213.449393 413.982616zM213.449393 272.831852l-49.817616 0c-4.585435 0-8.303106-3.719718-8.303106-8.302083l0-49.818639c0-4.586458 3.717671-8.302083 8.303106-8.302083l49.817616 0c4.585435 0 8.303106 3.715625 8.303106 8.302083l0 49.818639C221.7525 269.112134 218.035852 272.831852 213.449393 272.831852z"
                            p-id="6861"
                        />
                    </svg>
                    <div>
                        <p class="text-black text-lg font-medium">上传本地音视频文件</p>
                        <p class="text-sm text-gray-400">音视频转文字，提炼总结，一键导出</p>
                    </div>
                </div>

                <div
                    class="videodialogcard flex items-center bg-blue-100 bg-opacity-95 hover:bg-blue-200 cursor-pointer"
                    @click="videoDialogVisible = false;videoLinkUploadDialogVisible = true"
                >
                    <svg
                        t="1714675580159"
                        class="icon ml-5 mr-10"
                        viewBox="0 0 1024 1024"
                        version="1.1"
                        xmlns="http://www.w3.org/2000/svg"
                        p-id="8715"
                        width="60"
                        height="60"
                    >
                        <path
                            d="M913.937279 113.328092c-32.94432-32.946366-76.898391-51.089585-123.763768-51.089585s-90.819448 18.143219-123.763768 51.089585L416.737356 362.999454c-32.946366 32.94432-51.089585 76.898391-51.089585 123.763768s18.143219 90.819448 51.087539 123.763768c25.406646 25.40767 57.58451 42.144866 93.053326 48.403406 1.76418 0.312108 3.51915 0.463558 5.249561 0.463558 14.288424 0 26.951839-10.244318 29.519314-24.802896 2.879584-16.322757-8.016581-31.889291-24.339338-34.768875-23.278169-4.106528-44.38386-15.081487-61.039191-31.736818-21.61018-21.61018-33.509185-50.489928-33.509185-81.322144s11.899004-59.711963 33.509185-81.322144l15.864316-15.864316c-0.267083 1.121544-0.478907 2.267647-0.6191 3.440355-1.955538 16.45988 9.800203 31.386848 26.260084 33.344432 25.863041 3.072989 49.213865 14.378475 67.527976 32.692586 21.608134 21.608134 33.509185 50.489928 33.509185 81.322144s-11.901051 59.71401-33.509185 81.322144L318.53987 871.368764c-21.61018 21.61018-50.489928 33.511231-81.322144 33.511231-30.832216 0-59.711963-11.901051-81.322144-33.511231-21.61018-21.61018-33.509185-50.489928-33.509185-81.322144s11.899004-59.711963 33.509185-81.322144l169.43597-169.438017c11.720949-11.718903 11.720949-30.722722 0-42.441625-11.718903-11.718903-30.722722-11.718903-42.441625 0L113.452935 666.282852c-32.946366 32.94432-51.089585 76.898391-51.089585 123.763768 0 46.865377 18.143219 90.819448 51.089585 123.763768 32.94432 32.946366 76.898391 51.091632 123.763768 51.091632s90.819448-18.145266 123.763768-51.091632l249.673409-249.671363c32.946366-32.94432 51.089585-76.898391 51.089585-123.763768-0.002047-46.865377-18.145266-90.819448-51.089585-123.763768-27.5341-27.536146-64.073294-45.240367-102.885252-49.854455-3.618411-0.428765-7.161097-0.196475-10.508331 0.601704l211.589023-211.589023c21.61018-21.61018 50.489928-33.509185 81.322144-33.509185s59.711963 11.899004 81.322144 33.509185c21.61018 21.61018 33.509185 50.489928 33.509185 81.322144s-11.899004 59.711963-33.509185 81.322144l-150.180418 150.182464c-11.720949 11.718903-11.720949 30.722722 0 42.441625 11.718903 11.718903 30.722722 11.718903 42.441625 0l150.180418-150.182464c32.946366-32.94432 51.089585-76.898391 51.089585-123.763768C965.026864 190.226482 946.882622 146.272411 913.937279 113.328092z"
                            fill="#272636"
                            p-id="8716"
                        />
                    </svg>
                    <div>
                        <p class="text-black text-lg font-medium select-none">输入解析音视频链接</p>
                        <p class="text-sm text-gray-400 select-none">自动解析，自动总结</p>
                    </div>
                </div>
            </div>
        </el-dialog>

        <el-dialog
            :visible="videoUploadDialogVisible"
            title="上传音视频"
            width="30%"
            top="30vh"
            :close-on-click-modal="false"
            @close="closeVideoUploadDialog"
            class="select-none"
            :destroy-on-close="true"
        >
            <div class="flex">
                <el-upload
                    drag
                    class="upload-demo"
                    action="#"
                    ref="upload"
                    multiple
                    :limit="5"
                    :auto-upload="false"
                    :on-exceed="handleExceed"
                    :file-list="fileList"
                    :http-request="handleFileUpload"
                    :on-change="handleChange"
                    :on-remove="handleRemove"
                    accept=".mp3, .mp4"
                >
                    <i class="el-icon-upload"></i>
                    <div class="el-upload__text">
                        将文件拖到此处，或
                        <em>点击上传</em>
                    </div>
                    <div class="el-upload__tip flex justify-between" slot="tip">
                        <div>目前仅支持mp4，mp3文件，文件时长最长为1小时</div>
                        <div>{{fileCount}}/5</div>
                    </div>
                </el-upload>
                <div class="ml-5">
                    <div class="mb-5">
                        <div>音视频语言</div>
                        <div>
                            <el-radio-group v-model="lang" @change="changeLang">
                                <el-radio-button label="0">中文</el-radio-button>
                                <el-radio-button label="1">英语</el-radio-button>
                            </el-radio-group>
                        </div>
                    </div>
                    <div class="mb-5">
                        <div class>翻译</div>
                        <el-select v-model="transValue">
                            <el-option
                                v-for="item in uploadOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"
                            ></el-option>
                        </el-select>
                        <div></div>
                    </div>
                    <div class="flex justify-center">
                        <el-button class="w-1/2" type="primary">
                            上传
                            <i class="el-icon-upload el-icon--right"></i>
                        </el-button>
                    </div>
                </div>
            </div>
        </el-dialog>

        <el-dialog
            :visible="videoLinkUploadDialogVisible"
            title="上传音视频链接"
            width="30%"
            top="30vh"
            :close-on-click-modal="false"
            @close="closeVideoLinkUploadDialog"
            class="select-none"
            :destroy-on-close="true"
        >
            <div>
                <el-input
                    :autosize="{ minRows: 12, maxRows: 12}"
                    resize="none"
                    v-model="videoUrl"
                    min="1"
                    type="textarea"
                    placeholder="请输入音视频链接"
                ></el-input>
            </div>

            <div class="flex justify-center mt-5">
                <el-button class="w-1/4" type="primary">
                    上传
                    <i class="el-icon-upload el-icon--right"></i>
                </el-button>
            </div>
        </el-dialog>

        <el-dialog
            :visible="docUploadDialogVisible"
            title="上传文档、图片"
            width="25%"
            top="30vh"
            :close-on-click-modal="false"
            @close="closeDocUploadDialog"
            class="select-none"
            :destroy-on-close="true"
        >
            <div class="flex justify-center -mt-10">
                <el-menu
                    :default-active="docUploadDialogAcitveIndex"
                    class="el-menu-demo"
                    background-color="#fff"
                    mode="horizontal"
                    @select="handleDocTypeSelect"
                    style="border-bottom: 0px !important;"
                >
                    <el-menu-item index="0">文档阅读</el-menu-item>
                    <el-menu-item index="1">网页阅读</el-menu-item>
                </el-menu>
            </div>
            <div class="w-full flex justify-center mt-2" v-if="docUploadDialogAcitveIndex==='0'">
                <el-upload
                    drag
                    class="upload-demo"
                    action="#"
                    ref="upload"
                    multiple
                    :limit="1"
                    :file-list="fileList"
                    :http-request="handleFileUpload"
                    accept=".pdf, .doc, .docx, .img, .png, jepg"
                >
                    <i class="el-icon-upload"></i>
                    <div class="el-upload__text">
                        将文件拖到此处，或
                        <em>点击上传</em>
                    </div>
                    <div class="el-upload__tip" slot="tip">
                        <div class="flex justify-between mb-2">
                            <div class="w-5/6">
                                <el-input
                                    v-model="pdfUrl"
                                    min="1"
                                    placeholder="输入包含文档的URL，后缀需为XXX.pdf"
                                ></el-input>
                            </div>
                            <div>
                                <el-button>添加</el-button>
                            </div>
                        </div>
                        <div class="flex items-center mb-2">
                            <div class="text-gray-400">文件类型</div>
                            <div class="w-5/6">
                                <el-select v-model="docType" class="ml-3 w-full">
                                    <el-option
                                        v-for="item in docTypeOptions"
                                        :key="item.value"
                                        :label="item.label"
                                        :value="item.value"
                                    ></el-option>
                                </el-select>
                            </div>
                        </div>
                        <div class="text-gray-400">文件格式：PDF/ Word/ 图片/ HTML</div>
                        <div class="text-gray-400">文档大小：文件最大支持100M，图片最大支持20M</div>
                        <div class="text-gray-400">文档页数：PDF/Word最多支持100页</div>
                    </div>
                </el-upload>
            </div>

            <div v-else class="select-none">
                <div class="mb-3">
                    <el-input
                        :autosize="{ minRows: 12, maxRows: 12}"
                        resize="none"
                        v-model="htmlUrl"
                        min="1"
                        type="textarea"
                        placeholder="请输入或粘贴网址"
                    ></el-input>
                </div>

                <div class="text-gray-400 mb-3 text-xs">单次上传一个网址，请确保添加网址符合网络规范</div>
                <div class="flex justify-center">
                    <el-button class="w-1/4" type="primary">
                        上传
                        <i class="el-icon-upload el-icon--right"></i>
                    </el-button>
                </div>
            </div>
        </el-dialog>
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
            record_date: '04-25 16:11',
            videoDialogVisible: false,
            videoUploadDialogVisible: false,
            videoLinkUploadDialogVisible: false,
            docUploadDialogVisible: false,
            fileList: [],
            fileCount: 0,
            lang: 0,
            transValue: 0,
            videoUrl: '',
            htmlUrl: '',
            docType: 0,
            docUploadDialogAcitveIndex: '0',
            pdfUrl: '',
            openTrans: false,
            isUpload: false,
            uploadOptions: [
                {
                    value: 0,
                    label: '不翻译'
                },
                {
                    value: 1,
                    label: '英语'
                }
            ],
            docTypeOptions: [
                {
                    value: 0,
                    label: '文档'
                },
                {
                    value: 1,
                    label: '论文'
                },
                {
                    value: 2,
                    label: '图书'
                }
            ]
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
        },
        closeVideoDialog() {
            this.videoDialogVisible = false;
        },
        closeVideoUploadDialog() {
            this.videoUploadDialogVisible = false;
            this.uploadOptions[1].label = '英语';
            this.lang = 0;
            this.transValue = 0;
            this.fileCount = 0;
        },
        closeVideoLinkUploadDialog() {
            this.videoLinkUploadDialogVisible = false;
            this.link = '';
        },
        closeDocUploadDialog() {
            this.docUploadDialogVisible = false;
        },
        // 处理移除操作
        handleChange(file, fileList) {
            this.fileCount = fileList.length;
        },
        handleRemove(file, fileList) {
            this.fileCount = fileList.length;
        },
        // 处理超出图片个数操作
        handleExceed(files, fileList) {
            this.$message.warning(
                `当前限制选择 5 个文件，本次选择了 ${files.length} 个文件，共选择了 ${files.length + fileList.length} 个文件`
            );
        },
        // 处理文件上传操作
        handleFileUpload(file) {
            this.loading = true;
            // 调用后端服务器的接口
            uploadFile(file.file)
                .then(resp => {
                    this.form.installImgUrl = resp.url;
                })
                .catch(e => {
                    this.$message.error(e.message);
                    this.$refs.upload.clearFiles();
                });
        },
        submitUpload() {
            this.$refs.upload.submit();
        },
        changeLang() {
            if (this.lang == '0') {
                this.uploadOptions[1].label = '英语';
            } else {
                this.uploadOptions[1].label = '中文';
            }
        },
        isValidUrl() {
            const regex = /^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/)?[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$/;
            return regex.test(this.link);
        },
        handleDocTypeSelect(key, keyPath) {
            this.docUploadDialogAcitveIndex = key;
            this.pdfUrl = '';
            this.htmlUrl = '';
            console.log(this.docUploadDialogAcitveIndex);
        },
        clickTrans() {
            this.openTrans = !this.openTrans;
        },
        handleCheckAllChange(val) {
            this.checkedCities = val ? cityOptions : [];
            this.isIndeterminate = false;
        },
        handleCheckedCitiesChange(value) {
            let checkedCount = value.length;
            this.checkAll = checkedCount === this.cities.length;
            this.isIndeterminate = checkedCount > 0 && checkedCount < this.cities.length;
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

.videodialogcard {
    height: 110px;
    width: 85%;
    margin-top: 20px;
    margin-left: auto;
    margin-right: auto;
    border-radius: 14px;
    padding: 8px;
}

::v-deep .el-radio-button__inner {
    width: 100px;
    height: 36px;
    border: 1 !important;
    border-radius: 2px;
    font-size: 14px;
    font-weight: 400;
    color: #696969;
    line-height: 14px;
    outline: none;
    box-shadow: none;
}
/*激活样式*/
::v-deep .el-radio-button__orig-radio:checked + .el-radio-button__inner {
    background: #74baff;
    border: 0 !important;
    color: #ffffff;
    line-height: 14px;
    outline: none;
    box-shadow: none;
}

.el-menu--horizontal >>> .el-menu-item:not(.is-disabled):hover {
    background: transparent !important;
}

.hoverTrans:hover {
    margin-right: 0;
    transition: all 0.2s ease-in 0s;
}

.closeTrans {
    margin-right: -128px;
    transition: all 0.2s ease-in 0s;
}

.transListContainer {
    transition: opacity 0.3s;
}
</style>
