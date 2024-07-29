//go:build wireinject
// +build wireinject

package registry

import (
	"context"
	"net/http"

	"github.com/averak/hbaas/app/adapter/handler"
	"github.com/averak/hbaas/app/adapter/repoimpl"
	"github.com/averak/hbaas/app/adapter/usecaseimpl"
	"github.com/averak/hbaas/app/core/config"
	"github.com/averak/hbaas/app/infrastructure/connect/advice"
	"github.com/averak/hbaas/app/infrastructure/db"
	"github.com/averak/hbaas/app/infrastructure/google_cloud"
	"github.com/averak/hbaas/app/usecase"
	"github.com/averak/hbaas/testutils/testgoogle_cloud"
	"github.com/google/wire"
)

var SuperSet = wire.NewSet(
	repoimpl.SuperSet,
	usecaseimpl.SuperSet,
	usecase.SuperSet,
	db.NewConnection,
	advice.NewAdvice,
	google_cloud.NewPubSubClient,
	newFirebaseClient,
)

func newFirebaseClient(ctx context.Context) (google_cloud.FirebaseClient, error) {
	if config.Get().GetGoogleCloud().GetFirebase().GetUseStub() {
		return testgoogle_cloud.NewFirebaseClientStub(), nil
	}
	return google_cloud.NewFirebaseClient(ctx)
}

func InitializeAPIServerMux(ctx context.Context) (*http.ServeMux, error) {
	wire.Build(SuperSet, handler.SuperSet)
	return nil, nil
}
