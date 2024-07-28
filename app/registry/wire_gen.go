// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package registry

import (
	"context"
	"github.com/averak/hbaas/app/adapter/handler"
	"github.com/averak/hbaas/app/adapter/handler/debug/echo"
	"github.com/averak/hbaas/app/adapter/handler/global_kvs"
	"github.com/averak/hbaas/app/adapter/handler/leader_board"
	"github.com/averak/hbaas/app/adapter/handler/private_kvs"
	"github.com/averak/hbaas/app/adapter/handler/session"
	"github.com/averak/hbaas/app/adapter/repoimpl"
	"github.com/averak/hbaas/app/adapter/repoimpl/authentication_repoimpl"
	"github.com/averak/hbaas/app/adapter/repoimpl/echo_repoimpl"
	"github.com/averak/hbaas/app/adapter/repoimpl/global_kvs_repoimpl"
	"github.com/averak/hbaas/app/adapter/repoimpl/leader_board_repoimpl"
	"github.com/averak/hbaas/app/adapter/repoimpl/private_kvs_repoimpl"
	"github.com/averak/hbaas/app/adapter/repoimpl/user_repoimpl"
	"github.com/averak/hbaas/app/adapter/usecaseimpl"
	"github.com/averak/hbaas/app/core/config"
	"github.com/averak/hbaas/app/infrastructure/connect/advice"
	"github.com/averak/hbaas/app/infrastructure/db"
	"github.com/averak/hbaas/app/infrastructure/google_cloud"
	"github.com/averak/hbaas/app/usecase"
	"github.com/averak/hbaas/app/usecase/echo_usecase"
	"github.com/averak/hbaas/app/usecase/global_kvs_usecase"
	"github.com/averak/hbaas/app/usecase/leader_board_usecase"
	"github.com/averak/hbaas/app/usecase/private_kvs_usecase"
	"github.com/averak/hbaas/app/usecase/session_usecase"
	"github.com/averak/hbaas/testutils/testgoogle_cloud"
	"github.com/google/wire"
	"net/http"
)

// Injectors from wire.go:

func InitializeAPIServerMux(ctx context.Context) (*http.ServeMux, error) {
	connection, err := db.NewConnection()
	if err != nil {
		return nil, err
	}
	globalKVSRepository := global_kvs_repoimpl.NewRepository()
	usecase := global_kvs_usecase.NewUsecase(connection, globalKVSRepository)
	userRepository := user_repoimpl.NewRepository()
	adviceAdvice := advice.NewAdvice(connection, userRepository)
	globalKVSServiceHandler := global_kvs.NewHandler(usecase, adviceAdvice)
	leaderBoardRepository := leader_board_repoimpl.NewRepository()
	leader_board_usecaseUsecase := leader_board_usecase.NewUsecase(connection, leaderBoardRepository)
	leaderBoardServiceHandler := leader_board.NewHandler(leader_board_usecaseUsecase, adviceAdvice)
	privateKVSRepository := private_kvs_repoimpl.NewRepository()
	private_kvs_usecaseUsecase := private_kvs_usecase.NewUsecase(connection, privateKVSRepository)
	privateKVSServiceHandler := private_kvs.NewHandler(private_kvs_usecaseUsecase, adviceAdvice)
	firebaseClient, err := newFirebaseClient(ctx)
	if err != nil {
		return nil, err
	}
	identityVerifier := usecaseimpl.NewFirebaseIdentityVerifier(firebaseClient)
	authenticationRepository := authentication_repoimpl.NewRepository()
	session_usecaseUsecase := session_usecase.NewUsecase(connection, identityVerifier, authenticationRepository, userRepository)
	sessionServiceHandler := session.NewHandler(session_usecaseUsecase, adviceAdvice)
	echoRepository := echo_repoimpl.NewRepository()
	echo_usecaseUsecase := echo_usecase.NewUsecase(connection, echoRepository)
	echoServiceHandler := echo.NewHandler(echo_usecaseUsecase, adviceAdvice)
	serveMux := handler.New(globalKVSServiceHandler, leaderBoardServiceHandler, privateKVSServiceHandler, sessionServiceHandler, echoServiceHandler)
	return serveMux, nil
}

// wire.go:

var SuperSet = wire.NewSet(repoimpl.SuperSet, usecaseimpl.SuperSet, usecase.SuperSet, db.NewConnection, advice.NewAdvice, newFirebaseClient)

func newFirebaseClient(ctx context.Context) (google_cloud.FirebaseClient, error) {
	if config.Get().GetGoogleCloud().GetFirebase().GetUseStub() {
		return testgoogle_cloud.NewFirebaseClientStub(), nil
	}
	return google_cloud.NewFirebaseClient(ctx)
}
