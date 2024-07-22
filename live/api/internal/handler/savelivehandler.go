package handler

import (
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zjhx.com/live/api/internal/logic"
	"zjhx.com/live/api/internal/svc"
	"zjhx.com/live/api/internal/types"
)

func SaveLiveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SaveRequest
		fmt.Println(111)
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		fmt.Println(2222)
		l := logic.NewSaveLiveLogic(r.Context(), svcCtx)
		resp, err := l.SaveLive(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
