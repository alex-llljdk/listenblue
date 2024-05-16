<template>
    <div class="h-full w-full flex">
        <div class="h-full select-none" style="background: #e9efff">
            <el-aside width="80px">
                <router-link to="/home">
                    <div class="text-center mb-8 mt-8">
                        <div class="rounded-xl bg-white mx-4 py-3 flex justify-center">
                            <svg
                                t="1715682947773"
                                class="icon"
                                viewBox="0 0 1024 1024"
                                version="1.1"
                                xmlns="http://www.w3.org/2000/svg"
                                p-id="14609"
                                width="20"
                                height="20"
                            >
                                <path
                                    d="M238.208798 504.759085c114.078175 125.190257 330.334497 341.639984 441.040878 454.927143L785.790178 856.32207c-36.308951-36.902468-343.031681-348.783685-344.618828-358.506117 62.496379-64.083527 246.805082-251.967662 324.579423-339.65784l-95.228875-93.844342L238.208798 504.759085z"
                                    fill="#272636"
                                    p-id="14610"
                                />
                            </svg>
                        </div>
                    </div>
                </router-link>
                <div class="aside-icon text-center mb-8 block px-4 cursor-pointer">
                    <div class="flex justify-center mb-2">
                        <svg
                            t="1715676163293"
                            class="icon"
                            viewBox="0 0 1024 1024"
                            version="1.1"
                            xmlns="http://www.w3.org/2000/svg"
                            p-id="4267"
                            width="30"
                            height="30"
                        >
                            <path
                                d="M925.248 356.928l-258.176-258.176a64 64 0 0 0-45.248-18.752H144a64 64 0 0 0-64 64v736a64 64 0 0 0 64 64h736a64 64 0 0 0 64-64V402.176a64 64 0 0 0-18.752-45.248zM288 144h192V256H288V144z m448 736H288V736h448v144z m144 0H800V704a32 32 0 0 0-32-32H256a32 32 0 0 0-32 32v176H144v-736H224V288a32 32 0 0 0 32 32h256a32 32 0 0 0 32-32V144h77.824l258.176 258.176V880z"
                                p-id="4268"
                            />
                        </svg>
                    </div>
                    <div class="text-sm">
                        <p>保存</p>
                    </div>
                </div>
                <div class="aside-icon text-center mb-8 block px-4 cursor-pointer">
                    <div class="flex justify-center mb-2">
                        <svg
                            t="1715676434283"
                            class="icon"
                            viewBox="0 0 1024 1024"
                            version="1.1"
                            xmlns="http://www.w3.org/2000/svg"
                            p-id="10671"
                            width="25"
                            height="25"
                        >
                            <path
                                d="M463.082667 0v551.338667l-189.162667-189.141334-71.424 71.424L463.082667 694.186667v0.384l47.850666 51.136 48.597334-51.072v-0.426667l256.853333-256.213333-71.338667-72.149334-185.514666 185.514667V0z"
                                fill="#000000"
                                p-id="10672"
                            />
                            <path
                                d="M921.28 690.965333v229.696H101.333333V691.584H0v331.050667h101.333333v-1.237334h819.946667v0.64h101.333333V690.965333z"
                                fill="#000000"
                                p-id="10673"
                            />
                        </svg>
                    </div>
                    <div class="text-sm">
                        <p>导出</p>
                    </div>
                </div>
                <div class="aside-icon text-center mb-8 block px-4 cursor-pointer">
                    <div class="flex justify-center mb-2">
                        <svg
                            t="1715682811765"
                            class="icon"
                            viewBox="0 0 1024 1024"
                            version="1.1"
                            xmlns="http://www.w3.org/2000/svg"
                            p-id="12809"
                            width="30"
                            height="30"
                        >
                            <path
                                d="M454.613333 96.682667l-110.122666 223.146666-246.272 35.797334a64 64 0 0 0-35.456 109.141333l178.176 173.696-42.069334 245.290667a64 64 0 0 0 6.442667 40.576l3.114667 5.333333a64 64 0 0 0 83.328 21.546667L512 835.413333l220.245333 115.797334a64 64 0 0 0 92.885334-67.456l-42.069334-245.290667 178.176-173.653333a64 64 0 0 0-35.456-109.226667l-246.272-35.754667-110.08-223.146666a64 64 0 0 0-114.816 0zM512 173.226667l91.008 184.362666 3.328 6.144a85.333333 85.333333 0 0 0 60.928 40.533334l203.434667 29.525333-147.2 143.573333-4.821334 5.077334a85.333333 85.333333 0 0 0-19.712 70.442666l34.730667 202.624-201.813333-106.069333-4.778667-2.133333a42.666667 42.666667 0 0 0-34.986667 2.133333L290.304 855.466667l34.773333-202.624 0.896-6.912a85.333333 85.333333 0 0 0-25.429333-68.608l-147.242667-143.573334 203.52-29.525333a85.333333 85.333333 0 0 0 64.213334-46.677333L512 173.226667z"
                                fill="#000000"
                                p-id="12810"
                            />
                        </svg>
                    </div>
                    <div class="text-sm">
                        <p>收藏</p>
                    </div>
                </div>
                <div class="aside-icon text-center mb-8 block px-4 cursor-pointer">
                    <div class="flex justify-center mb-2">
                        <svg
                            t="1715683066737"
                            class="icon"
                            viewBox="0 0 1024 1024"
                            version="1.1"
                            xmlns="http://www.w3.org/2000/svg"
                            p-id="15596"
                            width="30"
                            height="30"
                        >
                            <path
                                d="M214.601143 298.642286v598.601143c0 22.893714 18.578286 41.398857 41.398857 41.398857h513.243429a41.398857 41.398857 0 0 0 41.398857-41.398857V298.715429H214.601143zM769.243429 1024H256a126.902857 126.902857 0 0 1-126.756571-126.683429V213.284571H896v683.885715c0 69.924571-56.832 126.756571-126.756571 126.756571z m212.114285-725.357714H42.642286A42.788571 42.788571 0 0 1 0 256c0-23.405714 19.236571-42.642286 42.642286-42.642286h938.715428c23.405714 0 42.642286 19.163429 42.642286 42.642286 0 23.405714-19.163429 42.642286-42.642286 42.642286zM768 213.357714H682.642286V128a42.715429 42.715429 0 0 0-42.642286-42.642286h-256a42.715429 42.715429 0 0 0-42.642286 42.642286v85.357714H256V128C256 57.417143 313.417143 0 384 0h256c70.582857 0 128 57.417143 128 128v85.357714z m-384 597.284572a42.788571 42.788571 0 0 1-42.642286-42.642286V469.357714c0-23.478857 19.163429-42.715429 42.642286-42.715428 23.405714 0 42.642286 19.236571 42.642286 42.715428V768c0 23.405714-19.163429 42.642286-42.642286 42.642286z m256 0a42.788571 42.788571 0 0 1-42.642286-42.642286V469.357714c0-23.478857 19.163429-42.715429 42.642286-42.715428 23.405714 0 42.642286 19.236571 42.642286 42.715428V768c0 23.405714-19.163429 42.642286-42.642286 42.642286z"
                                fill="#3e3e3e"
                                p-id="15597"
                            />
                        </svg>
                    </div>
                    <div class="text-sm">
                        <p>删除</p>
                    </div>
                </div>
            </el-aside>
        </div>
        <div class="flex h-full w-full" style="background: #f2f5fb">
            <div class="contentContainer w-full px-8 pt-6">
                <div class="contentHeader w-full pb-2">
                    <div class="flex justify-between w-full">
                        <div>
                            <div class="hover_input focus_input">
                                <input
                                    type="text"
                                    class="hip fip pl-2 font-semibold text-lg bg-transparent border-2 border-blue-50 border-opacity-30 focus:border-2 focus:outline-none rounded-md max-w-lg"
                                    v-model="recordInfo"
                                    :style="{ width: inputWidth + 'px' }"
                                    @input="handleInput"
                                />
                            </div>
                            <div>
                                <span class="text-xs text-gray-400 select-none ml-2">{{ savedTime }}</span>
                            </div>
                        </div>

                        <div class="flex text-lg items-center">
                            <el-popover placement="bottom-end" width="280" trigger="hover" visible-arrow="false">
                                <div class="p-2 select-none">
                                    <div class="font-semibold text-black text-sm leading-6">翻译</div>
                                    <div class="flex items-center mb-2">
                                        <div class="text-xs whitespace-nowrap mr-3">选择语言</div>
                                        <el-select v-model="transType" placeholder="请选择">
                                            <el-option
                                                v-for="item in transOptions"
                                                :key="item.value"
                                                :label="item.label"
                                                :value="item.value"
                                            >
                                            </el-option>
                                        </el-select>
                                    </div>
                                    <div class="flex items-center">
                                        <div class="text-xs mr-3">显示语言</div>
                                        <el-radio-group v-model="showType">
                                            <el-radio-button label="0">双语显示</el-radio-button>
                                            <el-radio-button label="1">纯译文显示</el-radio-button>
                                        </el-radio-group>
                                    </div>
                                </div>
                                <div
                                    class="tools-icon cursor-pointer w-10 h-10 rounded-lg flex items-center justify-center mr-3"
                                    slot="reference"
                                >
                                    <svg
                                        t="1715703178965"
                                        class="icon"
                                        viewBox="0 0 1024 1024"
                                        version="1.1"
                                        xmlns="http://www.w3.org/2000/svg"
                                        p-id="35183"
                                        width="30"
                                        height="30"
                                    >
                                        <path
                                            d="M889.362963 259.792593c-11.377778-64.474074-60.681481-113.777778-125.155556-125.155556C680.77037 121.362963 595.437037 113.777778 512 113.777778s-168.77037 7.585185-252.207407 20.859259c-64.474074 11.377778-113.777778 60.681481-125.155556 125.155556C121.362963 343.22963 113.777778 428.562963 113.777778 512s7.585185 168.77037 20.859259 252.207407c11.377778 64.474074 60.681481 113.777778 125.155556 125.155556 83.437037 13.274074 166.874074 20.859259 252.207407 20.859259 83.437037 0 168.77037-7.585185 254.103704-20.859259 64.474074-11.377778 113.777778-60.681481 125.155555-125.155556 11.377778-83.437037 18.962963-166.874074 18.962963-252.207407 0-83.437037-7.585185-166.874074-20.859259-252.207407z m-551.822222 30.34074c34.133333 24.651852 60.681481 51.2 83.437037 75.851852l-24.651852 24.651852c-20.859259-24.651852-49.303704-49.303704-83.437037-75.851852l24.651852-24.651852z m15.17037 422.874074l-13.274074-30.34074c7.585185-7.585185 11.377778-15.17037 11.377778-22.755556V468.385185h-70.162963v-34.133333h102.4V644.740741c15.17037-13.274074 30.340741-30.340741 47.407407-49.303704l9.481482 36.02963c-26.548148 30.340741-56.888889 56.888889-87.22963 81.54074z m365.985185-62.577777h-123.259259v83.437037h-34.133333v-83.437037h-113.777778v-34.133334h113.777778v-54.992592h-92.918519v-32.237037h92.918519v-45.511111h34.133333v45.511111h98.607407v32.237037h-98.607407v54.992592h123.259259v34.133334z m-3.792592-163.081482c-54.992593-7.585185-100.503704-20.859259-136.533334-41.718518-37.925926 18.962963-83.437037 36.02963-138.429629 47.407407l-17.066667-28.444444c49.303704-11.377778 91.022222-24.651852 125.155556-39.822223-28.444444-22.755556-49.303704-53.096296-62.577778-87.229629H455.111111v-30.340741h233.244445v28.444444c-15.17037 34.133333-39.822222 64.474074-77.748149 91.022223 32.237037 15.17037 70.162963 24.651852 115.674074 30.34074l-11.377777 30.340741z"
                                            p-id="35184"
                                            fill="#2c2c2c"
                                        />
                                        <path
                                            d="M650.42963 337.540741h-132.740741c13.274074 28.444444 34.133333 51.2 60.681481 70.162963 32.237037-20.859259 56.888889-43.614815 72.05926-70.162963z"
                                            p-id="35185"
                                            fill="#2c2c2c"
                                        />
                                    </svg>
                                </div>
                            </el-popover>
                            <el-tooltip class="item" effect="dark" content="展开视频" placement="bottom" v-if="!isVideoOpen">
                                <div
                                    class="tools-icon cursor-pointer w-10 h-10 rounded-lg flex items-center justify-center mr-3"
                                    @click="changeVideo"
                                >
                                    <svg
                                        t="1715702308488"
                                        class="icon"
                                        viewBox="0 0 1024 1024"
                                        version="1.1"
                                        xmlns="http://www.w3.org/2000/svg"
                                        p-id="34101"
                                        data-spm-anchor-id="a313x.search_index.0.i8.41e13a81zzRDDe"
                                        width="26"
                                        height="26"
                                    >
                                        <path
                                            d="M793.6 64H230.4C140.8 64 64 140.8 64 230.4v563.2C64 883.2 140.8 960 230.4 960h563.2c89.6 0 166.4-76.8 166.4-166.4V230.4C960 140.8 883.2 64 793.6 64z m-390.4 614.4V288l339.2 198.4-339.2 192z"
                                            p-id="34102"
                                            data-spm-anchor-id="a313x.search_index.0.i9.41e13a81zzRDDe"
                                        />
                                    </svg>
                                </div>
                            </el-tooltip>

                            <el-tooltip class="item" effect="dark" content="收起视频" placement="bottom" v-else>
                                <div
                                    class="tools-icon cursor-pointer w-10 h-10 rounded-lg flex items-center justify-center mr-3"
                                    @click="changeVideo"
                                >
                                    <svg
                                        t="1715702308488"
                                        class="icon"
                                        viewBox="0 0 1024 1024"
                                        version="1.1"
                                        xmlns="http://www.w3.org/2000/svg"
                                        p-id="34101"
                                        data-spm-anchor-id="a313x.search_index.0.i8.41e13a81zzRDDe"
                                        width="26"
                                        height="26"
                                    >
                                        <path
                                            d="M793.6 64H230.4C140.8 64 64 140.8 64 230.4v563.2C64 883.2 140.8 960 230.4 960h563.2c89.6 0 166.4-76.8 166.4-166.4V230.4C960 140.8 883.2 64 793.6 64z m-390.4 614.4V288l339.2 198.4-339.2 192z"
                                            p-id="34102"
                                            fill="#0552f7"
                                            data-spm-anchor-id="a313x.search_index.0.i9.41e13a81zzRDDe"
                                        />
                                    </svg>
                                </div>
                            </el-tooltip>

                            <el-popover placement="bottom-end" width="260" trigger="hover" visible-arrow="false">
                                <div class="p-1">
                                    <div class="font-semibold text-black text-sm leading-6">AI改写</div>
                                    <div class="font-normal text-xs text-gray-400 leading-5">
                                        在保留原文的要点内容基础上，将原文内容进行精简和改写
                                    </div>
                                    <div class="flex aiRadio mt-2">
                                        <div class="whitespace-nowrap">显示内容：</div>
                                        <el-radio-group v-model="aiType">
                                            <el-radio :label="0" class="mb-1">改写结果</el-radio>
                                            <el-radio :label="1">原文和改写结果</el-radio>
                                        </el-radio-group>
                                    </div>
                                    <div class="w-full mt-3">
                                        <el-button type="primary">改写总结</el-button>
                                    </div>
                                </div>
                                <div
                                    class="tools-icon cursor-pointer w-10 h-10 rounded-lg flex items-center justify-center"
                                    slot="reference"
                                >
                                    <svg
                                        t="1715704188840"
                                        class="icon"
                                        viewBox="0 0 1024 1024"
                                        version="1.1"
                                        xmlns="http://www.w3.org/2000/svg"
                                        p-id="42896"
                                        width="28"
                                        height="28"
                                    >
                                        <path
                                            d="M841.774545 116.363636A65.861818 65.861818 0 0 1 907.636364 182.225455v659.54909A65.861818 65.861818 0 0 1 841.774545 907.636364H182.225455A65.861818 65.861818 0 0 1 116.363636 841.774545V182.225455A65.861818 65.861818 0 0 1 182.225455 116.363636h659.54909z m-99.909818 289.582546H673.978182c-4.491636 0-5.632 1.675636-5.632 5.608727l0.581818 340.642909c0 3.351273 1.117091 5.585455 5.585455 5.585455h67.351272c4.491636 0 5.608727-1.675636 5.608728-5.585455V411.554909c0-4.491636-1.675636-5.585455-5.585455-5.585454zM462.429091 281.367273h-90.344727c-2.792727 0-3.933091 1.117091-3.933091 4.491636-0.558545 21.317818-2.234182 28.625455-4.491637 35.909818l-122.88 429.288728c-1.140364 5.632 0 6.749091 4.468364 6.74909h62.859636c4.491636 0 5.608727-0.558545 6.749091-4.491636l34.210909-129.629091h5.632c9.821091 0.046545 31.325091 0.139636 54.202182 0.186182l12.567273 0.023273 12.427636 0.023272h28.253091c14.917818 0 25.134545-0.093091 24.971637-0.232727l35.909818 129.629091c1.117091 2.792727 2.792727 3.933091 6.749091 3.933091h71.819636c3.933091 0 5.050182-1.698909 3.933091-5.608727L468.014545 285.858909c-1.117091-3.374545-2.234182-4.491636-5.608727-4.491636z m-44.893091 69.585454c11.217455 50.501818 39.842909 154.88 51.618909 202.589091h-103.237818c18.501818-66.792727 41.518545-152.087273 50.501818-202.589091z m290.117818-80.802909c-24.133818 0-43.752727 16.290909-43.752727 43.752727 0 28.625455 17.384727 44.357818 42.635636 43.776 26.368 0 44.916364-15.127273 44.916364-43.752727 0-25.832727-17.408-43.776-43.799273-43.776z"
                                            fill="#000000"
                                            p-id="42897"
                                        />
                                    </svg>
                                </div>
                            </el-popover>
                        </div>
                    </div>
                </div>
                <div class="contentBody overflow-auto">
                    <div class="w-full h-96 px-2" v-if="isVideoOpen">
                        <d-player id="dplayer" class="h-full" :options="playerOptions"></d-player>
                    </div>

                    <div class="scrollContainer overflow-auto" :style="{ height: isVideoOpen ? 'calc(100% - 384px)' : '100%' }">
                        <div class="dSJFEj" v-for="(item_recording, index) in recording_items" :key="index">
                            <div class="jgzocF">
                                <div class="jVabAE">
                                    <svg t="1714715933780" class="icon" viewBox="0 0 1024 1024" width="25" height="25">
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
                                <div class="gnYZJU" :class="item_recording.isActive ? 'gnYZJU_focus ring-1 ring-purple-300' : 'gnYZJU'">
                                    <div class="jItVgd">
                                        <div class="kCofWf">
                                            <div
                                                class="kFJVLr w-full"
                                                contenteditable
                                                @blur="updateContent(index, $event)"
                                                @focusin="inDiv(index, $event)"
                                                @focusout="outDiv(index, $event)"
                                            >
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="h-full w-1/2" style="background: #f2f5fb">
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
import VueDPlayer from 'vue-dplayer';
import 'vue-dplayer/dist/vue-dplayer.css';
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
    components: { quillEditor, 'd-player': VueDPlayer },
    data() {
        return {
            // recording_items 数组用于存储要动态生成的 div 元素的数据
            recording_items: [{ time_stamp: this.formatTime(this.currentTime), content: '1111', isActive: false },
                { time_stamp: this.formatTime(this.currentTime), content: '1111', isActive: false },
                { time_stamp: this.formatTime(this.currentTime), content: '1111', isActive: false },
                { time_stamp: this.formatTime(this.currentTime), content: '1111', isActive: false },
                { time_stamp: this.formatTime(this.currentTime), content: '1111', isActive: false },
                { time_stamp: this.formatTime(this.currentTime), content: '1111', isActive: false },
                { time_stamp: this.formatTime(this.currentTime), content: '1111', isActive: false }
            ],
            item_index: -1, // recording_items当前最大进行时item的下标
            currentTime: 0, // 当前时间，单位：秒
            maxTime: 3600, // 最大时间，单位：秒，即 01:00:00
            timer: null, // 定时器
            inputWidth: 204,
            recordInfo: '',
            savedTime: '',
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

            isVideoOpen: true,
            aiType: 0,
            transType: 0,
            showType: 0,
            transOptions: [
                {
                    value: 0,
                    label: '源语言'
                },
                {
                    value: 1,
                    label: '中英互译'
                }
            ],
            playerOptions: {
                container: document.getElementById('dplayer'), //播放器容器
                mutex: false, //  防止同时播放多个用户，在该用户开始播放时暂停其他用户
                theme: '#b7daff', // 风格颜色，例如播放条，音量条的颜色
                loop: false, // 是否自动循环
                lang: 'zh-cn', // 语言，'en', 'zh-cn', 'zh-tw'
                // screenshot: true, // 是否允许截图（按钮），点击可以自动将截图下载到本地
                hotkey: true, // 是否支持热键，调节音量，播放，暂停等
                preload: 'auto', // 自动预加载
                volume: 0.7, // 初始化音量
                playbackSpeed: [0.5, 0.75, 1, 1.25, 1.5, 2, 3], //可选的播放速度，可自定义
                video: {
                    url: 'https://cesium.com/public/SandcastleSampleData/big-buck-bunny_trailer.mp4', // 播放视频的路径
                    defaultQuality: 0 // 默认是高清
                    // pic: 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1606462956126&di=2d87964d4faf656af55d09d938640d97&imgtype=0&src=http%3A%2F%2Fattachments.gfan.com%2Fforum%2Fattachments2%2F201310%2F10%2F150326y7dzdd8d4kpjjdsd.jpg', // 视频封面图片
                    //  thumbnails: 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1606462956126&di=2d87964d4faf656af55d09d938640d97&imgtype=0&src=http%3A%2F%2Fattachments.gfan.com%2Fforum%2Fattachments2%2F201310%2F10%2F150326y7dzdd8d4kpjjdsd.jpg' // 进度条上的缩略图,需要通过dplayer-thumbnails来生成
                },
                subtitle: {
                    //字幕
                    url: '', //字幕网址
                    color: '#fff', //字幕颜色
                    fontSize: '20px',
                    bottom: '40px'
                },
                highlight: [
                    //进度条上的自定义时间标记
                    // 进度条时间点高亮
                    {
                        text: '10M',
                        time: 6
                    },
                    {
                        text: '20M',
                        time: 12
                    }
                ]
            }
        };
    },
    mounted() {
        this.updateInfo();
        this.initTitle();
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
        changeVideo() {
            this.isVideoOpen = !this.isVideoOpen;
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

.contentContainer {
    width: 65%;
    height: 100%;
    overflow: clip;
    display: flex;
    flex-direction: column;
    position: relative;
    margin: 0px auto;
}

.contentHeader {
    display: flex;
    -webkit-box-pack: justify;
    justify-content: space-between;
    top: 0px;
    left: 0px;
    -webkit-box-align: center;
    align-items: flex-end;
    width: 100%;
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
    font-size: 12px !important;
    font-weight: 530;
    color: #141414;
    line-height: 16px;
    outline: none;
    box-shadow: none;
    padding: 10px;
}
/*激活样式*/
::v-deep .el-radio-button__orig-radio:checked + .el-radio-button__inner {
    background: #4787fc;
    border: 1 !important;
    color: #faf5f5;
    font-size: 12px !important;
    font-weight: 500;
    line-height: 16px;
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

.aside-icon:hover > div > svg > path {
    fill: #0552f7 !important;
}

.aside-icon:hover > div > p {
    color: #0552f7 !important;
}

.tools-icon:hover > svg > path {
    fill: #0552f7 !important;
}

.tools-icon:hover {
    background-color: #abbce067;
}

.contentBody {
    height: 100%;
}

</style>