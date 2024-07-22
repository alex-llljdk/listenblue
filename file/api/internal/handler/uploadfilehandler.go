package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zjhx.com/file/api/internal/logic"
	"zjhx.com/file/api/internal/svc"
)

func uploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// var req types.UploadFileRequest
		//解析附带参数
		// fmt.Println("filename:   " + filename)
		userId := r.FormValue("user_id")
		dataType := r.FormValue("type")
		// 获取上传文件formdata
		if dataType != "5" {
			hash := r.FormValue("md5")
			filename := r.FormValue("name")
			file, _, err := r.FormFile("file")
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, err)
				return
			}
			//内部调用rpc对数据库进行操作
			l := logic.NewUploadFileLogic(r.Context(), svcCtx)
			resp, err := l.UploadFile(filename, userId, dataType, hash, file)
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, err)
			} else {
				httpx.OkJsonCtx(r.Context(), w, resp)
			}
		} else {
			//内部调用rpc对数据库进行操作
			url := r.FormValue("url")
			l := logic.NewUploadFileLogic(r.Context(), svcCtx)
			resp, err := l.UploadHtml(url, userId, dataType)
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, err)
			} else {
				httpx.OkJsonCtx(r.Context(), w, resp)
			}
		}

	}
}
