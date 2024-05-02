<template>
    <div class="h-full w-full bg-gray-100 flex">
        <div class="h-full bg-indigo-50">
            <el-aside width="80px">
                <router-link to="/home">
                    <div class="text-center mt-12">
                        <div class="rounded-lg bg-white mx-3 mt-7" style="font-size: 32px;">
                            <i class="el-icon-arrow-left"></i>
                        </div>
                    </div>
                </router-link>
                <div class="text-center mt-8">
                    <div>
                        <i class="el-icon-folder-checked text-4xl"></i>
                    </div>
                    <div class="text-base">保存</div>
                </div>
            </el-aside>
        </div>
        <div class="h-full pl-7 mt-12 w-1/2">
            <div class="hover_input focus_input">
                <input
                    type="text"
                    class="hip fip -ml-2 pl-2 font-semibold text-lg bg-transparent border-2 border-gray-100 focus:border-2 focus:outline-none rounded-md max-w-lg"
                    v-model="recordInfo"
                    :style="{ width: inputWidth + 'px' }"
                    @input="handleInput"
                />
            </div>

            <div>
                <span class="text-xs text-gray-400" style="user-select: none;">{{ savedTime }}</span>
            </div>

            <div class="ixtgUW">
                    <div class="w-1/2">
                        <img src="../../assets/img/voice_record.jpg" />
                    </div>

                    <!-- <div>
                        <div>音频语言</div>
                    </div> -->
            </div>
        </div>
        <div class="w-1/2"></div>
    </div>
</template>

<script>
export default {
    name: 'VoiceRecord',
    data() {
        return {
            inputWidth: 204,
            recordInfo: '',
            savedTime: ''
        };
    },
    mounted() {
        document.title = 'ListenBlue-会议记录';
        this.updateInfo();
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
</style>