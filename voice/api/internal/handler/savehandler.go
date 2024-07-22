package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zjhx.com/voice/api/internal/logic"
	"zjhx.com/voice/api/internal/svc"
)

func SaveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filename := r.FormValue("filename")
		userId := r.FormValue("user_id")
		dataType := r.FormValue("type")
		sourceContent := r.FormValue("content")
		innerHtml := r.FormValue("inner_Html")
		title := r.FormValue("title")
		// 获取上传文件formdata
		file, _, err := r.FormFile("file")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		l := logic.NewSaveLogic(r.Context(), svcCtx)
		resp, err := l.Save(filename, userId, dataType, sourceContent, innerHtml, title, file)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
