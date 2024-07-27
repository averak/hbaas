package repoimpl

import (
	"github.com/averak/hbaas/app/adapter/repoimpl/authentication_repoimpl"
	"github.com/averak/hbaas/app/adapter/repoimpl/echo_repoimpl"
	"github.com/averak/hbaas/app/adapter/repoimpl/user_repoimpl"
	"github.com/google/wire"
)

var SuperSet = wire.NewSet(
	authentication_repoimpl.New,
	echo_repoimpl.New,
	user_repoimpl.New,
)
