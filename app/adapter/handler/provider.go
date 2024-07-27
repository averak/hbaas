package handler

import (
	"net/http"

	"connectrpc.com/connect"
	"github.com/averak/hbaas/app/adapter/handler/debug/echo"
	"github.com/averak/hbaas/app/adapter/handler/session"
	"github.com/averak/hbaas/app/core/config"
	"github.com/averak/hbaas/app/infrastructure/connect/interceptor"
	"github.com/averak/hbaas/protobuf/api/apiconnect"
	"github.com/averak/hbaas/protobuf/api/debug/debugconnect"
	"github.com/google/wire"
)

var SuperSet = wire.NewSet(
	session.New,
	echo.New,
	New,
)

func New(
	session apiconnect.SessionServiceHandler,
	echo debugconnect.EchoServiceHandler,
) *http.ServeMux {
	opts := connect.WithInterceptors(interceptor.New()...)
	mux := http.NewServeMux()
	mux.Handle(apiconnect.NewSessionServiceHandler(session, opts))
	if config.Get().GetDebug() {
		mux.Handle(debugconnect.NewEchoServiceHandler(echo, opts))
	}
	return mux
}
