package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zjhx.com/record/api/internal/logic"
	"zjhx.com/record/api/internal/svc"
	"zjhx.com/record/api/internal/types"
)

func recentRecordByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecordByIdRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRecentRecordByIdLogic(r.Context(), svcCtx)
		resp, err := l.RecentRecordById(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
