package usecase

import (
	"github.com/averak/hbaas/app/usecase/echo_usecase"
	"github.com/averak/hbaas/app/usecase/global_kvs_usecase"
	"github.com/averak/hbaas/app/usecase/session_usecase"
	"github.com/google/wire"
)

var SuperSet = wire.NewSet(
	session_usecase.NewUsecase,
	global_kvs_usecase.NewUsecase,
	echo_usecase.NewUsecase,
)