package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/noahlsl/noah/consts"
)

type TimeoutMiddleware struct {
	timeout time.Duration
}

func NewTimeoutMiddleware(n int) *TimeoutMiddleware {

	return &TimeoutMiddleware{
		timeout: time.Duration(n) * time.Second,
	}
}

func (m *TimeoutMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), m.timeout)
		defer func() {
			if ctx.Err() == context.DeadlineExceeded {
				rs := res.NewRes().WithCode(consts.ErrCodeTimeout)
				_, _ = w.Write(rs.ToBytes())
				return
			}

			cancel()
		}()
		r = r.WithContext(ctx)
		next(w, r)
	}
}
func (m *TimeoutMiddleware) OriginalHandle(w http.ResponseWriter, r *http.Request) error {

	ctx, cancel := context.WithTimeout(r.Context(), m.timeout)
	defer func() {
		if ctx.Err() == context.DeadlineExceeded {
			rs := res.NewRes().WithCode(consts.ErrCodeTimeout)
			_, _ = w.Write(rs.ToBytes())
			return
		}

		cancel()
	}()
	r = r.WithContext(ctx)

	return nil
}
