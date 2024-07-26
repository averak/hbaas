//go:build wireinject
// +build wireinject

package registry

import (
	"context"
	"net/http"

	"github.com/averak/hbaas/app/adapter/handler"
	"github.com/averak/hbaas/app/adapter/repoimpl"
	"github.com/averak/hbaas/app/infrastructure/db"
	"github.com/averak/hbaas/app/usecase"
	"github.com/google/wire"
)

var SuperSet = wire.NewSet(
	repoimpl.SuperSet,
	usecase.SuperSet,
	db.NewConnection,
)

func InitializeAPIServerMux(ctx context.Context) (*http.ServeMux, error) {
	wire.Build(SuperSet, handler.SuperSet)
	return nil, nil
}
