package base

import (
	"net/http"

	"giligili/app/user/api/internal/logic/base"
	"giligili/app/user/api/internal/svc"
	"giligili/app/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := base.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err == nil {
			types.Response(w, resp, 200, "注册成功！")
		} else {
			types.Response(w, err.Error(), -1, "注册失败！")
		}
	}
}
