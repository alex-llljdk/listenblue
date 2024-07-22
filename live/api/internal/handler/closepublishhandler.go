package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zjhx.com/live/api/chat"
	"zjhx.com/live/api/internal/logic"
	"zjhx.com/live/api/internal/svc"
	"zjhx.com/live/api/internal/types"
)

func ClosePublishHandler(svcCtx *svc.ServiceContext, hub *chat.Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ClosePublishRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewClosePublishLogic(r.Context(), svcCtx)
		resp, err := l.ClosePublish(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			hub.SendEnd()
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
