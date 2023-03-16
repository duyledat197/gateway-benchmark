package main

import (
	"context"
	"sisu-network/gateway/idl/pb"
	"sisu-network/gateway/pkg/configs"
	"sisu-network/gateway/pkg/grpc_client"
	"sisu-network/gateway/pkg/http_server"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type server struct {
	configs *configs.Config

	//* loadConn
	userConn *grpc.ClientConn

	userDialClient *grpc_client.ConnClient

	//* load client
	userClient pb.UserServiceClient

	httpServer *http_server.HttpServer

	//* logger
	logger *zap.Logger

	processors []processor
	factories  []factory
}

type processor interface {
	Init(ctx context.Context) error
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type factory interface {
	Connect(ctx context.Context) error
	Stop(ctx context.Context) error
}
