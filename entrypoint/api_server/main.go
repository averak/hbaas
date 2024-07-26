package main

import (
	"context"

	"github.com/averak/hbaas/app/core/build_info"
	"github.com/averak/hbaas/app/core/config"
	"github.com/averak/hbaas/app/core/logger"
	"github.com/averak/hbaas/app/infrastructure/trace"
)

func main() {
	logger.Init("api-server", build_info.ServerVersion())
	if config.Get().GetGoogleCloud().GetTrace().GetEnabled() {
		trace.Init("user-api", build_info.ServerVersion(), config.Get().GetGoogleCloud().GetTrace().GetSamplingRate())
	}

	ctx := context.Background()
	logger.Notice(ctx, "Start server.")
}
