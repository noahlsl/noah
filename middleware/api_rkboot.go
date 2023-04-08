package middleware

import (
	"net/http"

	"github.com/rookie-ninja/rk-boot"
	_ "github.com/rookie-ninja/rk-gin/boot"
)

type RkBootMiddleware struct {
	boot *rkboot.Boot
}

func NewRkBootMiddleware() *RkBootMiddleware {
	return &RkBootMiddleware{
		boot: rkboot.NewBoot(),
	}
}

// Handle 监控中间件
func (m *RkBootMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m.boot.Bootstrap(r.Context())
		next(w, r)
		m.boot.WaitForShutdownSig(r.Context())
	}
}
