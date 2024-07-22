package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zjhx.com/video/api/internal/logic"
	"zjhx.com/video/api/internal/svc"
)

func uploadVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// var req types.UploadVideoRequest
		//解析附带参数
		filename := r.FormValue("filename")
		userId := r.FormValue("user_id")
		dataType := r.FormValue("type")
		md5 := r.FormValue("md5")
		sourceLangeage := r.FormValue("source_language")
		destLanguage := r.FormValue("dest_language")
		// 获取上传文件formdata
		file, _, err := r.FormFile("file")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUploadVideoLogic(r.Context(), svcCtx)
		//调用rpc业务创建数据库表
		resp, err := l.UploadVideo(filename, userId, dataType, md5, sourceLangeage, destLanguage, file)
		file.Close()
		//返回信息
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
