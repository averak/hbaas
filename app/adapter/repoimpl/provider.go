package repoimpl

import (
	"github.com/averak/hbaas/app/adapter/repoimpl/authentication_repoimpl"
	"github.com/averak/hbaas/app/adapter/repoimpl/echo_repoimpl"
	"github.com/averak/hbaas/app/adapter/repoimpl/global_kvs_repoimpl"
	"github.com/averak/hbaas/app/adapter/repoimpl/user_repoimpl"
	"github.com/google/wire"
)

var SuperSet = wire.NewSet(
	authentication_repoimpl.NewRepository,
	echo_repoimpl.NewRepository,
	global_kvs_repoimpl.NewRepository,
	user_repoimpl.NewRepository,
)