package middleware

import (
	"context"
	"net/http"

	"github.com/noahlsl/noah/tools/idx"
	"github.com/zeromicro/go-zero/core/trace"
)

func TraceMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		traceId := r.Header.Get(trace.TraceIdKey)
		if traceId == "" {
			traceId = idx.GetNanoId()
		}

		//traceId = trace.TraceIDFromContext(ctx)
		r.Header.Set(trace.TraceIdKey, traceId)
		ctx = context.WithValue(ctx, trace.TraceIdKey, traceId)
		r = r.WithContext(ctx)
		next(w, r)
	}
}

func OriginalTraceMiddleware(w http.ResponseWriter, r *http.Request) error {

	ctx := r.Context()
	traceId := r.Header.Get(trace.TraceIdKey)
	if traceId == "" {
		traceId = idx.GetNanoId()
	}

	//traceId = trace.TraceIDFromContext(ctx)
	r.Header.Set(trace.TraceIdKey, traceId)
	ctx = context.WithValue(ctx, trace.TraceIdKey, traceId)
	r = r.WithContext(ctx)

	return nil
}
