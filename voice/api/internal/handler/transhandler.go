package handler

import (
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zjhx.com/voice/api/internal/logic"
	"zjhx.com/voice/api/internal/svc"
	"zjhx.com/voice/api/internal/types"
)

func TransHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TransRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		fmt.Println("解析成功")
		l := logic.NewTransLogic(r.Context(), svcCtx)
		resp, err := l.Trans(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
