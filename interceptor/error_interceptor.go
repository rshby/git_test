package interceptor

import (
	"context"
	"git_test/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func WithErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			// set metadata
			md := metadata.Pairs("error", err.Error())
			_ = grpc.SetHeader(ctx, md)

			code, ok := mapErrorCodes[err]
			if !ok {
				return nil, status.Error(codes.Internal, "internal server error")
			}

			return nil, status.Error(code, err.Error())
		}
		return
	}
}

var mapErrorCodes = map[error]codes.Code{
	service.ErrCustomerNotFound: codes.NotFound,
}
