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
                <div class="aside-icon text-center mb-8 block px-4 cursor-pointer" @click="generateDocx">
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
            <div class="contentContainer w-full pt-6" v-if="resource_type !== 2">
                <div class="contentHeader w-full pb-2">
                    <div class="flex justify-between w-full px-8">
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
                            <el-tooltip effect="dark" content="搜索" placement="bottom" v-if="!isSearchContent">
                                <div
                                    class="tools-icon cursor-pointer w-10 h-10 rounded-lg flex items-center justify-center mr-3"
                                    @click="changeSearchStatus"
                                >
                                    <svg
                                        t="1719539453197"
                                        class="icon"
                                        viewBox="0 0 1024 1024"
                                        version="1.1"
                                        xmlns="http://www.w3.org/2000/svg"
                                        p-id="8356"
                                        width="26"
                                        height="26"
                                    >
                                        <path
                                            d="M944.951334 875.009893L777.542121 707.703008c48.605976-64.978515 77.564905-145.306286 77.564904-232.797042 0-215.401219-174.776856-389.973419-390.280404-389.973419C249.425402 85.034876 74.750874 259.607075 74.750874 475.008294S249.425402 864.981713 464.92895 864.981713c97.825922 0 186.953932-36.224243 255.411612-95.574698l165.055661 165.055661c8.18627 8.18627 18.930748 12.381733 29.777555 12.381733 10.846807 0 21.488958-4.195463 29.777556-12.381733 16.372539-16.474868 16.372539-43.080244 0-59.452783z m-480.022384-94.039772c-168.841811 0-306.268812-137.222344-306.268812-306.064155s137.427001-305.961827 306.268812-305.961827S771.197762 306.268812 771.197762 475.008294s-137.427001 305.961827-306.268812 305.961827z"
                                            p-id="8357"
                                        />
                                    </svg>
                                </div>
                            </el-tooltip>
                            <div class="search-top" v-else>
                                <div class="search-top-left" :class="{ focus: isSearchInputFocused }">
                                    <div class="search-input">
                                        <div>
                                            <svg
                                                height="14"
                                                viewBox="0 0 14 14"
                                                width="14"
                                                data-spm-anchor-id="5176.28158893.0.i62.22f86ac59e5h91"
                                            >
                                                <mask id="a" fill="#fff">
                                                    <path d="m0 0h16v16h-16z" fill="#fff" fill-rule="evenodd" />
                                                </mask>
                                                <path
                                                    d="m7 1.33333333c3.1296136 0 5.6666667 2.53705309 5.6666667 5.66666667 0 1.11726311-.3233397 2.15900619-.8815482 3.0367583l2.6695352 2.6703485-1.4142136 1.4142135-2.6099677-2.6106211c-.95224777.7253155-2.14106801 1.1559675-3.4304724 1.1559675-3.12961358 0-5.66666667-2.5370531-5.66666667-5.6666667 0-3.12961358 2.53705309-5.66666667 5.66666667-5.66666667zm0 2c-2.02504408 0-3.66666667 1.64162259-3.66666667 3.66666667s1.64162259 3.6666667 3.66666667 3.6666667 3.6666667-1.64162262 3.6666667-3.6666667-1.64162262-3.66666667-3.6666667-3.66666667z"
                                                    fill="#7d7d94"
                                                    fill-rule="evenodd"
                                                    mask="url(#a)"
                                                    transform="translate(-1 -1)"
                                                />
                                            </svg>
                                        </div>

                                        <input
                                            width="450"
                                            type="text"
                                            placeholder="输入内容回车进行搜索"
                                            class="search-input-field"
                                            v-model="searchInfo"
                                            @input="handleSearchInputChange"
                                            @focus="isSearchInputFocused = true"
                                            @blur="isSearchInputFocused = false"
                                            @keydown.enter="handleSearchEnterKey"
                                        />

                                        <div style="width: 14px"></div>

                                        <div class="clear-searchInfo" @click="clearSearchInfo" v-if="isShowClearSearchInfo">
                                            <svg
                                                t="1719557095120"
                                                class="icon"
                                                viewBox="0 0 1024 1024"
                                                version="1.1"
                                                xmlns="http://www.w3.org/2000/svg"
                                                p-id="13728"
                                                width="14"
                                                height="14"
                                            >
                                                <path
                                                    d="M512 1024c-285.257143 0-512-226.742857-512-512s226.742857-512 512-512 512 226.742857 512 512-226.742857 512-512 512z m-7.314286-585.142857l-146.285714-146.285714c-21.942857-21.942857-51.2-21.942857-73.142857 0-21.942857 21.942857-21.942857 51.2 0 73.142857l146.285714 146.285714-146.285714 146.285714c-21.942857 21.942857-21.942857 51.2 0 73.142857 21.942857 21.942857 51.2 21.942857 73.142857 0l146.285714-146.285714 146.285715 146.285714c21.942857 21.942857 51.2 21.942857 73.142857 0s21.942857-51.2 0-73.142857l-146.285715-146.285714 146.285715-146.285714c21.942857-21.942857 21.942857-51.2 0-73.142857s-51.2-21.942857-65.828572 0L504.685714 438.857143z"
                                                    fill="#B9B9B9"
                                                    p-id="13729"
                                                />
                                            </svg>
                                        </div>

                                        <div class="switchover" v-if="isShowSwitchButton">
                                            <span class="search-line"></span>
                                            <div class="switchover-button" @click="leftSwitchButton">
                                                <svg height="12" viewBox="0 0 8 12" width="8" class="inputArrow">
                                                    <mask id="a" fill="#fff">
                                                        <path d="m0 0h16v16h-16z" fill="#fff" fill-rule="evenodd" />
                                                    </mask>
                                                    <path
                                                        d="m11.8314211 3.91089519-1.5236962-1.45765507-4.93481912 4.94280601c-.37038971.36600473-.37394444.96297037-.0079397 1.33336008.00131535.00133111.00263466.00265831.00395792.00398156l.01473341.01473341 4.92406959 4.92407112 1.5236941-1.4672563-3.92681553-3.74462761-.00082918-.00084419c-.21621338-.22012815-.21584362-.57300593.00083061-.79268047z"
                                                        fill="#7d7d94"
                                                        mask="url(#a)"
                                                        transform="translate(-4.529672 -2.125432)"
                                                    />
                                                </svg>
                                            </div>
                                            <span class="switchover-inputNum">
                                                <span style="color: rgba(39, 38, 77, 0.65)">{{ searchedCurNum }}</span>
                                                <span>/</span>
                                                <span>{{ searchedTotalNum }}</span>
                                            </span>
                                            <div class="switchover-button" @click="rightSwitchButton">
                                                <svg
                                                    height="12"
                                                    viewBox="0 0 8 12"
                                                    width="8"
                                                    class="inputArrow"
                                                    style="transform: rotate(180deg)"
                                                >
                                                    <mask id="a" fill="#fff">
                                                        <path d="m0 0h16v16h-16z" fill="#fff" fill-rule="evenodd" />
                                                    </mask>
                                                    <path
                                                        d="m11.8314211 3.91089519-1.5236962-1.45765507-4.93481912 4.94280601c-.37038971.36600473-.37394444.96297037-.0079397 1.33336008.00131535.00133111.00263466.00265831.00395792.00398156l.01473341.01473341 4.92406959 4.92407112 1.5236941-1.4672563-3.92681553-3.74462761-.00082918-.00084419c-.21621338-.22012815-.21584362-.57300593.00083061-.79268047z"
                                                        fill="#7d7d94"
                                                        mask="url(#a)"
                                                        transform="translate(-4.529672 -2.125432)"
                                                    />
                                                </svg>
                                            </div>
                                        </div>

                                        <div class="search-error" v-if="isSearchInputError">
                                            <span class="search-line"></span>请输入内容
                                        </div>

                                        <div class="search-error" v-if="isSearchNone"><span class="search-line"></span>搜索无结果</div>

                                        <span class="cancel">
                                            <span class="line"></span>
                                            <span class="cancelText" @click="cancelSearchStatus">取消</span>
                                        </span>
                                    </div>
                                </div>
                            </div>

                            <el-popover placement="bottom-end" width="280" trigger="hover" visible-arrow="false">
                                <div class="p-2 select-none space-y-2">
                                    <div class="font-semibold text-black text-sm leading-6">一键替换</div>
                                    <div class="flex items-center">
                                        <div class="text-xs whitespace-nowrap mr-3">原文本</div>
                                        <el-input v-model="originText" placeholder="请输入原文本内容" append-slot="suffix">
                                            <template #suffix>
                                                <div class="flex items-center h-full">
                                                    <span class="font-bold text-yellow-500">{{ replaceNum }}</span>
                                                </div>
                                            </template>
                                        </el-input>
                                    </div>
                                    <div class="flex items-center">
                                        <div class="text-xs whitespace-nowrap mr-3">替换为</div>
                                        <el-input v-model="replaceText" placeholder="请输入替换内容"></el-input>
                                    </div>

                                    <div class="flex items-center">
                                        <div class="text-xs whitespace-nowrap mr-12"></div>
                                        <el-button
                                            :plain="true"
                                            :class="{
                                                'bg-blue-500 text-white': isReplaceAvailable,
                                                'bg-gray-500 text-white': !isReplaceAvailable
                                            }"
                                            @click="allReplace"
                                            class="custom-rounded w-48"
                                            >一键替换</el-button
                                        >
                                    </div>
                                </div>
                                <div
                                    class="tools-icon cursor-pointer w-10 h-10 rounded-lg flex items-center justify-center mr-3"
                                    slot="reference"
                                >
                                    <svg
                                        t="1719737321919"
                                        class="icon"
                                        viewBox="0 0 1024 1024"
                                        version="1.1"
                                        xmlns="http://www.w3.org/2000/svg"
                                        p-id="2326"
                                        width="26"
                                        height="26"
                                    >
                                        <path
                                            d="M461.72 472.43h-267a36 36 0 0 1-36-36v-267a36 36 0 0 1 36-36h267a36 36 0 0 1 36 36v267a36 36 0 0 1-36 36z m-231-72h195v-195h-195zM829.24 890.53h-267a36 36 0 0 1-36-36v-267a36 36 0 0 1 36-36h267a36 36 0 0 1 36 36v267a36 36 0 0 1-36 36z m-231-72h195v-195h-195zM848.8 480.18a36 36 0 0 1-36-36c0-124-100.89-224.91-224.91-224.91a36 36 0 0 1 0-72A296.9 296.9 0 0 1 884.8 444.18a36 36 0 0 1-36 36z"
                                            p-id="2327"
                                        />
                                        <path
                                            d="M848.8 480.18H718.35a36 36 0 0 1 0-72H848.8a36 36 0 0 1 0 72zM436.11 876.73A296.9 296.9 0 0 1 139.2 579.82a36 36 0 0 1 72 0c0 124 100.89 224.91 224.91 224.91a36 36 0 1 1 0 72z"
                                            p-id="2328"
                                        />
                                        <path d="M305.65 615.82H175.2a36 36 0 1 1 0-72h130.45a36 36 0 0 1 0 72z" p-id="2329" />
                                    </svg>
                                </div>
                            </el-popover>
                            <el-popover placement="bottom-end" width="280" trigger="hover" visible-arrow="false" v-if="aiType === 0">
                                <div class="p-2 select-none">
                                    <div class="font-semibold text-black text-sm leading-6">翻译</div>
                                    <div class="flex items-center">
                                        <el-radio-group v-model="showType">
                                            <el-radio-button label="0">源语言显示</el-radio-button>
                                            <el-radio-button label="1">纯译文显示</el-radio-button>
                                            <el-radio-button label="2">双语显示</el-radio-button>
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
                            <el-tooltip
                                class="item"
                                effect="dark"
                                content="展开视频"
                                placement="bottom"
                                v-if="!isVideoOpen && resource_type === 1"
                            >
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

                            <el-tooltip
                                class="item"
                                effect="dark"
                                content="收起视频"
                                placement="bottom"
                                v-else-if="isVideoOpen && resource_type === 1"
                            >
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
                                    <div class="aiRadio mt-2">
                                        <div class="whitespace-nowrap">显示内容：</div>
                                        <el-radio-group v-model="aiType">
                                            <el-radio :label="0">原文显示</el-radio>
                                            <el-radio :label="1" class="mb-1">改写结果</el-radio>
                                            <el-radio :label="2">原文和改写结果</el-radio>
                                        </el-radio-group>
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
                    <div class="audio-player" v-if="resource_type === 0">
                        <audio controls class="audio-control" ref="audio" v-if="showVoiceUrl">
                            <source :src="voice_url" />
                        </audio>
                    </div>

                    <div class="w-full h-96 px-8" v-else-if="isVideoOpen && resource_type === 1">
                        <d-player id="dplayer" class="h-full" :options="playerOptions"></d-player>
                    </div>

                    <div class="scrollContainer overflow-auto py-2 px-8" :style="computedResourceTypeStyle">
                        <div class="AIDIV mb-2">
                            <div class="zhinengsulan mb-3 select-none">
                                <img src="../../assets/img/zhinengsulan.png" alt class="h-12 pointer-events-none" />
                            </div>
                            <div class="keyword">
                                <div class="font-semibold text-sm text-black leading-6">关键词</div>
                                <div class="flex select-none">
                                    <div
                                        class="keywordleft overflow-hidden w-full mb-4"
                                        :style="{ height: openkeyword ? sourceDivHeight + 'px' : '40px' }"
                                        style="height 0.3s ease 0s;"
                                    >
                                        <div id="keywordh" class="flex flex-wrap">
                                            <div
                                                v-for="tag in recent_tag"
                                                :key="tag.id"
                                                class="py-0.5 px-2 bg-blue-100 rounded font-normal text-sm text-blue-400 mr-2 mt-3 leading-5 text-center cursor-pointer block"
                                            >
                                                {{ tag }}
                                            </div>
                                        </div>
                                    </div>
                                    <div class="keywordright relative top-4 w-16 text-xs leading-5 text-right cursor-pointer text-blue-400">
                                        <div class @click="changeKeyword">{{ keywordopentitle }}</div>
                                    </div>
                                </div>
                            </div>
                            <div class="abstract mb-2">
                                <div class="font-semibold text-sm text-black leading-6">全文概要</div>
                                <div class="mt-1">
                                    <div class="relative">
                                        <div class="relative flex mb-2">
                                            <div
                                                :class="{
                                                    breakabstractword: !openAbstract,
                                                    openabstractword: openAbstract
                                                }"
                                            >
                                                {{ text_summary }}
                                            </div>
                                        </div>
                                        <div class="text-xs text-blue-400 text-right">
                                            <span class="cursor-pointer" @click="changeAbstract">{{ abstractopentitle }}</span>
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <div class="aisum">
                                <div class="mb-2" style="width: 300px">
                                    <el-menu
                                        :default-active="aiSumActiveIndex"
                                        class="el-menu-demo"
                                        mode="horizontal"
                                        background-color="rgb(0 0 0 / 0%)"
                                        active-text-color="rgb(0 0 0)"
                                        @select="handleSelectAisum"
                                    >
                                        <el-menu-item index="1">章节速览</el-menu-item>
                                        <el-menu-item index="2">要点回顾</el-menu-item>
                                    </el-menu>
                                </div>
                                <div class="aisumContent mb-2" v-if="aiSumActiveIndex === '1'">
                                    <div class="relative">
                                        <div class="relative flex" v-for="(scanning_item, index) in scanning_items" :key="index">
                                            <div class="leftAisumContent mr-4 relative">
                                                <div class="h-10 flex items-center relative cursor-default select-none">
                                                    <div class="w-11 text-xs text-black select-none">
                                                        {{ scanning_item['time'] }}
                                                    </div>
                                                    <div>
                                                        <svg
                                                            t="1716204788686"
                                                            class="icon"
                                                            viewBox="0 0 1024 1024"
                                                            version="1.1"
                                                            xmlns="http://www.w3.org/2000/svg"
                                                            p-id="2618"
                                                            width="8"
                                                            height="8"
                                                        >
                                                            <path
                                                                d="M512 520.064m-414.464 0a3.238 3.238 0 1 0 828.928 0 3.238 3.238 0 1 0-828.928 0Z"
                                                                fill="#1296db"
                                                                p-id="2619"
                                                                data-spm-anchor-id="a313x.search_index.0.i0.44783a81zqBcfd"
                                                                class="selected"
                                                            />
                                                        </svg>
                                                    </div>
                                                </div>
                                            </div>
                                            <div class="rightAisumContent block w-full pt-0.5 cursor-pointer">
                                                <div class="min-h-10 mb-2 w-full bg-blue-100 rounded py-2 px-4 relative flex border">
                                                    <div style="visibilit: visible">
                                                        <div class="aisumtext mb-1" :class="{ breakaisum: !openEpSumScanning }">
                                                            {{ scanning_item['title'] }}
                                                        </div>
                                                        <div class="text-xs text-gray-400" v-if="openEpSumScanning">
                                                            {{ scanning_item['desc'] }}
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>
                                        <div class="text-xs text-blue-400 cursor-pointer ml-16" @click="changeEpScanning">
                                            {{ epopentitleScanning }}
                                        </div>
                                    </div>
                                </div>
                                <div class="aisumContent mb-2" v-if="aiSumActiveIndex === '2'">
                                    <div class="relative">
                                        <div class="relative flex" v-for="(review_item, index) in review_items" :key="index">
                                            <div class="leftAisumContent mr-4 relative">
                                                <div class="h-10 flex items-center relative cursor-default select-none">
                                                    <div class="w-11 text-sm font-bold text-blue-400 select-none">要点</div>
                                                    <div>
                                                        <svg
                                                            t="1716204788686"
                                                            class="icon"
                                                            viewBox="0 0 1024 1024"
                                                            version="1.1"
                                                            xmlns="http://www.w3.org/2000/svg"
                                                            p-id="2618"
                                                            width="8"
                                                            height="8"
                                                        >
                                                            <path
                                                                d="M512 520.064m-414.464 0a3.238 3.238 0 1 0 828.928 0 3.238 3.238 0 1 0-828.928 0Z"
                                                                fill="#1296db"
                                                                p-id="2619"
                                                                data-spm-anchor-id="a313x.search_index.0.i0.44783a81zqBcfd"
                                                                class="selected"
                                                            />
                                                        </svg>
                                                    </div>
                                                </div>
                                            </div>
                                            <div class="rightAisumContent block w-full pt-0.5 cursor-pointer">
                                                <div class="min-h-10 mb-2 w-full bg-blue-100 rounded py-2 px-4 relative flex border">
                                                    <div style="visibilit: visible">
                                                        <div class="aisumtext mb-1" :class="{ breakaisum: !openEpSumReview }">
                                                            {{ review_item['ques'] }}
                                                        </div>
                                                        <div class="text-xs text-gray-400" v-if="openEpSumReview">
                                                            {{ review_item['ans'] }}
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>
                                        <div class="text-xs text-blue-400 cursor-pointer ml-16" @click="changeEpReview">
                                            {{ epopentitleReview }}
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <div class="heng mt-4 mb-4">
                                <div class="flex items-center w-full" style="justify-content: space-between; -webkit-box-align: center">
                                    <div class="sc-djXQSe ccTlNE" style="display: block"></div>
                                    <div class="text-sm text-gray-300">以上内容由蓝心AI大模型生成</div>
                                    <div class="sc-jDba-dz jaCZOE" style="display: block"></div>
                                </div>
                            </div>
                            <div class="sourcetitle flex items-center">
                                <div class="mr-2">
                                    <svg
                                        t="1716206486538"
                                        class="icon"
                                        viewBox="0 0 1024 1024"
                                        version="1.1"
                                        xmlns="http://www.w3.org/2000/svg"
                                        p-id="3656"
                                        width="20"
                                        height="20"
                                    >
                                        <path
                                            d="M862.72 837.632c0 28.16-21.504 50.176-50.176 50.176h-10.24c6.656-13.312 10.24-26.624 10.24-41.984V178.176c0-14.848-3.584-30.208-10.24-41.984h10.24c28.16 0 50.176 21.504 50.176 50.176v651.264zM293.888 318.464c-17.92 0-33.28-11.776-33.28-25.088 0-13.312 15.36-25.088 33.28-25.088h355.328c17.92 0 33.28 11.776 33.28 25.088 0 13.312-15.36 25.088-33.28 25.088H293.888z m-8.704 168.448c-13.312 0-25.088-11.776-25.088-25.088 0-13.312 11.776-25.088 25.088-25.088h267.264c13.312 0 25.088 11.776 25.088 25.088 0 13.312-11.776 25.088-25.088 25.088H285.184z m0 152.064c-13.312 0-25.088-11.776-25.088-25.088s11.776-25.088 25.088-25.088h150.528c13.312 0 25.088 11.776 25.088 25.088s-11.776 25.088-25.088 25.088H285.184z m-81.92-552.448c-50.176 0-91.648 41.984-91.648 91.648v666.112c0 51.712 41.984 91.648 91.648 91.648h609.28c55.296 0 100.352-45.056 100.352-100.352V186.368c0-55.296-45.056-100.352-100.352-100.352h-609.28z"
                                            fill="#1296db"
                                            p-id="3657"
                                        />
                                    </svg>
                                </div>
                                <div class="text-lg font-semibold">原文</div>
                            </div>
                        </div>
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
                                    <div v-if="(aiType === 0 && (showType === '0' || showType === '2')) || aiType === 2" class="jItVgd">
                                        <div class="kCofWf">
                                            <div
                                                class="kFJVLr w-full"
                                                contenteditable
                                                @focusin="inDiv(index, $event)"
                                                @focusout="outDiv(index, $event)"
                                                @input="onContentChange(index, $event)"
                                                @compositionstart="onCompositionStart"
                                                @compositionend="onCompositionEnd(index, $event)"
                                                :ref="'editableDiv' + index"
                                            >
                                                <span
                                                    v-for="(item_content, content_index) in recording_items[index].content"
                                                    :key="content_index"
                                                    :id="generateSpanID(index, content_index)"
                                                    :class="computeSearchSpanClass('content', index, content_index)"
                                                    >{{ item_content }}</span
                                                >
                                            </div>
                                        </div>
                                    </div>

                                    <div v-if="aiType === 0 && showType === '2'" class="solid-line"></div>

                                    <div class="w-full text-left ml-12 mt-4" v-if="aiType === 1">
                                        <div class="text-xs text-blue-300">AI改写结果：</div>
                                    </div>
                                    <div class="w-full text-left ml-12" v-if="aiType === 2">
                                        <div class="text-xs text-blue-300">AI改写结果：</div>
                                    </div>

                                    <div
                                        v-if="(aiType === 0 && (showType === '1' || showType === '2')) || aiType === 1 || aiType === 2"
                                        class="jItVgd"
                                    >
                                        <div class="kCofWf">
                                            <div class="kFJVLr w-full">
                                                <span
                                                    v-for="(trans_item_content, trans_content_index) in recording_items[index].transContent"
                                                    :key="trans_content_index"
                                                    :id="generateTransSpanID(index, trans_content_index)"
                                                    :class="computeSearchSpanClass('trans', index, trans_content_index)"
                                                    >{{ trans_item_content }}</span
                                                >
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div class="h-1 w-full"></div>
                    </div>
                </div>
            </div>

            <iframe
                :src="pdf_url"
                style="border: none; width: 50%; height: 100%"
                class="p-2"
                v-if="pdf_url !== null && resource_type === 2"
            ></iframe>

            <div class="h-full w-1/2 bg-white px-8 pt-3 pb-16">
                <el-tabs v-model="activePDFName" @tab-click="handlePDFClick" class v-if="resource_type === 2">
                    <el-tab-pane name="first">
                        <template #label>
                            <span class="text-base">智能速览</span>
                        </template>
                    </el-tab-pane>
                    <el-tab-pane name="second">
                        <template #label>
                            <span class="text-base">笔记</span>
                        </template>
                    </el-tab-pane>
                </el-tabs>

                <div class="introduction" style="overflow-y: auto" v-if="activePDFName === 'first' && resource_type === 2">
                    <div class="introduction-content">
                        <div>
                            <span class="taskTitle">{{ pdfFileTitle }}</span>
                            <div class="title-wrapper">
                                <span class="title">全文概述</span>
                            </div>
                            <div class="card">
                                <div class="markdown-body">
                                    <p>{{ pdfOverview }}</p>
                                </div>
                            </div>
                        </div>

                        <div>
                            <div class="title-wrapper mt-4">
                                <span class="title">关键要点</span>
                            </div>
                            <div class="card">
                                <div class="markdown-body ml-8">
                                    <ol>
                                        <li v-for="(pdfKeyPoint, index) in pdfKeyPoints" :key="index">
                                            {{ pdfKeyPoint }}
                                        </li>
                                    </ol>
                                </div>
                            </div>
                        </div>

                        <div>
                            <div class="title-wrapper mt-4">
                                <span class="title">文档速读</span>
                            </div>
                            <div class="card">
                                <div class="markdown-body ml-4">
                                    <div class="aisumContent mb-2">
                                        <div class="relative">
                                            <div class="relative flex" v-for="(read_item, index) in pdfQuickReading" :key="index">
                                                <div class="leftAisumContent mr-4 relative">
                                                    <div class="h-10 flex items-center relative cursor-default select-none">
                                                        <div class="w-11 text-xs text-black select-none">
                                                            {{ read_item['page'] }}
                                                        </div>
                                                        <div>
                                                            <svg
                                                                t="1716204788686"
                                                                class="icon"
                                                                viewBox="0 0 1024 1024"
                                                                version="1.1"
                                                                xmlns="http://www.w3.org/2000/svg"
                                                                p-id="2618"
                                                                width="8"
                                                                height="8"
                                                            >
                                                                <path
                                                                    d="M512 520.064m-414.464 0a3.238 3.238 0 1 0 828.928 0 3.238 3.238 0 1 0-828.928 0Z"
                                                                    fill="#1296db"
                                                                    p-id="2619"
                                                                    data-spm-anchor-id="a313x.search_index.0.i0.44783a81zqBcfd"
                                                                    class="selected"
                                                                />
                                                            </svg>
                                                        </div>
                                                    </div>
                                                </div>
                                                <div class="rightAisumContent block w-full pt-0.5 cursor-pointer">
                                                    <div class="min-h-10 mb-2 w-full bg-blue-100 rounded py-2 px-4 relative flex border">
                                                        <div style="visibilit: visible">
                                                            <div class="aisumtext breakaisum mb-1">
                                                                {{ read_item['title'] }}
                                                            </div>
                                                            <div class="text-xs text-gray-400">
                                                                {{ read_item['content'] }}
                                                            </div>
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="h-full editorPDFViewport -mt-6 mx-3" v-else-if="activePDFName === 'second' && resource_type === 2">
                    <quill-editor
                        class="h-full"
                        v-model="content"
                        :options="editorOption"
                        @blur="onEditorBlur($event)"
                        @focus="onEditorFocus($event)"
                        @change="onEditorChange($event)"
                    ></quill-editor>
                </div>
                <div class="h-full editorViewport" v-else-if="resource_type !== 2">
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
import flvjs from 'flv.js';
import { quillEditor } from 'vue-quill-editor';
import { voiceTrans } from '../../api/voiceTrans';
import { saveVoice } from '../../api/voiceTrans';
import { getRecord, getPDFRecord, getAiRewriting } from '../../api/record';
import { userStore } from '../../store/store';
import 'quill/dist/quill.core.css';
import 'quill/dist/quill.snow.css';
import 'quill/dist/quill.bubble.css';
import VueDPlayer from 'vue-dplayer';
import 'vue-dplayer/dist/vue-dplayer.css';
import { encodeWithSalt, decodeWithSalt, Salt } from '../../utils/util';

