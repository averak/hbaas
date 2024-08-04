package repoimpl

import (
	"github.com/averak/hbaas/app/adapter/repoimpl/authentication_repoimpl"
	"github.com/averak/hbaas/app/adapter/repoimpl/echo_repoimpl"
	"github.com/averak/hbaas/app/adapter/repoimpl/global_kvs_repoimpl"
	"github.com/averak/hbaas/app/adapter/repoimpl/leader_board_repoimpl"
	"github.com/averak/hbaas/app/adapter/repoimpl/private_kvs_repoimpl"
	"github.com/averak/hbaas/app/adapter/repoimpl/room_repoimpl"
	"github.com/averak/hbaas/app/adapter/repoimpl/user_profile_repoimpl"
	"github.com/averak/hbaas/app/adapter/repoimpl/user_repoimpl"
	"github.com/google/wire"
)

var SuperSet = wire.NewSet(
	authentication_repoimpl.NewRepository,
	echo_repoimpl.NewRepository,
	global_kvs_repoimpl.NewRepository,
	leader_board_repoimpl.NewRepository,
	private_kvs_repoimpl.NewRepository,
	room_repoimpl.NewRepository,
	user_profile_repoimpl.NewRepository,
	user_repoimpl.NewRepository,
)
