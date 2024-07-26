package handler

import (
	"net/http"

	"connectrpc.com/connect"
	"github.com/averak/hbaas/app/adapter/handler/debug/echo"
	"github.com/averak/hbaas/app/core/config"
	"github.com/averak/hbaas/app/infrastructure/connect/interceptor"
	"github.com/averak/hbaas/protobuf/api/debug/debugconnect"
	"github.com/google/wire"
)

var SuperSet = wire.NewSet(
	echo.New,
	New,
)

func New(
	echo debugconnect.EchoServiceHandler,
) *http.ServeMux {
	opts := connect.WithInterceptors(interceptor.New()...)
	mux := http.NewServeMux()
	if config.Get().GetDebug() {
		mux.Handle(debugconnect.NewEchoServiceHandler(echo, opts))
	}
	return mux
}
