package usecase

import (
	"github.com/averak/hbaas/app/usecase/echo_usecase"
	"github.com/google/wire"
)

var SuperSet = wire.NewSet(
	echo_usecase.New,
)
