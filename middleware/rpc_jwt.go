package middleware

import (
	"context"

	"google.golang.org/grpc"
)

func TraceInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	resp, err = handler(ctx, req)
	if err != nil {

	}

	return resp, err
}
