package advice

import (
	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/infrastructure/connect/error_response"
	"github.com/averak/hbaas/protobuf/api/api_errors"
)

func checkUserStatus(user model.User) error {
	if user.Status == model.UserStatusActive {
		return nil
	}
	return error_response.New(api_errors.ErrorCode_COMMON_INVALID_USER_AVAILABILITY, api_errors.ErrorSeverity_ERROR_SEVERITY_WARNING, "user is not active")
}
