package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zjhx.com/video/api/internal/logic"
	"zjhx.com/video/api/internal/svc"
)

func uploadLargeVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//解析附带参数
		filename := r.FormValue("filename")
		userId := r.FormValue("user_id")
		dataType := r.FormValue("type")
		md5 := r.FormValue("md5")
		current := r.FormValue("current")
		total := r.FormValue("total")

		// 获取上传文件formdata
		file, _, err := r.FormFile("file")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUploadLargeVideoLogic(r.Context(), svcCtx)
		resp, err := l.UploadLargeVideo(filename, userId, dataType, md5, current, total, file)
		file.Close()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
