package grpc_client

import (
	"context"

	"sisu-network/gateway/pkg/hystrix_config"

	"github.com/afex/hystrix-go/hystrix"
	"google.golang.org/grpc"
)

// UnaryClientInterceptor ...
func UnaryClientInterceptor(isEnableHystrix bool) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if isEnableHystrix {
			hystrix.ConfigureCommand(method, hystrix_config.HystrixConfig())

			err := hystrix.Do(method, func() (err error) {
				err = invoker(ctx, method, req, reply, cc, opts...)
				return err
			}, nil)

			return err
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
