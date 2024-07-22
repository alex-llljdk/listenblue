package handler

import (
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zjhx.com/user/api/internal/logic"
	"zjhx.com/user/api/internal/svc"
)

func UploadImgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file_name := r.FormValue("file_name")
		file, _, err := r.FormFile("file")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		fmt.Println(file)
		l := logic.NewUploadImgLogic(r.Context(), svcCtx)
		resp, err := l.UploadImg(file_name, file)
		file.Close()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)

		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
