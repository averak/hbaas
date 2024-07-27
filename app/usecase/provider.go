package usecase

import (
	"github.com/averak/hbaas/app/usecase/echo_usecase"
	"github.com/averak/hbaas/app/usecase/session_usecase"
	"github.com/google/wire"
)

var SuperSet = wire.NewSet(
	session_usecase.New,
	echo_usecase.New,
)
