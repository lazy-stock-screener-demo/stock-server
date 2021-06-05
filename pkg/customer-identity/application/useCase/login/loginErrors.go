package login

import (
	appcore "stock-contexts/pkg/shared/application"
)

// NewUserNotFoundErr define implicit error type
func NewUserNotFoundErr(err interface{}) appcore.Result {
	return appcore.NewResultUseCaseErr(
		appcore.ErrProps{
			Message: "Couldn't user by user name",
			Type:    "UserNotFound",
			Error:   err,
		},
	)
}
