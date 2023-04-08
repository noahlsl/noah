package middleware

import (
	"context"
	"net/http"

	"github.com/noahlsl/noah/consts"
	"github.com/noahlsl/noah/tools/ipx"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type IPMiddleware struct {
	r   *redis.ClusterClient
	key string
}

func NewIPMiddleware(r *redis.ClusterClient, key string) *IPMiddleware {
	return &IPMiddleware{
		r:   r,
		key: key,
	}
}

func (l *IPMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ipStr := ipx.GetClientIp()
		result, err := l.r.SIsMember(context.Background(), l.key, ipStr).Result()
		if err != nil {
			rs := res.NewRes().WithCode(consts.ErrCodeRequestLimit)
			_, _ = w.Write(rs.ToBytes())
			return
		}

		if !result {
			rs := res.NewRes().WithCode(consts.ErrCodeRequestLimit)
			_, _ = w.Write(rs.ToBytes())
			return
		}

		next(w, r)
	}
}

func (l *IPMiddleware) OriginalHandle(w http.ResponseWriter, r *http.Request) error {

	ipStr := ipx.GetClientIp()
	result, err := l.r.SIsMember(context.Background(), l.key, ipStr).Result()
	if err != nil {
		return errors.WithMessage(consts.ErrRequestLimit, err.Error())
	}

	if !result {
		return consts.ErrRequestLimit
	}

	return nil
}
