package usecaseimpl

import (
	"github.com/google/wire"
)

var SuperSet = wire.NewSet(
	NewFirebaseIdentityVerifier,
	NewBaasUserDeletionTaskQueue,
)
