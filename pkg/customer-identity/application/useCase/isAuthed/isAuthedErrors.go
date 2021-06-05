package isauthed

import (
	appcore "stock-contexts/pkg/shared/application"
)

// NewUserNotFoundErr define implicit error type
func NewUserNotFoundErr(err interface{}) appcore.Result {
	return appcore.NewResultUseCaseErr(
		appcore.ErrProps{
			Message: "user is not login",
			Type:    "UserNotFound",
			Error:   err,
		},
	)
}

// NewUserNotFoundErr define implicit error type
func NewSignatureExpiredErr(err interface{}) appcore.Result {
	return appcore.NewResultUseCaseErr(
		appcore.ErrProps{
			Message: "Token signature expired.",
			Type:    "TokenSignatureExpired",
			Error:   err,
		},
	)
}

// NewUserNotFoundErr define implicit error type
func NewNoAccessTokenProvidedErr(err interface{}) appcore.Result {
	return appcore.NewResultUseCaseErr(
		appcore.ErrProps{
			Message: "No access token provided.",
			Type:    "NoAccessToken",
			Error:   err,
		},
	)
}
