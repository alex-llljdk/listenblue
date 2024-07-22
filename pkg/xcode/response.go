package xcode

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zjhx.com/pkg/xcode/types"
)

func ErrHandler(err error) (int, any) {
	code := CodeFromError(err)

	return http.StatusOK, types.Status{
		Code:    int32(code.Code()),
		Message: code.Message(),
	}
}

func JwtUnauthorizedResult(w http.ResponseWriter, r *http.Request, err error) {
	httpx.WriteJson(w, http.StatusOK, &types.Status{
		Code:    int32(401),
		Message: "鉴权失败，请重新登录",
	})
}
