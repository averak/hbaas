package handler

import (
	"net/http"

	"connectrpc.com/connect"
	"github.com/averak/hbaas/app/adapter/handler/debug/echo"
	"github.com/averak/hbaas/app/adapter/handler/global_kvs"
	"github.com/averak/hbaas/app/adapter/handler/leader_board"
	"github.com/averak/hbaas/app/adapter/handler/private_kvs"
	"github.com/averak/hbaas/app/adapter/handler/session"
	"github.com/averak/hbaas/app/adapter/handler/user"
	"github.com/averak/hbaas/app/core/config"
	"github.com/averak/hbaas/app/infrastructure/connect/interceptor"
	"github.com/averak/hbaas/protobuf/api/apiconnect"
	"github.com/averak/hbaas/protobuf/api/debug/debugconnect"
	"github.com/google/wire"
)

var SuperSet = wire.NewSet(
	global_kvs.NewHandler,
	leader_board.NewHandler,
	private_kvs.NewHandler,
	session.NewHandler,
	user.NewHandler,
	echo.NewHandler,
	New,
)

func New(
	globalKVS apiconnect.GlobalKVSServiceHandler,
	leaderBoard apiconnect.LeaderBoardServiceHandler,
	privateKVS apiconnect.PrivateKVSServiceHandler,
	session apiconnect.SessionServiceHandler,
	user apiconnect.UserServiceHandler,
	echo debugconnect.EchoServiceHandler,
) *http.ServeMux {
	opts := connect.WithInterceptors(interceptor.New()...)
	mux := http.NewServeMux()
	mux.Handle(apiconnect.NewGlobalKVSServiceHandler(globalKVS, opts))
	mux.Handle(apiconnect.NewLeaderBoardServiceHandler(leaderBoard, opts))
	mux.Handle(apiconnect.NewPrivateKVSServiceHandler(privateKVS, opts))
	mux.Handle(apiconnect.NewSessionServiceHandler(session, opts))
	mux.Handle(apiconnect.NewUserServiceHandler(user, opts))
	if config.Get().GetDebug() {
		mux.Handle(debugconnect.NewEchoServiceHandler(echo, opts))
	}
	return mux
}
