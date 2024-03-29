package user

import (
	"net/http"

	"giligili/app/user/api/internal/logic/user"
	"giligili/app/user/api/internal/svc"
	"giligili/app/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChangePasswordByPasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChangePasswordByPasswordRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewChangePasswordByPasswordLogic(r.Context(), svcCtx)
		resp, err := l.ChangePasswordByPassword(&req)
		if err == nil {
			types.Response(w, resp, 200, "密码修改成功！")
		} else {
			types.Response(w, err, -1, "密码修改失败！")
		}
	}
}
