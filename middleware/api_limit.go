package middleware

import (
	"net/http"
	"time"

	"github.com/juju/ratelimit"
	"github.com/noahlsl/noah/consts"
)

type LimitMiddleware struct {
	bucket *ratelimit.Bucket
}

// NewLimitMiddleware  限流中间件
// fillInterval 时间段
// cap 容量
// quantum 生产速度
func NewLimitMiddleware(fillInterval time.Duration, cap, quantum int64) *LimitMiddleware {
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, cap, quantum)
	return &LimitMiddleware{
		bucket: bucket,
	}
}

func (l *LimitMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if l.bucket.TakeAvailable(1) < 1 {
			rs := res.NewRes().WithCode(consts.ErrCodeSysBadRequest)
			_, _ = w.Write(rs.ToBytes())
			return
		}

		next(w, r)
	}
}

func (l *LimitMiddleware) OriginalHandle(w http.ResponseWriter, r *http.Request) error {

	if l.bucket.TakeAvailable(1) < 1 {
		return consts.ErrRequestLimit
	}

	return nil
}
