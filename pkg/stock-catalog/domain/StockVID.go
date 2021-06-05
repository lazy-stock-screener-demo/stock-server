package catalogdomain

import (
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"
)

// StockVIDProps struct
type StockVIDProps struct {
	value string
}

// StockVID struct
type StockVID struct {
	valueObject domaincore.IValueObject
	props       StockVIDProps
	Result      appcore.IResult
}

// GetValue method
func (s StockVID) GetValue() string {
	return s.props.value
}

// NewStockVID constructor
func NewStockVID(props StockVIDProps) StockVID {
	return StockVID{
		valueObject: domaincore.NewValueObject(props),
		props:       props,
		Result:      appcore.NewResultOk(),
	}
}

// // StockVIDProps struct
// type StockVIDProps struct {
// 	value string
// }

// // StockVID struct
// type StockVID struct {
// 	valueObject domaincore.IValueObject
// 	props       StockVIDProps
// }

// // GetValue method
// func (s StockVID) GetValue() string {
// 	return s.props.value
// }

// // NewStockVID constructor
// func NewStockVID(props StockVIDProps) appcore.Result {
// 	return appcore.NewResultOk(StockVID{
// 		valueObject: domaincore.NewValueObject(props),
// 		props:       props,
// 	},
// 	)
// }
