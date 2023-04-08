package middleware

import (
	"net/http"
)

type CorsMiddleware struct {
	AllowOrigin string
}

func CorsDataMiddleware(f string) *CorsMiddleware {
	return &CorsMiddleware{
		AllowOrigin: f,
	}
}

func (m *CorsMiddleware) CorsMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		r.Header.Set("Access-Control-Allow-Origin", "*")
		if m.AllowOrigin != "" {
			r.Header.Set("Access-Control-Allow-Origin", m.AllowOrigin)
		}
		r.Header.Set("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token")
		r.Header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		r.Header.Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		r.Header.Set("Access-Control-Allow-Credentials", "true")
		// 允许放行OPTIONS请求
		if method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next(w, r)
	}
}

func (m *CorsMiddleware) OriginalCorsMiddleware(w http.ResponseWriter, r *http.Request) error {

	method := r.Method
	r.Header.Set("Access-Control-Allow-Origin", "*")
	if m.AllowOrigin != "" {
		r.Header.Set("Access-Control-Allow-Origin", m.AllowOrigin)
	}
	r.Header.Set("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token")
	r.Header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	r.Header.Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	r.Header.Set("Access-Control-Allow-Credentials", "true")
	// 允许放行OPTIONS请求
	if method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}

	return nil
}
