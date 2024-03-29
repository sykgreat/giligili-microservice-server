package middleware

import (
	"giligili/common/enum"
	"giligili/common/jwt"

	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strconv"
)

type JwtAuthMiddleware struct {
	Jwt   *jwt.Jwt
	Redis *redis.Redis
}

func NewJwtAuthMiddleware(jwt *jwt.Jwt, redis *redis.Redis) *JwtAuthMiddleware {
	return &JwtAuthMiddleware{
		Jwt:   jwt,
		Redis: redis,
	}
}

func (m *JwtAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		_, claims, err := m.Jwt.ParseToken(tokenString)
		if err != nil {
			httpx.OkJson(w, struct {
				Code int         `json:"code"`
				Msg  string      `json:"msg"`
				Data interface{} `json:"data,omitempty"`
			}{
				Code: http.StatusUnauthorized,
				Msg:  "token验证失败！",
				Data: err.Error(),
			})
			return
		}
		// 验证token存在 -> 判断token类型
		if claims.TokenType == 0 { // accessToken
			// 读取缓存
			accessToken, err := m.Redis.Get(enum.UserModule + enum.Token + strconv.Itoa(int(claims.UserId)) + ":" + enum.AccessToken)
			if err != nil { // accessToken 不存在
				httpx.OkJson(w, struct {
					Code int         `json:"code"`
					Msg  string      `json:"msg"`
					Data interface{} `json:"data,omitempty"`
				}{
					Code: http.StatusUnauthorized,
					Msg:  "token验证失败！",
					Data: err.Error(),
				})
				return
			}
			if accessToken == tokenString { // accessToken 未过期
				// 验证权限
				//user := service.SelectUserByID(claims.UserId)
				//role := dto.GetRoleString(user.Role)
				//if !authentication.Check(role, ctx.FullPath(), ctx.Request.Method) {
				//	http.Error(w, "权限不足", http.StatusUnauthorized)
				//	return
				//}
				r = r.WithContext(context.WithValue(r.Context(), "userId", claims.UserId))
				next(w, r)
			} else {
				// 刷新accessToken 和 refreshToken
				httpx.OkJson(w, struct {
					Code int         `json:"code"`
					Msg  string      `json:"msg"`
					Data interface{} `json:"data,omitempty"`
				}{
					Code: http.StatusUnauthorized,
					Msg:  "token验证失败！",
					Data: err.Error(),
				})
				return
			}
		} else if claims.TokenType == 1 { // refreshToken
			// 读取缓存
			refreshToken, err := m.Redis.Get(enum.UserModule + enum.Token + strconv.Itoa(int(claims.UserId)) + ":" + enum.RefreshToken)
			if err != nil { // refreshToken 不存在
				httpx.OkJson(w, struct {
					Code int         `json:"code"`
					Msg  string      `json:"msg"`
					Data interface{} `json:"data,omitempty"`
				}{
					Code: http.StatusUnauthorized,
					Msg:  "token验证失败！",
					Data: err.Error(),
				})
				return
			}
			if refreshToken != "" { // refreshToken 存在
				// 刷新accessToken
				accessToken, err := m.Redis.Get(enum.UserModule + enum.Token + strconv.Itoa(int(claims.UserId)) + ":" + enum.AccessToken)
				if err != nil || accessToken == "" { // accessToken 不存在
					accessToken, err = m.Jwt.GenerateAccessToken(claims.UserId, claims.Email)
					if err != nil {
						httpx.OkJson(w, struct {
							Code int         `json:"code"`
							Msg  string      `json:"msg"`
							Data interface{} `json:"data,omitempty"`
						}{
							Code: http.StatusUnauthorized,
							Msg:  "token验证失败！",
							Data: err.Error(),
						})
						return
					}
					if err := m.Redis.SetexCtx(
						r.Context(),
						enum.UserModule+":"+enum.Token+":"+strconv.Itoa(int(claims.UserId))+":"+enum.AccessToken,
						accessToken,
						7*24*60*60,
					); err != nil {
						httpx.OkJson(w, struct {
							Code int         `json:"code"`
							Msg  string      `json:"msg"`
							Data interface{} `json:"data,omitempty"`
						}{
							Code: http.StatusUnauthorized,
							Msg:  "token验证失败！",
							Data: err.Error(),
						})
						return
					}
				}

				// 将新的accessToken放入body中返回 -> 用于前端刷新
				httpx.OkJson(w, struct {
					Code int         `json:"code"`
					Msg  string      `json:"msg"`
					Data interface{} `json:"data,omitempty"`
				}{
					Code: http.StatusOK,
					Msg:  "token刷新成功！",
					Data: accessToken,
				})
				return
			} else {
				httpx.OkJson(w, struct {
					Code int         `json:"code"`
					Msg  string      `json:"msg"`
					Data interface{} `json:"data,omitempty"`
				}{
					Code: http.StatusUnauthorized,
					Msg:  "token验证失败！",
					Data: err.Error(),
				})
				return
			}
		} else {
			httpx.OkJson(w, struct {
				Code int         `json:"code"`
				Msg  string      `json:"msg"`
				Data interface{} `json:"data,omitempty"`
			}{
				Code: http.StatusUnauthorized,
				Msg:  "token验证失败！",
				Data: err.Error(),
			})
			return
		}
	}
}
