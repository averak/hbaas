package main

import (
	"context"

	"github.com/averak/hbaas/app/core/build_info"
	"github.com/averak/hbaas/app/core/logger"
)

func main() {
	logger.Init("api-server", build_info.ServerVersion())

	ctx := context.Background()
	logger.Notice(ctx, "Start server.")
}
