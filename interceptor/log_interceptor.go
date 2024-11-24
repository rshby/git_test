package interceptor

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func WithLogInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		logrus.Infof("melewati interceptor log")
		return handler(ctx, req)
	}
}
