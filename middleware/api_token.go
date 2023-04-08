package middleware

import (
	"context"
	"fmt"
	"github.com/noahlsl/noah/consts"
	"github.com/zeromicro/go-zero/core/trace"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type TokenMiddleware struct {
	key  string
	r    *redis.ClusterClient
	flag []any
}

func NewTokenMiddleware(r *redis.ClusterClient, key string, flag ...any) *TokenMiddleware {
	return &TokenMiddleware{
		key:  key,
		r:    r,
		flag: flag,
	}
}

func (m *TokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")
		if token == "" {
			token = r.Header.Get("token")
		}

		if token == "" {
			token = r.Form.Get("token")
		}

		if token == "" {
			_, _ = w.Write(consts.ErrTokenExpired)
			return
		}

		var flag []any
		for _, f := range m.flag {
			if f == nil {
				continue
			}
			flag = append(flag, f)
		}
		flag = append(flag, token)
		key := fmt.Sprintf(m.key, flag...)
		var (
			err    error
			result string
		)
		for i := 0; i <= 3; i++ {
			result, err = m.r.Get(r.Context(), key).Result()
			if err != nil {
				if i == 3 {
					_, _ = w.Write(consts.ErrTokenExpired)
					return
				}

				continue
			}
			break
		}
		// 设置 trace
		r.Header.Set("trace", trace.TraceIDFromContext(r.Context()))
		// 客户端请求使用的用户(user)ID
		r.Header.Set("uid", result)
		// 后台请求时候使用的管理员(administrator)ID
		r.Header.Set("aid", result)
		// 回写token
		r.Header.Set("token", token)
		next(w, r)
	}
}

func (m *TokenMiddleware) OriginalHandle(w http.ResponseWriter, r *http.Request) error {

	var ctx = context.Background()
	token := r.Header.Get("Authorization")
	if token == "" {
		token = r.Header.Get("token")
	}

	if token == "" {
		token = r.URL.Query().Get("token")
	}

	if token == "" {
		token = r.Form.Get("token")
	}

	if token == "" {
		return consts.ErrSysTokenExpired
	}

	var flag []any
	for _, f := range m.flag {
		if f == nil {
			continue
		}
		flag = append(flag, f)
	}
	flag = append(flag, token)
	key := fmt.Sprintf(m.key, flag...)
	var (
		err    error
		result string
	)
	for i := 0; i <= 3; i++ {
		result, err = m.r.Get(ctx, key).Result()
		if err != nil {
			if i == 3 {
				return consts.ErrSysTokenExpired
			}
			continue
		}
		break
	}

	// 设置 trace
	r.Header.Set("trace", trace.TraceIDFromContext(r.Context()))
	// 客户端请求使用的用户(user)ID
	r.Header.Set("uid", result)
	// 后台请求时候使用的管理员(administrator)ID
	r.Header.Set("aid", result)
	// 回写token
	r.Header.Set("token", token)

	return nil
}
