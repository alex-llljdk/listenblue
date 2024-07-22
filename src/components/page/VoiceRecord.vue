<template>
    <div class="h-full w-full flex">
        <div class="h-full" style="background: #e9efff">
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
                <div v-if="showRecording" class="text-center mt-8" @click="handleSave">
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
                </div>
            </el-aside>
        </div>
        <div class="flex h-full w-full" style="background: #f2f5fb">
            <div class="jphwIe">
                <div class="WFeMQ">
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
                            <div v-if="showRecording">
                                <span class="text-xs text-gray-400 select-none ml-2">{{ savedTime }}</span>
                            </div>
                            <div v-else class="select-none ml-2">&nbsp;</div>
                        </div>

                        <div v-if="showRecording" class="flex text-lg items-center">
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
                        </div>
                    </div>
                </div>
                <div class="h-full w-full">
                    <div v-if="showRecording" class="h-full w-full">
                        <div class="dDPXpc" style="overflow: auto" ref="scrollableDiv">
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

                                        <!-- <div v-if="!item_recording.isDivTrans" class="solid-line"></div> -->

                                        <div v-if="translation === '0' && !item_recording.isDivTrans" class="solid-line"></div>

                                        <div v-if="translation === '0'" :class="item_recording.isDivTrans ? 'transJItVgd' : 'jItVgd'">
                                            <div class="kCofWf">
                                                <div class="kFJVLr w-full">
                                                    <span
                                                        v-for="(trans_item_content, trans_content_index) in recording_items[index]
                                                            .transContent"
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
                        <div class="kGnJQu">
                            <div class="bottomContain">
                                <div class="content flex">
                                    <div class="h-28" v-if="recording">
                                        <img src="../../assets/img/recording.gif" class="h-28 w-16" />
                                    </div>
                                    <div class="h-28" v-else>
                                        <img src="../../assets/img/recording.jpg" class="h-28 w-16" />
                                    </div>

                                    <div class="running-content">
                                        <div class="running-text">
                                            <span>{{ recording ? '录音中...' : '录音已暂停' }}</span>
                                        </div>
                                        <div class="running-time">
                                            <span>{{ formatTime(currentTime) }}</span>
                                            <span style="opacity: 0.45; margin: 0px 5px">/</span>
                                            <span style="opacity: 0.45">01:00:00</span>
                                        </div>
                                    </div>
                                </div>
                                <div class="right-btn">
                                    <div class="red_container" title="结束录音" @click="centerDialogVisible = true">
                                        <svg t="1714747107927" class="icon" viewBox="0 0 1024 1024" width="25" height="25">
                                            <path
                                                d="M718.9248 239.328c24.1984 17.2096 45.7152 36.704 64.5376 58.4896s34.9568 45.312 48.4032 70.592c13.4464 25.28 23.6608 52.032 30.6496 80.2688 6.9952 28.2368 10.4896 56.8704 10.4896 85.9136 0 50.016-9.5424 96.9408-28.64 140.7744-19.0848 43.8336-44.9088 82.016-77.4464 114.5536s-70.7264 58.3552-114.5536 77.4464C608.5376 886.4576 561.6192 896 511.5968 896c-49.472 0-96.1344-9.5424-139.9616-28.64-43.8336-19.0848-82.1504-44.9088-114.9568-77.4464s-58.6176-70.7264-77.4464-114.5536-28.2368-90.7584-28.2368-140.7744c0-28.4992 3.36-56.4736 10.0864-83.8976 6.72-27.4304 16.2624-53.5104 28.64-78.2464 12.3712-24.7424 27.6928-47.872 45.984-69.376 18.2848-21.5168 38.72-40.8768 61.3056-58.0928 11.84-8.6016 24.608-11.8272 38.3232-9.6768 13.7152 2.1504 24.8704 8.8768 33.472 20.1664 8.608 11.296 11.84 23.936 9.6896 37.92-2.1568 13.984-8.8768 25.28-20.1728 33.8816C324.4352 352 298.4896 382.3872 280.4736 418.4192c-18.016 36.032-27.0336 74.7584-27.0336 116.1664 0 35.5008 6.7264 68.9728 20.1728 100.4352 13.4464 31.4624 31.8656 58.8928 55.2576 82.2848 23.3984 23.392 50.8224 41.952 82.2848 55.6608 31.4624 13.7216 64.9344 20.576 100.4288 20.576 35.5008 0 68.9792-6.8544 100.4352-20.576 31.4624-13.7152 58.8928-32.2688 82.2848-55.6608 23.3984-23.392 41.952-50.8224 55.6608-82.2848 13.7216-31.4624 20.576-64.9344 20.576-100.4352 0-41.952-9.6832-81.6128-29.0432-118.9888s-46.5216-68.1664-81.472-92.3712c-11.84-8.064-18.9632-19.0912-21.3824-33.0688-2.4192-13.9904 0.4032-26.8928 8.4672-38.7264 8.0704-11.296 19.0912-18.1504 33.0752-20.5696C694.1824 228.4352 707.0912 231.264 718.9248 239.328L718.9248 239.328zM511.5968 537.0048c-13.984 0-25.9456-4.9728-35.8912-14.9184-9.952-9.952-14.9248-21.92-14.9248-35.904L460.7808 179.6288c0-13.984 4.9728-26.08 14.9248-36.3008S497.6128 128 511.5968 128c14.528 0 26.7648 5.1072 36.704 15.328 9.9584 10.2208 14.9312 22.3168 14.9312 36.3008l0 306.5536c0 13.984-4.9728 25.952-14.9312 35.904C538.3552 532.032 526.1184 537.0048 511.5968 537.0048L511.5968 537.0048zM511.5968 537.0048"
                                                fill="#d1685a"
                                                p-id="13331"
                                            />
                                        </svg>
                                    </div>

                                    <el-dialog
                                        title="结束录音"
                                        :visible.sync="centerDialogVisible"
                                        width="16%"
                                        top="40vh"
                                        class="font-bold"
                                    >
                                        <span>结束后无法在本记录继续录音</span>
                                        <span slot="footer" class="dialog-footer">
                                            <el-button @click="noSaveClose">直接退出</el-button>
                                            <el-button type="primary" @click="saveClose">保存并退出</el-button>
                                        </span>
                                    </el-dialog>

                                    <div class="blue_container" v-if="recording" @click="toggleDivs" title="暂停录音">
                                        <svg t="1714749774206" class="icon" viewBox="0 0 1024 1024" width="25" height="25">
                                            <path d="M640 832V192h128v640h-128zM256 192h128v640H256V192z" fill="#0590DF" p-id="15295" />
                                        </svg>
                                    </div>

                                    <div class="blue_container" v-else @click="toggleDivs" title="继续录音">
                                        <svg t="1714746776172" class="icon" viewBox="0 0 1024 1024" width="25" height="25">
                                            <path
                                                d="M512 648c92.8 0 168-75.2 168-168V304c0-92.8-75.2-168-168-168S344 211.2 344 304v176c0 92.8 75.2 168 168 168zM392 304c0-65.6 54.4-120 120-120s120 54.4 120 120v176c0 65.6-54.4 120-120 120s-120-54.4-120-120V304z"
                                                fill="#2090E0"
                                                p-id="3768"
                                            />
                                            <path
                                                d="M768 408c-12.8 0-24 11.2-24 24v56C744 616 640 720 512 720s-232-104-232-232V432c0-12.8-11.2-24-24-24s-24 11.2-24 24v56c0 145.6 113.6 267.2 256 278.4v73.6h-64c-12.8 0-24 11.2-24 24s11.2 24 24 24h176c12.8 0 24-11.2 24-24s-11.2-24-24-24h-64v-73.6c142.4-12.8 256-132.8 256-278.4V432c0-12.8-11.2-24-24-24z"
                                                fill="#2090E0"
                                                p-id="3769"
                                            />
                                        </svg>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div v-else class="ixtgUW">
                        <div class="kGfivz">
                            <div class="epacMt">
                                <img src="../../assets/img/voice_record.jpg" />
                            </div>

                            <div class="w-2/5">
                                <div class="hObNWG">
                                    <div>音频语言</div>
                                </div>
                                <div class="pb-3 pt-1 w-full">
                                    <el-radio-group v-model="language" class="w-full">
                                        <el-radio-button label="0" class="w-1/2">中文</el-radio-button>
                                        <el-radio-button label="1" class="w-1/2">英文</el-radio-button>
                                    </el-radio-group>
                                </div>

                                <div class="hObNWG">
                                    <div>翻译</div>
                                </div>
                                <div class="pb-7 pt-1 w-full">
                                    <el-radio-group v-model="translation" class="w-full">
                                        <el-radio-button label="0" class="w-1/2">翻译</el-radio-button>
                                        <el-radio-button label="1" class="w-1/2">不翻译</el-radio-button>
                                    </el-radio-group>
                                </div>
                            </div>

                            <div class="flex justify-center w-2/5 h-16">
                                <button class="custom-button" @click="startRecording">开始录音</button>
                            </div>
                            <div class="justify-center w-2/5 h-48"></div>
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
import { Loading, Message } from 'element-ui';
import { quillEditor } from 'vue-quill-editor';
import { voiceTrans } from '../../api/voiceTrans';
import { saveVoice } from '../../api/voiceTrans';
import { userStore } from '../../store/store';
import 'quill/dist/quill.core.css';
import 'quill/dist/quill.snow.css';
import 'quill/dist/quill.bubble.css';
import lamejs from 'lamejs';
const userstore = userStore();
// var voiceArray = [];
//Recorder
var Recorder = function (stream, socket) {
    var sampleBits = 16; //输出采样数位 8, 16
    var sampleRate = 16000; //输出采样率
    console.log('sampleBits: ', sampleBits, 'sampleRate: ', sampleRate);
    var context = new AudioContext();
    var audioInput = context.createMediaStreamSource(stream);
    var recorder = context.createScriptProcessor(1024, 1, 1);
    var wholeAudioDataBuffer = []; // 整个录音过程的音频数据
    var ws = socket;
    var sendSize = 1280; // 每次send大小

    var audioData = {
        size: 0, //录音文件长度
        allSize: 0,
        allbuffer: [], //所有
        buffer: [], //录音缓存
        inputSampleRate: 32000, //输入采样率
        inputSampleBits: 16, //输入采样数位 8, 16
        outputSampleRate: sampleRate, //输出采样数位
        oututSampleBits: sampleBits, //输出采样率
        clear: function () {
            this.buffer = [];
            this.size = 0;
        },
        input: function (data) {
            this.buffer.push(new Float32Array(data));
            this.allbuffer.push(new Float32Array(data));
            this.size += data.length;
            this.allSize += data.length;
        },
        compress: function () {
            //合并压缩
            //合并
            var data = new Float32Array(this.size);
            var offset = 0;
            for (var i = 0; i < this.buffer.length; i++) {
                data.set(this.buffer[i], offset);
                offset += this.buffer[i].length;
            }
            //压缩
            var compression = parseInt(this.inputSampleRate / this.outputSampleRate);
            //console.log(compression)
            var length = data.length / compression;
            var result = new Float32Array(length);
            var index = 0,
                j = 0;
            while (index < length) {
                result[index] = data[j];
                j += compression;
                index++;
            }
            return result;
        },
        encodePCM: function () {
            //这里不对采集到的数据进行其他格式处理，如有需要均交给服务器端处理。
            var sampleRate = Math.min(this.inputSampleRate, this.outputSampleRate);
            var sampleBits = Math.min(this.inputSampleBits, this.oututSampleBits);
            var bytes = this.compress();
            var dataLength = bytes.length * (sampleBits / 8);
            var buffer = new ArrayBuffer(dataLength);
            var data = new DataView(buffer);
            var offset = 0;
            for (var i = 0; i < bytes.length; i++, offset += 2) {
                var s = Math.max(-1, Math.min(1, bytes[i]));
                data.setInt16(offset, s < 0 ? s * 0x8000 : s * 0x7fff, true);
            }
            return new Blob([data]);
        },

        compressAll: function () {
            //合并压缩
            //合并
            var data = new Float32Array(this.allSize);
            console.log(this.allSize);
            console.log(this.allbuffer.length);
            var offset = 0;
            for (var i = 0; i < this.allbuffer.length; i++) {
                data.set(this.allbuffer[i], offset);
                offset += this.allbuffer[i].length;
            }
            console.log(122211);
            //压缩
            var compression = parseInt(this.inputSampleRate / this.outputSampleRate);
            //console.log(compression)
            var length = data.length / compression;
            var result = new Float32Array(length);
            var index = 0,
                j = 0;
            while (index < length) {
                result[index] = data[j];
                j += compression;
                index++;
            }
            return result;
        },
        encodePCMAll: function () {
            //这里不对采集到的数据进行其他格式处理，如有需要均交给服务器端处理。
            var sampleRate = Math.min(this.inputSampleRate, this.outputSampleRate);
            var sampleBits = Math.min(this.inputSampleBits, this.oututSampleBits);
            var bytes = this.compressAll();
            var dataLength = bytes.length * (sampleBits / 8);
            var buffer = new ArrayBuffer(dataLength);
            var data = new DataView(buffer);
            var offset = 0;
            for (var i = 0; i < bytes.length; i++, offset += 2) {
                var s = Math.max(-1, Math.min(1, bytes[i]));
                data.setInt16(offset, s < 0 ? s * 0x8000 : s * 0x7fff, true);
            }
            return new Blob([data]);
        }
    };

    var sendData = function () {
        //对以获取的数据进行处理(分包)
        var reader = new FileReader();
        reader.onload = (e) => {
            var outbuffer = e.target.result;
            var arr = new Uint8Array(outbuffer);
            // console.log('arr: ', arr);
            if (arr.length > 0) {
                var tmparr = new Uint8Array(sendSize);
                var j = 0;
                for (var i = 0; i < arr.byteLength; i++) {
                    // wholeAudioDataBuffer.push(1);
                    // voiceArray.push(arr[i]);
                    tmparr[j++] = arr[i];
                    if ((i + 1) % sendSize == 0) {
                        // console.log("tmparr: ", tmparr);
                        ws.send(tmparr);

                        if (arr.byteLength - i - 1 >= sendSize) {
                            tmparr = new Uint8Array(sendSize);
                        } else {
                            tmparr = new Uint8Array(arr.byteLength - i - 1);
                        }
                        j = 0;
                    }
                    if (i + 1 == arr.byteLength && (i + 1) % sendSize != 0) {
                        ws.send(tmparr);
                    }
                }
            }
        };
        reader.readAsArrayBuffer(audioData.encodePCM());
        audioData.clear(); //每次发送完成则清理掉旧数据
    };

    this.start = function () {
        audioInput.connect(recorder);
        recorder.connect(context.destination);
    };

    this.stop = function () {
        recorder.disconnect();
    };

    this.getBufferAll = function () {
        return audioData.allbuffer;
    };

    this.getBlob = function () {
        return audioData.encodePCMAll();
    };
    this.clear = function () {
        audioData.clear();
    };

    recorder.onaudioprocess = function (e) {
        var inputBuffer = e.inputBuffer.getChannelData(0);
        audioData.input(inputBuffer);
        sendData();
        // console.log('wholeAudioDataBuffer: ', wholeAudioDataBuffer);
        // console.log('发送音频流');
    };
};

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
    components: { quillEditor },
    name: 'VoiceRecord',
    data() {
        return {
            // recording_items 数组用于存储要动态生成的 div 元素的数据
            recording_items: [],
            item_index: -1, // recording_items当前最大进行时item的下标
            curSentences: [], // 当前item_index正在识别的句子数组
            isAddItem: true, // curSentences最后一个句子必须是onebest状态才可以添加recording_item
            isVar: false, // curSentences最后一个元素是否是var状态
            activeTime: 0, // curSentences最后一个句子的时间戳

            originText: '', // 原文本内容
            replaceText: '', // 替换文本

            currentTime: 0, // 当前时间，单位：秒
            maxTime: 3600, // 最大时间，单位：秒，即 01:00:00
            timer: null, // 定时器

            showRecording: false, // 控制是否显示录音时的内容
            recording: true, // 是否正在录音
            inputWidth: 204,
            recordInfo: '',
            savedTime: '',
            isComposing: false, // 标志位，用于跟踪输入法输入状态

            language: '0', // 0 中文, 1 英文
            translation: '0', // 0 翻译, 1 不翻译
            sourceLanguage: 'zh-CHS',
            destLanguage: 'en',
            transTimer: null, // 用于翻译的毫秒定时器

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
            wsUrl: 'ws://localhost:10003/voice/ws',
            recorder: null, //多媒体对象，用来处理音频
            curVoiceIndex: 0, // 已读音频数据的下标

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
            centerDialogVisible: false,
            saveFileName: '',

            mediaStream: null
        };
    },
    mounted() {
        document.title = 'ListenBlue-会议记录';
        this.updateInfo();
        this.initTitle();
        this.saveFileName = this.getSaveFileName;
    },
    created() {},
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
            return `${year}-${month}-${date} ${hours}:${minutes}`;
        },
        editor() {
            return this.$refs.myQuillEditor.quill;
        },
        isReplaceAvailable() {
            return false;
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
        }
    },
    beforeDestroy() {
        console.log('Leaving the route...');
        this.clearRecorderAndCloseSocket();
        this.closeMicrophone();
    },
    methods: {
        closeMicrophone() {
            if (this.mediaStream) {
                let tracks = this.mediaStream.getAudioTracks();
                tracks.forEach((track) => track.stop());
                this.mediaStream = null;
                console.log('麦克风已关闭');
            } else {
                console.log('没有麦克风需要关闭');
            }
        },
        clearRecorderAndCloseSocket() {
            if (this.recorder !== null) {
                this.recorder.stop();
                this.recorder.clear();
                if (this.socket.readyState === WebSocket.CLOSED) {
                    console.log('Socket is closed');
                } else {
                    console.log('Socket is not closed');
                    this.socket.close();
                }
            }
        },
        getHMSTime() {
            const now = new Date();
            const hours = String(now.getHours()).padStart(2, '0');
            const minutes = String(now.getMinutes()).padStart(2, '0');
            const seconds = String(now.getSeconds()).padStart(2, '0');
            return `${hours}:${minutes}:${seconds}`;
        },
        getCurSavedTime() {
            const hms = this.getHMSTime();
            return `已保存于 ${hms}`;
        },
        updateInfo() {
            const ymd_hms = this.getYMDHMSTime;
            const hms = this.getHMSTime();
            this.recordInfo = `${ymd_hms} 记录`;
            this.savedTime = this.getCurSavedTime();
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
        scrollToBottom() {
            const div = this.$refs.scrollableDiv;
            div.scrollTop = div.scrollHeight;
        },
        addRecordingItem() {
            // 添加一个新的item到数组中
            if (this.isAddItem) {
                this.curSentences = [];
                this.recording_items.push({
                    time_stamp: this.formatTime(this.currentTime),
                    content: [],
                    isActive: false,
                    isDivTrans: true,
                    transContent: [],
                    isTrans: false, // 与定时器配合, 当前Item是否需要请求翻译
                    transMillisecondStamp: 0 // 翻译时的毫秒时间戳
                });
                this.item_index++;

                if (this.translation === '1') {
                    this.recording_items[this.item_index].isDivTrans = false;
                }

                if (this.item_index > 0) {
                    this.recording_items[this.item_index - 1].isDivTrans = false;
                }
            }
            this.$nextTick(() => {
                this.scrollToBottom();
            });
        },
        deleteItem() {
            // 检查数组长度是否大于0
            if (this.recording_items.length > 0) {
                // 检查最后一个item的content是否为空, 删除最后一个item
                let last_item_content = this.concatContent(this.recording_items[this.recording_items.length - 1].content);
                if (last_item_content.length === 0) {
                    this.recording_items.pop();
                    this.item_index--;
                }
            }
            this.$nextTick(() => {
                this.scrollToBottom();
            });
        },
        toggleDivs() {
            this.recording = !this.recording; // 录音与暂停状态切换
            if (this.recording) {
                this.startTimer();
                this.addRecordingItem();
                console.log(this.item_index);
                this.recorder.start();
            } else {
                this.stopTimer();
                this.deleteItem();
                this.recorder.stop();
            }
        },
        customSleep(ms) {
            return new Promise((resolve) => {
                setTimeout(resolve, ms);
            });
        },
        startTimer() {
            this.timer = setInterval(() => {
                this.currentTime++;
                if (this.currentTime - this.activeTime >= 5 && this.isAddItem === true && this.curSentences.length > 0) {
                    this.addRecordingItem();
                }
            }, 1000); // Increment time every second
        },
        stopTimer() {
            clearInterval(this.timer);
        },
        async startTransTimer() {
            this.transTimer = setInterval(async () => {
                if (this.translation === '0') {
                    for (let i = 0; i < this.recording_items.length; i++) {
                        let i_time = Date.now() - this.recording_items[i].transMillisecondStamp;
                        if (this.recording_items[i].isTrans && i_time >= 500) {
                            let i_content = this.concatContent(this.recording_items[i].content);
                            voiceTrans(i_content, this.sourceLanguage, this.destLanguage).then((response) => {
                                this.recording_items[i].transContent = [response.data];
                            });
                            this.recording_items[i].isTrans = false;
                            this.recording_items[i].transMillisecondStamp = Date.now();
                            await this.customSleep(500);
                        }
                    }
                }
            }, 5); // Increment time every second
        },
        stopTransTimer() {
            clearInterval(this.transTimer);
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
        getCaretPosition(element) {
            let caretOffset = 0;
            const doc = element.ownerDocument || element.document;
            const win = doc.defaultView || doc.parentWindow;
            const sel = win.getSelection();
            if (sel.rangeCount > 0) {
                const range = sel.getRangeAt(0);
                const preCaretRange = range.cloneRange();
                preCaretRange.selectNodeContents(element);
                preCaretRange.setEnd(range.startContainer, range.startOffset);
                caretOffset = preCaretRange.toString().length;
            }
            return caretOffset;
        },
        setCaretPosition(element, offset) {
            const doc = element.ownerDocument || element.document;
            const win = doc.defaultView || doc.parentWindow;
            const range = doc.createRange();
            const sel = win.getSelection();

            let charIndex = 0;
            let foundStart = false,
                stop = false;

            (function traverseNodes(node) {
                if (stop) return;

                if (node.nodeType === 3) {
                    // Text node
                    const nextCharIndex = charIndex + node.length;
                    if (!foundStart && offset >= charIndex && offset <= nextCharIndex) {
                        range.setStart(node, offset - charIndex);
                        foundStart = true;
                    }
                    if (foundStart && offset >= charIndex && offset <= nextCharIndex) {
                        range.setEnd(node, offset - charIndex);
                        stop = true;
                    }
                    charIndex = nextCharIndex;
                } else {
                    for (let i = 0; i < node.childNodes.length; i++) {
                        traverseNodes(node.childNodes[i]);
                        if (stop) break;
                    }
                }
            })(element);

            sel.removeAllRanges();
            sel.addRange(range);
        },
        onCompositionStart() {
            this.isComposing = true;
        },
        onCompositionEnd(index, event) {
            this.isComposing = false;
            this.onContentChange(index, event); // 输入法完成时触发内容变化处理
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

            if (this.translation === '0') {
                this.recording_items[index].isTrans = true;
            }

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
        inDiv(index, event) {
            this.recording_items[index].isActive = true;
        },
        outDiv(index, event) {
            this.recording_items[index].isActive = false;
        },
        //启动录音
        startRecording() {
            console.log('语言：', this.language);
            console.log('翻译：', this.translation);
            // 设置翻译的源语言，目标语言
            if (this.translation === '0') {
                if (this.language === '0') {
                    this.sourceLanguage = 'zh-CHS';
                    this.destLanguage = 'en';
                } else {
                    this.sourceLanguage = 'en';
                    this.destLanguage = 'zh-CHS';
                }
            }

            navigator.getUserMedia = navigator.getUserMedia || navigator.webkitGetUserMedia;
            if (!navigator.getUserMedia) {
                console.error('无法访问麦克风:', error);
                alert('无法访问麦克风，请检查权限设置或硬件连接');
            } else {
                var That = this;
                navigator.getUserMedia(
                    {
                        audio: true
                    },
                    function (mediaStream) {
                        // 麦克风可用，执行成功时的操作
                        That.mediaStream = mediaStream;
                        That.initWebSocketAndRecording(mediaStream);
                    },
                    function (error) {
                        console.error('无法访问麦克风:', error);
                        alert('无法访问麦克风，请检查权限设置或硬件连接');
                    }
                );
            }
        },
        initWebSocketAndRecording(mediaStream) {
            console.log('initWesocket==');
            // 建立WebSocket连接
            this.socket = new WebSocket(this.wsUrl);

            // 监听WebSocket的open事件，当连接打开时触发
            this.socket.onopen = (event) => {
                console.log('socket.onopen成功');
                this.recorder = new Recorder(mediaStream, this.socket);

                // setInterval(() => {
                //     this.socket.send("111111111111111111111");
                //     this.socket.send("222222222222222222222");
                // }, 1000);
            };
            // 监听WebSocket的message事件，当收到服务器消息时触发
            this.socket.onmessage = (event) => {
                var data = JSON.parse(event.data);
                console.log('onmessage: ', data);
                switch (
                    data.action //根据WebSocket返回值的某个字段，区分改做什么事情
                ) {
                    case 'started': // // {"type":"connect","data":"7f0000010b540000000a"}
                        // 保证socket连接成功再开始录音
                        console.log('recorder 开始录音');
                        this.showRecording = true;
                        if (this.recording) {
                            this.addRecordingItem();
                            this.item_index = 0;
                            this.startTimer();
                            this.startTransTimer();
                        }
                        this.recorder.start();

                        //可删除，只是用来测试后端自动断开连接的时间
                        this.second = 0;
                        setInterval(() => {
                            this.second++;
                        }, 1000);
                        break;
                    case 'result':
                        if (data.data.var !== undefined) {
                            if (this.isVar) {
                                this.curSentences[this.curSentences.length - 1] = data.data.var;
                            } else {
                                this.curSentences.push(data.data.var);
                            }
                            this.isAddItem = false;
                            this.isVar = true;
                        } else if (data.data.onebest !== undefined) {
                            if (this.isVar) {
                                this.curSentences[this.curSentences.length - 1] = data.data.onebest;
                            } else {
                                this.curSentences.push(data.data.onebest);
                            }
                            this.isAddItem = true;
                            this.isVar = false;
                        }

                        let result_content = '';
                        for (let i = 0; i < this.curSentences.length; i++) {
                            result_content += this.curSentences[i];
                        }
                        this.recording_items[this.item_index].content = [result_content];
                        if (this.translation === '0') {
                            this.recording_items[this.item_index].isTrans = true;
                        }

                        if (data.data.onebest !== undefined && this.curSentences.length >= 3) {
                            this.addRecordingItem();
                        }

                        break;
                }
                this.activeTime = this.currentTime;

                // console.log('socket.onmessage==', event.data);
                // var data = JSON.parse(event.data);
                // switch (
                //     data.type //根据WebSocket返回值的某个字段，区分改做什么事情
                // ) {
                //     case 'connect': // // {"type":"connect","data":"7f0000010b540000000a"}
                //         //写你自己的逻辑
                //         break;
                //     case 'xxx':
                //         //写你自己的逻辑
                //         break;
                // }
            };
            // 监听WebSocket的close事件，当连接关闭时触发
            this.socket.onclose = (event) => {
                console.log('socket.onclose==');
                console.log(`WebSocket连接${this.second}秒后关闭了`);
                console.log('WebSocket is closed. Code: ' + event.code + ' Reason: ' + event.wasClean);
                //连接关闭，就重新连接
                // this.reconnect();
            };

            // 监听WebSocket的error事件，当发生错误时触发
            this.socket.onerror = (error) => {
                console.error('socket.onerror:', error);
            };
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
                    for (let j = 0; j < this.recording_items[i].content.length; j++) {
                        if (this.recording_items[i].content[j] === this.originText) {
                            this.recording_items[i].content[j] = this.replaceText;
                            i_replace = true;
                        }
                    }

                    if (i_replace && this.translation === '0') {
                        await this.customSleep(500);
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
        },
        changeSearchStatus() {
            this.dictIndexID = {};
            this.dictSpanClass = {};
            this.searchInfo = '';
            this.isSearchContent = true;
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

            if (this.translation === '0') {
                for (let i = 0; i < this.recording_items.length; i++) {
                    let i_trans_content = this.concatContent(this.recording_items[i].transContent);
                    this.recording_items[i].transContent = [i_trans_content];
                }
            }

            this.dictIndexID = {};
            this.dictSpanClass = {};
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

                    // 翻译文本
                    if (this.translation === '0') {
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
        concatContent(curContent) {
            let result_content = '';
            for (let i = 0; i < curContent.length; i++) {
                result_content += curContent[i];
            }
            return result_content;
        },
        computeSearchSpanClass(prefix, one_index, two_index) {
            let span_id = `${prefix}-${one_index}-${two_index}`;
            if (this.dictSpanClass.hasOwnProperty(span_id)) {
                return this.dictSpanClass[span_id];
            } else {
                return '';
            }
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
        uploadVoiceRecord(loadingInstance) {
            this.closeMicrophone();
            let buffer = this.recorder.getBufferAll();
            let f32buffer = this.float32Flatten(buffer);
            let int16buffer = this.FloatArray2Int16(f32buffer);
            let blob = this.encodeMono(1, 48000, int16buffer);

            const formData = new FormData();
            formData.append('file', blob); //文件
            formData.append('filename', this.saveFileName); //文件名
            formData.append('user_id', userstore.user_id);
            formData.append('type', 0);
            formData.append('title', this.recordInfo);
            let text = {
                record_timestamp: this.recordInfo,
                source_language: this.sourceLanguage,
                dest_language: this.destLanguage,
                content: this.content,
                savedTime: this.getCurSavedTime()
            };
            let tmp_recordings = [];
            for (let i = 0; i < this.recording_items.length; ++i) {
                tmp_recordings.push({
                    timestamp: this.recording_items[i].time_stamp,
                    content: this.concatContent(this.recording_items[i].content)
                });
                console.log(this.recording_items[i].time_stamp);
            }
            formData.append('inner_Html', JSON.stringify(text));
            formData.append('content', JSON.stringify(tmp_recordings));

            saveVoice(formData) //调用api
                .then((res) => {
                    if (res.code === 200) {
                        this.$nextTick(() => {
                            // 以服务的方式调用的 Loading 需要异步关闭
                            loadingInstance.close();
                        });
                        this.$router.push('/home');
                    } else {
                        this.$nextTick(() => {
                            // 以服务的方式调用的 Loading 需要异步关闭
                            loadingInstance.close();
                        });
                        Message.error('上传失败');
                    }
                });
        },
        handleSave() {
            this.$confirm('保存并退出？', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'info'
            }).then(() => {
                const loadingInstance = Loading.service({
                    lock: true,
                    text: '正在转写中，请稍等',
                    spinner: 'el-icon-loading',
                    background: 'rgba(0, 0, 0, 0.8)'
                });
                this.sleep(2000)
                this.uploadVoiceRecord(loadingInstance);
            });
        },
        noSaveClose() {
            this.centerDialogVisible = false;
            this.$router.push('/home');
        },
        saveClose() {
            this.centerDialogVisible = false;
            const loadingInstance = Loading.service({
                lock: true,
                text: '正在转写中，请稍等',
                spinner: 'el-icon-loading',
                background: 'rgba(0, 0, 0, 0.8)'
            });
            this.sleep(2000)
            this.uploadVoiceRecord(loadingInstance);
            this.$router.push('/home');
        },
        FloatArray2Int16(floatbuffer) {
            var int16Buffer = new Int16Array(floatbuffer.length);
            for (var i = 0, len = floatbuffer.length; i < len; i++) {
                if (floatbuffer[i] < 0) {
                    int16Buffer[i] = 0x8000 * floatbuffer[i];
                } else {
                    int16Buffer[i] = 0x7fff * floatbuffer[i];
                }
            }
            return int16Buffer;
        },

        encodeMono(channels, sampleRate, samples) {
            var buffer = [];
            var mp3enc = new lamejs.Mp3Encoder(channels, sampleRate, 256);
            var remaining = samples.length;
            var maxSamples = 1152;
            for (var i = 0; remaining >= maxSamples; i += maxSamples) {
                var mono = samples.subarray(i, i + maxSamples);
                var mp3buf = mp3enc.encodeBuffer(mono);
                if (mp3buf.length > 0) {
                    buffer.push(new Int8Array(mp3buf));
                }
                remaining -= maxSamples;
            }
            var d = mp3enc.flush();
            if (d.length > 0) {
                buffer.push(new Int8Array(d));
            }
            var blob = new Blob(buffer, { type: 'audio/mp3' });
            return blob;
        },
        float32Flatten(chunks) {
            //get the total number of frames on the new float32array
            const nFrames = chunks.reduce((acc, elem) => acc + elem.length, 0);

            //create a new float32 with the correct number of frames
            const result = new Float32Array(nFrames);

            //insert each chunk into the new float32array
            let currentFrame = 0;
            chunks.forEach((chunk) => {
                result.set(chunk, currentFrame);
                currentFrame += chunk.length;
            });
            return result;
        },
        sleep(time) {
            var timeStamp = new Date().getTime();
            var endTime = timeStamp + time;
            while (true) {
                if (new Date().getTime() > endTime) {
                    return;
                }
            }
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

.jphwIe {
    width: 65%;
    height: 100%;
    overflow: clip;
    display: flex;
    flex-direction: column;
    position: relative;
    margin: 0px auto;
    padding-left: 28px;
    padding-right: 18px;
}

.WFeMQ {
    display: flex;
    -webkit-box-pack: justify;
    justify-content: space-between;
    top: 0px;
    left: 0px;
    -webkit-box-align: center;
    align-items: flex-end;
    width: 100%;
    height: 96px;
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
    font-size: 15px !important;
    font-weight: 530;
    color: #696969;
    line-height: 24px;
    outline: none;
    box-shadow: none;
}
/*激活样式*/
::v-deep .el-radio-button__orig-radio:checked + .el-radio-button__inner {
    background: #409eff;
    border: 1 !important;
    color: #fdfdfd;
    font-size: 16px !important;
    font-weight: 500;
    line-height: 24px;
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

.dDPXpc {
    height: calc(100vh - 88px - 112px);
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

.transJItVgd {
    display: flex;
    align-items: flex-start;
    width: 100%;
    padding: 16px 24px;
    background-color: rgba(142, 85, 243, 0.45);
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

.custom-rounded {
    border-radius: 8px; /* 自定义圆角大小 */
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
</style>