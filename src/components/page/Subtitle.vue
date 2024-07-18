<template>
    <div class="h-full w-full flex items-center justify-center">
        <div class="text-container">
            <span class="text-content break-words text-blue-400">
                <!-- {{ showSubtitle }} -->
                今天我很开新撒飞洒飞洒飞洒飞洒飞洒飞洒飞洒飞洒发生的撒大大撒今天我很开新撒飞洒飞洒飞洒飞洒飞洒飞洒飞洒飞洒发生的撒大大撒今天我很开新撒飞洒飞洒ddd飞洒飞洒飞大大萨达萨达萨达萨达洒飞洒飞洒飞洒发生的撒大大撒ddddddddddddddd大撒大撒ddd
            </span>
        </div>
    </div>

    <!-- <div class="bg-transparent text-blue-400 text-7xl leading-loose"></div> -->
</template>

<script>
import { LiveSubtitleUrl } from '../../api/live';
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
export default {
    destroy() {
        if (this.interval) {
            clearInterval(this.interval);
        }
        this.socket.close()
    },
    mounted() {
        this.startRecording();
    },
    data() {
        return {
            mediaStream: null,
            recorder: null,
            url: null,
            showSubtitle: ''
        };
    },
    methods: {
        goBack() {
            this.$router.go(-1);
        },
        //启动录音
        startRecording() {
            navigator.getUserMedia = navigator.getUserMedia || navigator.webkitGetUserMedia;
            if (!navigator.getUserMedia) {
                console.error('无法访问麦克风');
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
            this.url = LiveSubtitleUrl;
            this.socket = new WebSocket(this.url);

            // 监听WebSocket的open事件，当连接打开时触发
            this.socket.onopen = (event) => {
                console.log('socket.onopen成功');
                this.recorder = new Recorder(mediaStream, this.socket);
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
                        this.recorder.start();
                        break;
                    case 'result':
                        if (data.data.var) {
                            this.showSubtitle = data.data.var;
                        } else {
                            this.showSubtitle = data.data.onebest;
                        }

                        break;
                }
            };
            // 监听WebSocket的close事件，当连接关闭时触发
            this.socket.onclose = (event) => {
                console.log('socket.onclose==');
                console.log(`WebSocket连接${this.second}秒后关闭了`);
                console.log('WebSocket is closed. Code: ' + event.code + ' Reason: ' + event.wasClean);
                if (this.interval) {
                    clearInterval(this.interval);
                }
                //连接关闭，就重新连接
                // this.reconnect();
            };

            // 监听WebSocket的error事件，当发生错误时触发
            this.socket.onerror = (error) => {
                console.error('socket.onerror:', error);
            };

            this.interval = setInterval(() => {
                this.socket.send(
                    "ping"
                );
            },  60*1000);
        }
    }
};
</script>


<style scoped>
.text-container {
    position: relative;
    width: 1200px; /* 定义宽度 */
    height: 60px;
    border: 1px solid #000; /* 边框用于可视化 div 边界 */
}

.text-container:before {
    content: '';
    display: inline-block;
    vertical-align: middle;
    height: 100%;
}

.text-content {
    display: inline-block;
    vertical-align: middle;
    position: absolute;
    bottom: 0;
    left: 0;
    right: 0;
    text-align: center; /* 确保文本水平居中 */
    margin: 0 auto; /* 可用于调整文本的左右边距，如果需要 */
    font-size: 40px;
}
</style>
