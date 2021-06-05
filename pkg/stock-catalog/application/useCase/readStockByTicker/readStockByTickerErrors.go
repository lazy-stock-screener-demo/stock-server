package readstockbyticker

import (
	appcore "stock-contexts/pkg/shared/application"
)

// TickerNotFoundErr define exact error type
// type TickerNotFoundErr struct {
// 	appcore.Result
// }

// // NewTickerNotFoundErr define implicit error type
// func NewTickerNotFoundErr(err interface{}) appcore.Result {
// 	// t = template.Must(template.New("t2").
// 	// 	Parse("Couldn't find a stock by ticker view id: {{.a}}\n"))
// 	return appcore.NewResultUseCaseErr(
// 		appcore.ErrProps{
// 			Message: "Couldn't find a stock by ticker view id:",
// 			Type:    "TickerNotFound",
// 			Error:   err,
// 		},
// 	)
// }

// NewTickerNotFoundErr define implicit error type
func NewTickerNotFoundErr(err interface{}) appcore.Result {
	// t = template.Must(template.New("t2").
	// 	Parse("Couldn't find a stock by ticker view id: {{.a}}\n"))
	return appcore.NewResultUseCaseErr(
		appcore.ErrProps{
			Message: "Couldn't find a stock by ticker view id:",
			Type:    "TickerNotFound",
			Error:   err,
		},
	)
}
