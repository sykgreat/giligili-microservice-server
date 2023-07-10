package middleware

import (
	"context"
	"giligili/common/enum"
	"giligili/common/jwt"
	"github.com/zeromicro/go-zero/core/stores/redis"
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
		token, claims, err := m.Jwt.ParseToken(tokenString)
		if err != nil || !token.Valid {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// 验证token存在 -> 判断token类型
		if claims.TokenType == 0 { // accessToken
			// 读取缓存
			accessToken, err := m.Redis.Get(enum.UserModule + ":" + enum.Token + ":" + strconv.Itoa(int(claims.UserId)) + ":" + enum.AccessToken)
			if err != nil { // accessToken 不存在
				http.Error(w, err.Error(), http.StatusUnauthorized)
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
				http.Error(w, "token验证失败", http.StatusUnauthorized)
				return
			}
		} else if claims.TokenType == 1 { // refreshToken
			// 读取缓存
			refreshToken, err := m.Redis.Get(enum.UserModule + ":" + enum.Token + ":" + strconv.Itoa(int(claims.UserId)) + ":" + enum.RefreshToken)
			if err != nil { // refreshToken 不存在
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			if refreshToken != "" { // refreshToken 存在
				// 刷新accessToken
				accessToken, err := m.Redis.Get(enum.UserModule + ":" + enum.Token + ":" + strconv.Itoa(int(claims.UserId)) + ":" + enum.AccessToken)
				if err != nil || accessToken == "" { // accessToken 不存在
					accessToken, err = m.Jwt.GenerateAccessToken(claims.UserId, claims.Email)
					if err != nil {
						http.Error(w, err.Error(), http.StatusUnauthorized)
						return
					}
					if err := m.Redis.SetexCtx(
						r.Context(),
						enum.UserModule+":"+enum.Token+":"+strconv.Itoa(int(claims.UserId))+":"+enum.AccessToken,
						accessToken,
						168,
					); err != nil {
						http.Error(w, err.Error(), http.StatusUnauthorized)
						return
					}
				}

				// 将新的accessToken放入body中返回 -> 用于前端刷新
				_, err = w.Write([]byte(accessToken))
				if err != nil {
					http.Error(w, err.Error(), http.StatusUnauthorized)
					return
				}
				return
			} else {
				http.Error(w, "token验证失败", http.StatusUnauthorized)
				return
			}
		} else {
			http.Error(w, "token验证失败", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
