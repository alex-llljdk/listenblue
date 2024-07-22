package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zjhx.com/video/api/internal/logic"
	"zjhx.com/video/api/internal/svc"
	"zjhx.com/video/api/internal/types"
)

func mergeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//解析附带参数
		var req types.MergeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewMergeLogic(r.Context(), svcCtx)
		resp, err := l.Merge(req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
