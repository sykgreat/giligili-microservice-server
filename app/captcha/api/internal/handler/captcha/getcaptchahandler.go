package captcha

import (
	"net/http"

	"giligili/app/captcha/api/internal/logic/captcha"
	"giligili/app/captcha/api/internal/svc"
	"giligili/app/captcha/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCaptchaRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := captcha.NewGetCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.GetCaptcha(&req)
		if err == nil {
			types.Response(w, resp, 200, "获取验证码成功！")
		} else {
			types.Response(w, err, -1, "获取验证码失败！")
		}
	}
}
