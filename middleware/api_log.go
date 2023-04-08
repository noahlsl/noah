package middleware

import (
	"github.com/noahlsl/noah/tools/logx"
	"net/http"
)

type LogMiddleware struct {
}

func NewLogMiddleware() *LogMiddleware {
	return &LogMiddleware{}
}

func (m *LogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		logx.Debugw("request", logx.GetFields(r)...)
		next(w, r)
	}
}
func (m *LogMiddleware) OriginalHandle(w http.ResponseWriter, r *http.Request) error {

	logx.Debugw("request", logx.GetFields(r)...)
	return nil
}
