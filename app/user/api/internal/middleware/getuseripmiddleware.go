package middleware

import (
	"net"
	"net/http"
)

type GetUserIpMiddleware struct {
}

func NewGetUserIpMiddleware() *GetUserIpMiddleware {
	return &GetUserIpMiddleware{}
}

func (m *GetUserIpMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		remoteAddr := r.RemoteAddr
		if ip := r.Header.Get("X-Real-IP"); ip != "" {
			remoteAddr = ip
		} else if ip = r.Header.Get("X-Forwarded-For"); ip != "" {
			remoteAddr = ip
		} else {
			remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
		}

		if remoteAddr == "::1" {
			remoteAddr = "127.0.0.1"
		}

		r.Header.Set("X-Real-IP", remoteAddr)
		next(w, r)

	}
}
