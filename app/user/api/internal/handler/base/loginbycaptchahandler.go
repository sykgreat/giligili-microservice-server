package base

import (
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

		l := base.NewLoginByCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.LoginByCaptcha(&req)
		if err == nil {
			types.Response(w, resp, 200, "操作成功！")
		} else {
			types.Response(w, err, -1, "操作失败！")
		}
	}
}
