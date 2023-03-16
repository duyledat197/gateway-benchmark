package main

import (
	"context"
)

func main() {
	ctx := context.Background()
	if err := gracefulShutdown(ctx, start); err != nil {
		srv.logger.Sugar().Fatalln(err)
	}
}
