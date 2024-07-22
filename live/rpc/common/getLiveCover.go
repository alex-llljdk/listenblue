package common

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"zjhx.com/live/code"
	"zjhx.com/live/gmodel"
	"zjhx.com/live/rpc/internal/svc"
	"zjhx.com/pkg/cons"
)

func GetLiveCover(svcCtx *svc.ServiceContext) {
	ticker := time.NewTicker(60 * time.Second) // 创建一个每秒执行一次的定时器
	defer ticker.Stop()                        // 确保在不再需要时停止定时器

	for {
		select {
		case <-ticker.C: // 每当定时器触发时
			func() {
				roomList, err := svcCtx.BizRedis.Lrange("RoomList", 0, -1)
				if err != nil {
					logx.Error(err)
					return
				}
				if len(roomList) == 0 {
					return
				}
				for _, v := range roomList {
					res, err := svcCtx.BizRedis.Get(cons.LiveKey + v)
					if err != nil {
						logx.Error(err)
						return
					}
					var roomInfo gmodel.LiveRoom
					err = json.Unmarshal([]byte(res), &roomInfo)
					if err != nil {
						logx.Error(err)
						return
					}
					if !roomInfo.IsOpen {
						continue
					}

					spfilename := strings.Split(roomInfo.CoverUrl, "/")

					ffmpegCmd := exec.Command("ffmpeg",
						"-i", roomInfo.Url, // 输入流
						"-frames:v", "1", // 截取视频的第一帧
						"-s", "200x100",
						"-f", "image2", // 输出格式
						"-y", code.LiveDataUrl+"/"+spfilename[len(spfilename)-1]) // 输出文件

					// 执行FFmpeg命令
					err = ffmpegCmd.Run()
					if err != nil {
						fmt.Printf("Error: %s\n", err)
					} else {
						fmt.Println("Screenshot captured successfully.")
					}
				}

			}() // 调用需要定期执行的函数
		}
	}
}
