package handler

import (
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zjhx.com/user/api/internal/logic"
	"zjhx.com/user/api/internal/svc"
	"zjhx.com/user/api/internal/types"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)
		fmt.Println("返回错误？")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			fmt.Println("返回错误？")

		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
