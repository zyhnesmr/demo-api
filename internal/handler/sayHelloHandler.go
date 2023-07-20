package handler

import (
	"net/http"

	"demo-api/internal/logic"
	"demo-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SayHelloHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewSayHelloLogic(r.Context(), svcCtx)
		resp, err := l.SayHello()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
