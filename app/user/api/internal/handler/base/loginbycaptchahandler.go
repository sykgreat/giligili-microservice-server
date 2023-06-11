package base

import (
	"context"
	"net/http"

	"giligili/app/user/api/internal/logic/base"
	"giligili/app/user/api/internal/svc"
	"giligili/app/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginByCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginByCaptchaRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		get := r.Header.Get("X-Real-IP")
		ctx := context.WithValue(r.Context(), "X-Real-IP", get)

		l := base.NewLoginByCaptchaLogic(ctx, svcCtx)
		resp, err := l.LoginByCaptcha(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
