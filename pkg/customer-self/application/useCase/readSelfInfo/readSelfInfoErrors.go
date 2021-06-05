package readselfinfo

import (
	appcore "stock-contexts/pkg/shared/application"
)

// NewCustomerNotFoundErr define implicit error type
func NewCustomerNotFoundErr(err interface{}) appcore.Result {
	return appcore.NewResultUseCaseErr(
		appcore.ErrProps{
			Message: "Couldn't find a Customer by customer vid:",
			Type:    "CustomerNotFound",
			Error:   err,
		},
	)
}
