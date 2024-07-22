package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zjhx.com/live/api/internal/logic"
	"zjhx.com/live/api/internal/svc"
	"zjhx.com/live/api/internal/types"
	"zjhx.com/pkg/xcode"
)

func JoinRoomHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.JoinRoomRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewJoinRoomLogic(r.Context(), svcCtx)
		resp, err := l.JoinRoom(&req)
		if err != nil {
			_, er := xcode.ErrHandler(err)
			httpx.OkJsonCtx(r.Context(), w, er)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
