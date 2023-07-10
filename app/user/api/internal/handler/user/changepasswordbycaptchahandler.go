package user

import (
	"net/http"

	"giligili/app/user/api/internal/logic/user"
	"giligili/app/user/api/internal/svc"
	"giligili/app/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChangePasswordByCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChangePasswordByCaptchaRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewChangePasswordByCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.ChangePasswordByCaptcha(&req)
		if err == nil {
			types.Response(w, resp, 200, "操作成功！")
		} else {
			types.Response(w, err, -1, "操作失败！")
		}
	}
}
