package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/noahlsl/noah/consts"
	"github.com/noahlsl/noah/tools/md5x"
	"github.com/redis/go-redis/v9"
)

// IdempotenceMiddleware 幂等性中间件
type IdempotenceMiddleware struct {
	num int
	r   *redis.ClusterClient
	key string
}

func NewIdempotenceMiddleware(r *redis.ClusterClient, num int) *IdempotenceMiddleware {
	return &IdempotenceMiddleware{
		r:   r,
		num: num,
		key: "base:limit:idempotence:%v",
	}
}

func (m *IdempotenceMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		param := r.Header.Get("param")
		if param != "" && m.num != 0 {
			key := fmt.Sprintf(m.key, md5x.ByString(param, token))
			ok := m.r.SetNX(context.Background(), key, 0,
				time.Duration(m.num)*time.Second).Val()
			if !ok {
				_, _ = w.Write(consts.ErrIdempotence)
				return
			}
		}

		next(w, r)
	}
}

func (m *IdempotenceMiddleware) OriginalHandle(w http.ResponseWriter, r *http.Request) error {

	param := r.Header.Get("param")
	token := r.Header.Get("token")
	if param != "" && m.num != 0 {
		key := fmt.Sprintf(m.key, md5x.ByString(param, token))
		ok := m.r.SetNX(context.Background(), key, 0,
			time.Duration(m.num)*time.Second).Val()
		if !ok {
			_, _ = w.Write(consts.ErrIdempotence)
			return consts.ErrRequestLimit
		}
	}

	return nil
}
