package main

import (
	"context"
	"sisu-network/gateway/idl/pb"
	"sisu-network/gateway/pkg/grpc_client"
	"sisu-network/gateway/pkg/http_server"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func (s *server) loadLogger() error {
	// s.logger = logger.NewZapLogger("INFO", true)
	return nil
}

func (s *server) loadConfig(ctx context.Context) error {
	return nil
}

func (s *server) loadClients(ctx context.Context) error {
	//* load grpc clients

	defaultOptions := &grpc_client.Options{
		IsEnableHystrix:            false,
		IsEnableClientLoadBalancer: false,
		IsEnableTracing:            false,
		IsEnableRetry:              true,
		IsEnableMetrics:            false,
		IsEnableSecure:             false,
		IsEnableValidator:          true,
	}

	s.userDialClient = &grpc_client.ConnClient{
		ServiceName: "User",
		Options:     defaultOptions,
	}
	s.userClient = pb.NewUserServiceClient(s.userDialClient.Conn)

	s.factories = append(s.factories, s.userDialClient)

	return nil
}

func (s *server) loadServers(ctx context.Context) error {
	s.httpServer = &http_server.HttpServer{
		Address: s.configs.HTTP,
		Logger:  s.logger,
		Handlers: func(ctx context.Context, mux *runtime.ServeMux) error {
			if err := pb.RegisterUserServiceHandlerClient(ctx, mux, s.userClient); err != nil {
				return err
			}

			return nil
		},
		Options: &http_server.Options{},
	}

	s.processors = append(s.processors, s.httpServer)

	return nil
}
