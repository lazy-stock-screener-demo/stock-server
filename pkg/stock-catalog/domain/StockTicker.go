package catalogdomain

import (
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"
)

// StockTickerProps struct
type StockTickerProps struct {
	Value string
}

// StockTicker struct
type StockTicker struct {
	valueObject domaincore.IValueObject
	props       StockTickerProps
	Result      appcore.IResult
}

// GetValue method
func (s StockTicker) GetValue() string {
	return s.props.Value
}

// NewStockTicker constructor
func NewStockTicker(props StockTickerProps) StockTicker {
	guard := false
	if guard == false {
		return StockTicker{
			valueObject: domaincore.NewValueObject(props),
			props:       props,
			Result: appcore.NewResultFailed(appcore.ErrProps{
				Message: "message",
			}),
		}
	}

	//OK case
	return StockTicker{
		valueObject: domaincore.NewValueObject(props),
		props:       props,
		Result:      appcore.NewResultOk(),
	}
}

// // StockTickerProps struct
// type StockTickerProps struct {
// 	Value string
// }

// // StockTicker struct
// type StockTicker struct {
// 	valueObject domaincore.IValueObject
// 	Props       StockTickerProps
// }

// // GetValue method
// func (s StockTicker) GetValue() string {
// 	return s.Props.Value
// }

// // NewStockTicker constructor
// func NewStockTicker(props StockTickerProps) appcore.Result {
// 	return appcore.NewResultOk(StockTicker{
// 		valueObject: domaincore.NewValueObject(props),
// 		Props:       props,
// 	},
// 	)
// }