import JSZip from 'jszip';
import Docxtemplater from 'docxtemplater';
import { saveAs } from 'file-saver';
import PizZipUtils from 'pizzip/utils/index.js';
import PizZip from 'pizzip';

const userstore = userStore();
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
            recording_items: [],
            copy_recording_items: null,
            currentTime: 0, // 当前时间，单位：秒
            timer: null, // 定时器
            inputWidth: 204,
            recordInfo: '',
            savedTime: '',
            content: null,
            sourceLanguage: '',
            destLanguage: '',
            resource_type: 0,

            openkeyword: false,
            keywordopentitle: '展开全部',
            openAbstract: false,
            abstractopentitle: '展开全部',
            aiSumActiveIndex: '1',
            openEpSumScanning: false,
            openEpSumReview: false,
            epopentitleScanning: '查看章节摘要',
            epopentitleReview: '查看章节摘要',
            recent_tag: [],
            text_summary: '',

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
            showType: '0', // 0源语言显示, 1纯译文显示, 2双语显示

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
                    url: '', // 播放视频的路径
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
            },

            originText: '', // 原文本内容
            replaceText: '', // 替换文本

            isSearchInputFocused: false, // 搜索框聚焦状态产生的变化, 不需要关心
            isSearchContent: false, // 显示搜索框
            searchInfo: '', // 搜索框里面的内容
            isSearchInputError: false, // 搜索框为空时起作用
            isSearchNone: false, // 搜索结果为None时
            isShowClearSearchInfo: false, // 当搜索框不为空时，显示清除图案
            isShowSwitchButton: false, // 搜索到有内容, 就有可切换状态
            searchedCurNum: 1,
            searchedTotalNum: 1,
            dictSpanClass: {}, // {id: class}
            dictIndexID: {}, // {index: id}
            replaceNum: 0,

            activePDFName: 'first',

            scanning_items: [],
            review_items: [],

            pdfFileTitle: '',
            pdfOverview: '',
            pdfKeyPoints: [],
            pdfQuickReading: [],

            showVoiceUrl: false,
            voice_url: '',
            pdf_url: '',

            recordId: null,

            guide_template: 'http://172.30.230.46/template/guide_template.docx',
            origin_text: [],
            origin_template: 'http://172.30.230.46/template/origin_template.docx',
            AI_rewrite_template: '',
            pdf_template: 'http://172.30.230.46/template/pdf_template.docx'
        };
    },
    mounted() {
        this.tempRoute = Object.assign({}, this.$route);
        this.recordId = decodeWithSalt(this.$route.query.id, Salt);

        this.initTitle();
        getRecord(this.recordId).then((response) => {
            let ret_data = response;
            if (ret_data['datatype'] > 1) {
                this.resource_type = 2;
                this.pdf_url = ret_data['path'];
                console.log(ret_data['path']);
                this.pdfFileTitle = ret_data['title'];
                this.pdfOverview = ret_data['summary'];
                this.pdfKeyPoints = [];
                for (let i = 0; i < ret_data['keyword'].length; ++i) {
                    this.pdfKeyPoints.push(ret_data['keyword'][i]['desc']);
                }

                for (let i = 0; i < ret_data['scanning'].length; ++i) {
                    let pageStr = ret_data['scanning'][i]['index'];
                    let parts = pageStr.split('_');
                    let number = '第' + parts[1] + '页';
                    this.pdfQuickReading.push({
                        page: number,
                        title: ret_data['scanning'][i]['title'],
                        content: ret_data['scanning'][i]['desc']
                    });
                }
            } else {
                let ret_inner_text = ret_data['inner_text'];
                this.resource_type = ret_data['datatype'];

                if (this.resource_type === 0 || this.resource_type === 1) {
                    this.content = ret_inner_text['content'] === '' ? null : ret_inner_text['content'];
                    this.recordInfo = ret_inner_text['recordInfo'] === '' ? this.getYMDHMSTime : ret_inner_text['recordInfo'];
                    this.savedTime = ret_inner_text['savedTime'] === '' ? this.getCurSavedTime : ret_inner_text['savedTime'];
                    this.sourceLanguage = ret_inner_text['source_language'];
                    this.destLanguage = ret_inner_text['dest_language'];
                    let trans_text = ret_data['trans_text'];
                    this.origin_text = trans_text;
                    for (let i = 0; i < trans_text.length; ++i) {
                        this.addRecordingItem(trans_text[i]['timestamp'], trans_text[i]['content']);
                    }

                    for (let i = 0; i < ret_data['keyword'].length; ++i) {
                        this.recent_tag.push(ret_data['keyword'][i]['key']);
                    }

                    this.text_summary = ret_data['summary'];

                    this.review_items = ret_data['review'];
                    this.scanning_items = ret_data['scanning'];

                    if (this.resource_type === 0) {
                        this.voice_url = ret_data['path'];
                        this.showVoiceUrl = true;
                    } else if (this.resource_type === 1) {
                        const filename = ret_data['path'].split('/').pop();
                        // 获取扩展名
                        const extension = filename.split('.').pop().toLowerCase();
                        if (extension === 'flv') {
                            this.playerOptions['video'] = {
                                url: ret_data['path'],
                                type: 'customFlv',
                                customType: {
                                    customFlv: function (video, player) {
                                        const flvPlayer = flvjs.createPlayer({
                                            type: 'flv',
                                            url: video.src
                                        });
                                        flvPlayer.attachMediaElement(video);
                                        flvPlayer.load();
                                    }
                                }
                            };
                        } else if (extension === 'mp4') {
                            this.playerOptions['video']['url'] = ret_data['path'];
                        }
                    }
                }
            }
        });

        // getRecord("163").then((response) => {
        //   let ret_data = response;
        //   let ret_inner_text = ret_data["inner_text"];
        //   this.resource_type = ret_data["datatype"] === 3 ? 2 : ret_data["datatype"];

        //   if (this.resource_type === 0 || this.resource_type === 1) {
        //     this.content =
        //       ret_inner_text["content"] === "" ? null : ret_inner_text["content"];
        //     this.recordInfo =
        //       ret_inner_text["recordInfo"] === ""
        //         ? this.getYMDHMSTime
        //         : ret_inner_text["recordInfo"];
        //     this.savedTime =
        //       ret_inner_text["savedTime"] === ""
        //         ? this.getCurSavedTime
        //         : ret_inner_text["savedTime"];
        //     this.sourceLanguage = ret_inner_text["source_language"];
        //     this.destLanguage = ret_inner_text["dest_language"];
        //     let trans_text = ret_data["trans_text"];
        //     this.origin_text = trans_text;
        //     for (let i = 0; i < trans_text.length; ++i) {
        //       this.addRecordingItem(trans_text[i]["timestamp"], trans_text[i]["content"]);
        //     }

        //     for (let i = 0; i < ret_data["keyword"].length; ++i) {
        //       this.recent_tag.push(ret_data["keyword"][i]["key"]);
        //     }

        //     this.text_summary = ret_data["summary"];

        //     this.review_items = ret_data["review"];
        //     this.scanning_items = ret_data["scanning"];

        //     if (this.resource_type === 0) {
        //       this.voice_url =
        //         "http://172.30.230.46/voice/2/0916e14b0d2f684c6fedc18b0e040b38/%E5%B0%8F%E5%93%81%EF%BC%9A%E5%A4%A9%E7%8E%8B%E5%B7%A8%E6%98%9F-%E5%96%9C%E5%89%A7%E5%90%AC%E6%88%91%E7%9A%84.mp3";
        //       this.showVoiceUrl = true;
        //     } else if (this.resource_type === 1) {
        //       this.playerOptions["video"]["url"] =
        //         "https://cesium.com/public/SandcastleSampleData/big-buck-bunny_trailer.mp4";
        //     }
        //   }
        // });
    },
    computed: {
        getSaveFileName() {
            const now = new Date();
            const year = now.getFullYear();
            const month = String(now.getMonth() + 1).padStart(2, '0');
            const date = String(now.getDate()).padStart(2, '0');
            const hours = String(now.getHours()).padStart(2, '0');
            const minutes = String(now.getMinutes()).padStart(2, '0');
            const seconds = String(now.getSeconds()).padStart(2, '0');
            const milliseconds = String(now.getMilliseconds()).padStart(3, '0');
            return `${year}-${month}-${date}-${hours}-${minutes}-${seconds}-${milliseconds}`;
        },
        getYMDHMSTime() {
            const now = new Date();
            const year = now.getFullYear();
            const month = String(now.getMonth() + 1).padStart(2, '0');
            const date = String(now.getDate()).padStart(2, '0');
            const hours = String(now.getHours()).padStart(2, '0');
            const minutes = String(now.getMinutes()).padStart(2, '0');
            return `${year}-${month}-${date} ${hours}:${minutes} 记录`;
        },
        getExportTime() {
            const now = new Date();
            const year = now.getFullYear();
            const month = String(now.getMonth() + 1).padStart(2, '0');
            const date = String(now.getDate()).padStart(2, '0');
            const hours = String(now.getHours()).padStart(2, '0');
            const minutes = String(now.getMinutes()).padStart(2, '0');
            const seconds = String(now.getSeconds()).padStart(2, '0');
            const milliseconds = String(now.getMilliseconds()).padStart(3, '0');
            return `${year} 年 ${month} 月 ${date} 日 ${hours}:${minutes}`;
        },
        getCurSavedTime() {
            const hms = this.getHMSTime();
            return `已保存于 ${hms}`;
        },
        editor() {
            return this.$refs.myQuillEditor.quill;
        },
        isReplaceAvailable() {
            return false;
        },
        computedResourceTypeStyle() {
            if (this.resource_type === 0) {
                return {
                    height: 'calc(100% - 100px)'
                };
            } else if (this.resource_type === 1) {
                return {
                    height: this.isVideoOpen ? 'calc(100% - 400px)' : '100%'
                };
            } else {
            }
        }
    },
    watch: {
        originText(newValue) {
            this.replaceNum = 0;
            this.cancelSearchStatus();
            this.searchInfo = newValue;
            if (this.calcSearchContentInfo()) {
                this.replaceNum = this.searchedTotalNum;
            } else {
                this.replaceNum = 0;
            }
        },
        async showType(newValue) {
            if (newValue === '0') {
                console.log('newValue: ', newValue);
            } else if (newValue === '1' || newValue === '2') {
                console.log('newValue: ', newValue);
                for (let i = 0; i < this.recording_items.length; i++) {
                    let i_content = this.concatContent(this.recording_items[i].content);
                    voiceTrans(i_content, this.sourceLanguage, this.destLanguage).then((response) => {
                        this.recording_items[i].transContent = [response.data];
                    });
                    await this.customSleep(50);
                    console.log('i_content: ', i_content);
                }
            }
        },

        async aiType(newValue) {
            if (newValue === 0 && this.copy_recording_items !== null) {
                this.recording_items = this.copy_recording_items;
            } else if (newValue === 1 || newValue === 2) {
                this.copy_recording_items = this.recording_items;
                this.recording_items = this.copy_recording_items;
                var request = [];
                for (let i = 0; i < this.recording_items.length; i++) {
                    request.push(this.recording_items[i].content[0]);
                }
                getAiRewriting(request).then((response) => {
                    if (response.code === 200) {
                        
                        var res = JSON.parse(response.data);
                        console.log(res)
                        for (let i = 0; i < this.recording_items.length; i++) {
                            this.recording_items[i].transContent = [JSON.parse(res[i]).content];
                        }
                        this.$message({
                            message: 'AI改写成功',
                            type: 'success',
                            offset: 500
                        });
                    }
                    // this.recording_items[i].transContent = [response.data];
                });

                // console.log('newValue: ', newValue);
                // for (let i = 0; i < this.recording_items.length; i++) {
                //     let i_content = this.concatContent(this.recording_items[i].content);
                //     voiceTrans(i_content, this.sourceLanguage, this.destLanguage).then(response => {
                //         this.recording_items[i].transContent = [response.data];
                //     });
                //     await this.customSleep(50);
                //     console.log('i_content: ', i_content);
                // }
            }
        }
    },
    methods: {
        handleSelectAisum(index) {
            this.aiSumActiveIndex = index;
        },
        loadFile(url, callback) {
            PizZipUtils.getBinaryContent(url, callback);
        },
        generateDocx() {
            if (this.resource_type === 0 || this.resource_type === 1) {
                this.$confirm('导出智能速览及原文？').then(async (_) => {
                    // 导出摘要
                    let keys_content = '';
                    for (let i = 0; i < this.recent_tag.length; ++i) {
                        keys_content += this.recent_tag[i] + ' ';
                    }

                    let chapters_content = [];
                    for (let i = 0; i < this.scanning_items.length; ++i) {
                        chapters_content.push({
                            timestamp: this.scanning_items[i]['time'],
                            title: this.scanning_items[i]['title'],
                            content: this.scanning_items[i]['desc']
                        });
                    }

                    let export_guide_data = {
                        recordInfo: this.recordInfo,
                        exportTime: this.getExportTime,
                        keyPoints: keys_content,
                        summary: this.text_summary,
                        chapters: chapters_content,
                        keyInfo: this.review_items
                    };

                    this.loadFile(this.guide_template, (error, content) => {
                        if (error) {
                            throw error;
                        }

                        // 创建一个PizZip实例
                        const zip = new PizZip(content);

                        const doc = new Docxtemplater(zip, {
                            paragraphLoop: true,
                            linebreaks: true
                        });
                        // 设置数据
                        doc.setData(export_guide_data);

                        try {
                            // 渲染文档
                            doc.render();
                        } catch (error) {
                            const e = {
                                message: error.message,
                                name: error.name,
                                stack: error.stack,
                                properties: error.properties
                            };
                            console.log(JSON.stringify({ error: e }));
                            throw error;
                        }

                        const out = doc.getZip().generate({
                            type: 'blob',
                            mimeType: 'application/vnd.openxmlformats-officedocument.wordprocessingml.document'
                        });

                        // 下载文件
                        saveAs(out, this.recordInfo + '_导读.docx');
                    });

                    let export_origin_data = {
                        recordInfo: this.recordInfo,
                        exportTime: this.getExportTime,
                        originText: this.origin_text
                    };
                    // 导出原文
                    this.loadFile(this.origin_template, (error, content) => {
                        if (error) {
                            throw error;
                        }

                        // 创建一个PizZip实例
                        const zip = new PizZip(content);

                        const doc = new Docxtemplater(zip, {
                            paragraphLoop: true,
                            linebreaks: true
                        });
                        // 设置数据
                        doc.setData(export_origin_data);

                        try {
                            // 渲染文档
                            doc.render();
                        } catch (error) {
                            const e = {
                                message: error.message,
                                name: error.name,
                                stack: error.stack,
                                properties: error.properties
                            };
                            console.log(JSON.stringify({ error: e }));
                            throw error;
                        }

                        const out = doc.getZip().generate({
                            type: 'blob',
                            mimeType: 'application/vnd.openxmlformats-officedocument.wordprocessingml.document'
                        });

                        // 下载文件
                        saveAs(out, this.recordInfo + '_原文.docx');
                    });
                });
            } else if (this.resource_type === 2) {
                this.$confirm('导出文档智能速览？').then(async (_) => {
                    // 导出摘要
                    let keys_content = [];
                    for (let i = 0; i < this.pdfKeyPoints.length; ++i) {
                        keys_content.push({
                            index: i,
                            desc: this.pdfKeyPoints[i]
                        });
                    }

                    let chapters_content = [];
                    for (let i = 0; i < this.pdfQuickReading.length; ++i) {
                        chapters_content.push({
                            index: this.pdfQuickReading[i]['page'],
                            title: this.pdfQuickReading[i]['title'],
                            desc: this.pdfQuickReading[i]['content']
                        });
                    }

                    let export_data = {
                        title: this.pdfFileTitle,
                        exportTime: this.getExportTime,
                        summary: this.pdfOverview,
                        keyword: keys_content,
                        scanning: chapters_content
                    };

                    this.loadFile(this.pdf_template, (error, content) => {
                        if (error) {
                            throw error;
                        }

                        // 创建一个PizZip实例
                        const zip = new PizZip(content);

                        const doc = new Docxtemplater(zip, {
                            paragraphLoop: true,
                            linebreaks: true
                        });
                        // 设置数据
                        doc.setData(export_data);

                        try {
                            // 渲染文档
                            doc.render();
                        } catch (error) {
                            const e = {
                                message: error.message,
                                name: error.name,
                                stack: error.stack,
                                properties: error.properties
                            };
                            console.log(JSON.stringify({ error: e }));
                            throw error;
                        }

                        const out = doc.getZip().generate({
                            type: 'blob',
                            mimeType: 'application/vnd.openxmlformats-officedocument.wordprocessingml.document'
                        });

                        // 下载文件
                        saveAs(out, this.pdfFileTitle + '_智能速览.docx');
                    });
                });
            }
        },
        handlePDFClick(tab, event) {
            console.log(tab, event);
        },
        getHMSTime() {
            const now = new Date();
            const hours = String(now.getHours()).padStart(2, '0');
            const minutes = String(now.getMinutes()).padStart(2, '0');
            const seconds = String(now.getSeconds()).padStart(2, '0');
            return `${hours}:${minutes}:${seconds}`;
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
        onCompositionStart() {
            this.isComposing = true;
        },
        onCompositionEnd(index, event) {
            this.isComposing = false;
            this.onContentChange(index, event); // 输入法完成时触发内容变化处理
        },
        generateTransSpanID(index, trans_content_index) {
            return `trans-${index}-${trans_content_index}`;
        },
        generateSpanID(index, content_index) {
            return `content-${index}-${content_index}`;
        },
        scrollToSpanID(span_id) {
            const element = document.getElementById(span_id);
            if (element) {
                element.scrollIntoView({ behavior: 'smooth' });
            }
        },
        computeSearchSpanClass(prefix, one_index, two_index) {
            let span_id = `${prefix}-${one_index}-${two_index}`;
            if (this.dictSpanClass.hasOwnProperty(span_id)) {
                return this.dictSpanClass[span_id];
            } else {
                return '';
            }
        },
        onContentChange(index, event) {
            if (this.isComposing) {
                return; // 如果正在输入法输入，则不处理
            }
            // this.recording_items[index].content = [event.target.innerText];
            const context = this.$refs['editableDiv' + index][0];
            const caretPosition = this.getCaretPosition(context);

            // Update the model
            const contentArray = Array.from(context.childNodes).map((node) => node.textContent);
            this.recording_items[index].content = contentArray;

            this.$nextTick(() => {
                this.setCaretPosition(context, caretPosition);
                this.scrollToBottom();
            });

            if (this.searchInfo !== '') {
                if (this.calcSearchContentInfo()) {
                    this.isShowClearSearchInfo = true;
                    this.isShowSwitchButton = true;
                    this.isSearchNone = false;
                    this.isSearchInputError = false;
                } else {
                    this.scrollToSpanID(this.dictIndexID[this.searchedCurNum]);
                    this.isShowSwitchButton = false;
                    this.isSearchNone = true;
                }
            }
        },
        initTitle() {
            document.getElementsByClassName('ql-editor')[0].dataset.placeholder = '请在这里记录您的想法';
            for (let item of titleConfig) {
                let tip = document.querySelector('.quill-editor ' + item.Choice);
                if (!tip) continue;
                tip.setAttribute('title', item.title);
            }
        },
        concatContent(curContent) {
            let result_content = '';
            for (let i = 0; i < curContent.length; i++) {
                result_content += curContent[i];
            }
            return result_content;
        },
        addRecordingItem(item_timestamp, item_content) {
            // 添加一个新的item到数组中
            this.recording_items.push({
                time_stamp: item_timestamp,
                content: [item_content],
                transContent: ['']
            });
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
        },
        changeKeyword() {
            this.sourceDivHeight = document.getElementById('keywordh').offsetHeight;
            this.openkeyword = !this.openkeyword;
            if (this.openkeyword == true) {
                this.keywordopentitle = '收起';
            } else {
                this.keywordopentitle = '展开全部';
            }
        },
        changeAbstract() {
            this.openAbstract = !this.openAbstract;
            if (this.openAbstract == true) {
                this.abstractopentitle = '收起';
            } else {
                this.abstractopentitle = '展开全部';
            }
        },
        changeEpScanning() {
            this.openEpSumScanning = !this.openEpSumScanning;
            if (this.openEpSumScanning == true) {
                this.epopentitleScanning = '收起';
            } else {
                this.epopentitleScanning = '查看章节摘要';
            }
        },
        changeEpReview() {
            this.openEpSumReview = !this.openEpSumReview;
            if (this.openEpSumReview == true) {
                this.epopentitleReview = '收起';
            } else {
                this.epopentitleReview = '查看章节摘要';
            }
        },

        changeSearchStatus() {
            this.dictIndexID = {};
            this.dictSpanClass = {};
            this.searchInfo = '';
            this.isSearchContent = true;
        },

        handleSearchInputChange() {
            if (this.searchInfo === '') {
                this.isSearchInputError = false;
                this.isShowClearSearchInfo = false;
                this.isShowSwitchButton = false;
                this.isSearchNone = false;
            } else {
                this.isShowClearSearchInfo = true;
            }
        },
        handleSearchEnterKey() {
            if (this.searchInfo === '') {
                this.dictIndexID = {};
                this.dictSpanClass = {};
                this.isSearchInputError = true;
            } else {
                if (this.calcSearchContentInfo()) {
                    this.isShowClearSearchInfo = true;
                    this.isShowSwitchButton = true;
                    this.isSearchNone = false;
                    this.isSearchInputError = false;
                } else {
                    this.scrollToSpanID(this.dictIndexID[this.searchedCurNum]);
                    this.isShowSwitchButton = false;
                    this.isSearchNone = true;
                }
            }
        },
        calcSearchContentInfo() {
            if (this.searchInfo === '') {
                return false;
            } else {
                this.searchedCurNum = 1;
                let prefix_index = 0;
                let cur_index = 0;
                let cur_span_id = '';
                for (let i = 0; i < this.recording_items.length; i++) {
                    prefix_index = i;
                    let content_suffix_index = -1;
                    let trans_suffix_index = -1;

                    let i_content = this.concatContent(this.recording_items[i].content);
                    let i_trans_content = this.concatContent(this.recording_items[i].transContent);
                    if (i_content === '') continue;

                    // 源语言文本
                    if ((this.aiType === 0 && (this.showType === '0' || this.showType === '2')) || this.aiType === 2) {
                        let result_content = [];
                        let preIndex = 0;
                        let curIndex = 0;
                        while ((curIndex = i_content.indexOf(this.searchInfo, curIndex)) !== -1) {
                            result_content.push(i_content.substring(preIndex, curIndex));
                            content_suffix_index = content_suffix_index + 1;
                            cur_span_id = this.generateSpanID(prefix_index, content_suffix_index);
                            this.dictSpanClass[cur_span_id] = '';

                            result_content.push(i_content.substring(curIndex, curIndex + this.searchInfo.length));
                            content_suffix_index = content_suffix_index + 1;
                            cur_index = cur_index + 1;
                            cur_span_id = this.generateSpanID(prefix_index, content_suffix_index);
                            if (cur_index === this.searchedCurNum) {
                                this.dictSpanClass[cur_span_id] = 'searched-actived';
                            } else {
                                this.dictSpanClass[cur_span_id] = 'searched';
                            }
                            this.dictIndexID[cur_index] = cur_span_id;

                            curIndex += this.searchInfo.length; // 继续搜索下一个匹配位置
                            preIndex = curIndex;
                        }
                        result_content.push(i_content.substring(preIndex, i_content.length));
                        content_suffix_index = content_suffix_index + 1;
                        cur_span_id = this.generateSpanID(prefix_index, content_suffix_index);
                        this.dictSpanClass[cur_span_id] = '';
                        this.recording_items[i].content = result_content;
                    }

                    // 翻译文本
                    if ((this.aiType === 0 && (this.showType === '1' || this.showType === '2')) || this.aiType === 1 || this.aiType === 2) {
                        let trans_result_content = [];
                        let trans_preIndex = 0;
                        let trans_curIndex = 0;
                        while ((trans_curIndex = i_trans_content.indexOf(this.searchInfo, trans_curIndex)) !== -1) {
                            trans_result_content.push(i_trans_content.substring(trans_preIndex, trans_curIndex));
                            trans_suffix_index = trans_suffix_index + 1;
                            cur_span_id = this.generateTransSpanID(prefix_index, trans_suffix_index);
                            this.dictSpanClass[cur_span_id] = '';

                            trans_result_content.push(i_trans_content.substring(trans_curIndex, trans_curIndex + this.searchInfo.length));
                            trans_suffix_index = trans_suffix_index + 1;
                            cur_index = cur_index + 1;
                            cur_span_id = this.generateTransSpanID(prefix_index, trans_suffix_index);
                            if (cur_index === this.searchedCurNum) {
                                this.dictSpanClass[cur_span_id] = 'searched-actived';
                            } else {
                                this.dictSpanClass[cur_span_id] = 'searched';
                            }
                            this.dictIndexID[cur_index] = cur_span_id;

                            trans_curIndex += this.searchInfo.length; // 继续搜索下一个匹配位置
                            trans_preIndex = trans_curIndex;
                        }
                        trans_result_content.push(i_trans_content.substring(trans_preIndex, i_trans_content.length));
                        trans_suffix_index = trans_suffix_index + 1;
                        cur_span_id = this.generateSpanID(prefix_index, trans_suffix_index);
                        this.dictSpanClass[cur_span_id] = '';
                        this.recording_items[i].transContent = trans_result_content;
                    }
                }

                if (cur_index > 0) {
                    this.searchedTotalNum = cur_index;
                    return true;
                } else {
                    return false;
                }
            }
        },
        clearSearchInfo() {
            this.isShowClearSearchInfo = false;
            this.isShowSwitchButton = false;
            this.isSearchNone = false;
            this.searchInfo = '';
            this.dictIndexID = {};
            this.dictSpanClass = {};
        },
        leftSwitchButton() {
            let origin_searchedCurNum = this.searchedCurNum;
            if (this.searchedCurNum === 1) {
                this.searchedCurNum = this.searchedTotalNum;
            } else {
                this.searchedCurNum = this.searchedCurNum - 1;
            }
            this.dictSpanClass[this.dictIndexID[origin_searchedCurNum]] = 'searched';
            this.dictSpanClass[this.dictIndexID[this.searchedCurNum]] = 'searched-actived';
            this.scrollToSpanID(this.dictIndexID[this.searchedCurNum]);
        },
        rightSwitchButton() {
            let origin_searchedCurNum = this.searchedCurNum;
            this.searchedCurNum = (this.searchedCurNum % this.searchedTotalNum) + 1;
            this.dictSpanClass[this.dictIndexID[origin_searchedCurNum]] = 'searched';
            this.dictSpanClass[this.dictIndexID[this.searchedCurNum]] = 'searched-actived';
            this.scrollToSpanID(this.dictIndexID[this.searchedCurNum]);
        },
        cancelSearchStatus() {
            this.isSearchContent = false;
            this.isSearchInputError = false;
            this.isShowClearSearchInfo = false;
            this.isShowSwitchButton = false;
            this.isSearchNone = false;
            this.searchInfo = '';

            for (let i = 0; i < this.recording_items.length; i++) {
                let i_content = this.concatContent(this.recording_items[i].content);
                this.recording_items[i].content = [i_content];
            }

            if ((this.aiType === 0 && (this.showType === '1' || this.showType === '2')) || this.aiType === 1 || this.aiType === 2) {
                for (let i = 0; i < this.recording_items.length; i++) {
                    let i_trans_content = this.concatContent(this.recording_items[i].transContent);
                    this.recording_items[i].transContent = [i_trans_content];
                }
            }

            this.dictIndexID = {};
            this.dictSpanClass = {};
        },
        customSleep(ms) {
            return new Promise((resolve) => {
                setTimeout(resolve, ms);
            });
        },
        async allReplace() {
            if (this.originText === '') {
                this.$message({
                    message: '请输入要替换的原文本内容',
                    type: 'error',
                    offset: 500
                });
            } else if (this.replaceNum === 0) {
                this.$message({
                    message: '没有可替换的文本内容',
                    type: 'info',
                    offset: 500
                });
            } else {
                for (let i = 0; i < this.recording_items.length; i++) {
                    let i_replace = false;
                    if ((this.aiType === 0 && (this.showType === '0' || this.showType === '2')) || this.aiType === 0) {
                        for (let j = 0; j < this.recording_items[i].content.length; j++) {
                            if (this.recording_items[i].content[j] === this.originText) {
                                this.recording_items[i].content[j] = this.replaceText;
                                i_replace = true;
                            }
                        }
                    }

                    if ((this.aiType === 0 && this.showType === '1') || this.aiType === 1 || this.aiType === 2) {
                        for (let j = 0; j < this.recording_items[i].transContent.length; j++) {
                            if (this.recording_items[i].transContent[j] === this.originText) {
                                this.recording_items[i].transContent[j] = this.replaceText;
                                i_replace = true;
                            }
                        }
                    }

                    if (i_replace && this.showType === '2' && this.aiType === 0) {
                        await this.customSleep(50);
                        let i_content = this.concatContent(this.recording_items[i].content);
                        voiceTrans(i_content, this.sourceLanguage, this.destLanguage).then((response) => {
                            this.recording_items[i].transContent = [response.data];
                        });
                    }
                }
                this.$message({
                    message: '一键替换成功',
                    type: 'success',
                    offset: 500
                });
            }
            this.replaceNum = 0;
            this.dictIndexID = {};
            this.dictSpanClass = {};
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
.editorPDFViewport {
    height: 100%;
    background-color: #fff;
    border-radius: 10px !important;
    padding: 40px 40px 140px;
    position: relative;
    z-index: 1;
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

.solid-line {
    border-top: 1px solid #e9e1e1;
    width: calc(100% - 26px);
    height: 1px;
    margin: 0;
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
.keywordleft {
    height: 56px;
}

.breakabstractword {
    font-size: 14px;
    line-height: 24px;
    overflow: hidden;
    text-align: justify;
    color: rgb(71, 70, 103);
    transition: height 0.3s ease 0s;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 3;
    line-height: 1.2em;
    word-break: break-all;
    text-overflow: ellipsis;
}

.openabstractword {
    font-size: 14px;
    line-height: 24px;
    overflow: hidden;
    text-align: justify;
    color: rgb(71, 70, 103);
    transition: height 0.3s ease 0s;
    line-height: 1.2em;
    word-break: break-all;
}

.djVVR {
    font-size: 14px;
    line-height: 24px;
    opacity: 0;
    position: absolute;
    top: 0px;
    left: 0px;
    width: 100%;
    z-index: -1;
    pointer-events: none;
}

.aisumtext {
    font-size: 14px;
    letter-spacing: 0px;
    box-sizing: border-box;
    overflow-wrap: anywhere;
    position: relative;
    z-index: 2;
    color: rgb(39, 38, 77);
    cursor: pointer;
    -webkit-font-smoothing: antialiased;
}

.breakaisum {
    word-break: break-all;
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    -webkit-line-clamp: 1;
    -webkit-box-orient: vertical;
}

.heng > div > div:nth-child(2n + 1) {
    box-shadow: rgb(255, 255, 255) 0px 1px 0px;
}

.ccTlNE {
    margin-right: 14px;
    flex: 1 1 0%;
    height: 1px;
    background-color: rgba(39, 38, 77, 0.1);
}

.heng > div > div:nth-child(2n + 1) {
    box-shadow: rgb(255, 255, 255) 0px 1px 0px;
}

.jaCZOE {
    margin-left: 14px;
    flex: 1 1 0%;
    height: 1px;
    background-color: rgba(39, 38, 77, 0.1);
}

.search-top {
    display: flex;
    -webkit-box-align: center;
    align-items: center;
}

.search-top-left {
    border-radius: 10px;
    border: 2px solid rgba(96, 92, 229, 0);
}

.search-top-left:hover {
    border-color: rgba(102, 62, 166, 0.493);
}
.search-top-left.focus {
    border-color: rgba(105, 59, 179, 0.764);
    box-shadow: 0px 0px 8px rgba(105, 59, 179, 0.4); /* 添加更强的模糊效果 */
}

.search-input {
    width: 400px;
    border-radius: 8px;
    padding: 4px 12px;
    font-size: 20px;
    height: 36px;
    border: 1px solid rgb(233, 233, 237);
    display: flex;
    -webkit-box-align: center;
    align-items: center;
    background: rgb(255, 255, 255);
}

.cancel {
    font-size: 12px;
    color: rgba(39, 38, 77, 0.65);
    display: flex;
    -webkit-box-align: center;
    align-items: center;
    -webkit-box-pack: justify;
    justify-content: space-between;
    height: 18px;
    width: 42px;
    margin-left: 7px;
    cursor: pointer;
}

.line {
    display: inline-block;
    width: 1px;
    height: 18px;
    background-color: rgb(233, 233, 237);
    margin-right: 12px;
    border: 0px;
}

.cancelText {
    cursor: pointer;
    display: block;
    width: 29px;
    font-size: 14px;
    font-weight: 400;
}
.cancelText:hover {
    color: rgb(86, 24, 187); /* 悬停时的字体颜色 */
}

.search-input-field {
    width: 300px;
    background: none;
    margin: 0px 7px;
    outline: none;
    border: 0px solid rgb(204, 204, 204);
    font-weight: 400;
    font-size: 14px;
    color: rgb(39, 38, 77);
}

.search-error {
    width: 131px;
    font-size: 12px;
    color: rgb(255, 128, 128);
    display: flex;
}

.search-line {
    width: 1px;
    height: 18px;
    background-color: rgb(233, 233, 237);
    margin-left: 12px;
    margin-right: 10px;
}

.switchover {
    display: flex;
    -webkit-box-align: center;
    align-items: center;
}

.switchover-button {
    width: 24px;
    height: 24px;
    -webkit-box-pack: center;
    justify-content: center;
    -webkit-box-align: center;
    align-items: center;
    display: flex;
    cursor: pointer;
}

.switchover-inputNum {
    font-weight: 400;
    font-size: 12px;
    color: rgba(39, 38, 77, 0.45);
    height: 24px;
    line-height: 24px;
    user-select: none;
}

.clear-searchInfo {
    cursor: pointer;
}

.clear-searchInfo:hover path {
    fill: #888; /* 悬停时颜色加深 */
}

.searched {
    background-color: rgb(255, 223, 164);
    color: rgba(39, 38, 77, 0.85);
}

.searched-actived {
    background-color: rgba(159, 156, 238, 0.6);
    color: rgba(74, 77, 38, 0.85);
}

.audio-player {
    background-color: #f0f8ff; /* 温柔的背景色 */
    border-radius: 15px; /* 圆角边框 */
    padding: 20px; /* 内边距 */
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1); /* 阴影效果 */
    width: calc(100% - 40px); /* 适应父容器宽度 */
    margin: 0 auto; /* 居中对齐 */
}

.audio-control {
    width: 100%; /* 音频控件宽度适应容器 */
    border: none; /* 去掉默认边框 */
    outline: none; /* 去掉默认轮廓 */
}

.introduction {
    height: 100%;
    display: flex;
    flex-direction: column;
    background: rgb(255, 255, 255);
}

.introduction-content {
    flex: 1 1;
}

.taskTitle {
    flex: 1 1;
    display: inline-block;
    font-family: PingFang SC;
    font-size: 20px;
    font-weight: 500;
    -webkit-font-smoothing: antialiased;
    line-height: 40px;
    letter-spacing: 0px;
    margin-bottom: 12px;
}

.title-wrapper {
    margin-bottom: 8px;
    display: flex;
    align-items: center;
    justify-content: space-between;
}

.title {
    flex: 1 1;
    font-family: PingFang SC;
    font-size: 14px;
    font-weight: 600;
    -webkit-font-smoothing: antialiased;
    line-height: 24px;
    letter-spacing: 0px;
}

.card {
    border-radius: 8px;
}

.markdown-body {
    font-weight: 400;
    font-size: 14px;
    color: rgb(86, 98, 121);
    line-height: 24px;
}

ol {
    list-style-type: decimal;
}
</style>
