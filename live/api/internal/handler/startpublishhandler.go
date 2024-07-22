package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zjhx.com/live/api/internal/logic"
	"zjhx.com/live/api/internal/svc"
	"zjhx.com/live/api/internal/types"
	"zjhx.com/pkg/cons"
)

func StartPublishHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OnPublishRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewStartPublishLogic(r.Context(), svcCtx)
		resp, err := l.StartPublish(&req)

		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
			go func(name string) {
				//开启录像
				if resp.IsRecord {
					time.Sleep(5 * time.Second)
					res1, err := http.Get(fmt.Sprintf("%s/control/record/start?app=live&name=%s&rec=all", cons.RequestIp, name))
					fmt.Println("-------------------------", res1, err)
				}
			}(req.Name)
		}
	}
}
