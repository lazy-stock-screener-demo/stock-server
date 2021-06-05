package catalogdomain

import (
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"
)

// IStockView provide a way to access its own value
type IStockView interface {
	GetStockVID() StockVID
	GetStockTicker() StockTicker
}

// StockViewProps struct
type StockViewProps struct {
	StockVID
	StockTicker
	ProfitDetails
	SafetyDetails
}

// StockView struct
type StockView struct {
	valueObject domaincore.IValueObject
	props       StockViewProps
	Result      appcore.IResult
}

// GetStockVID method
func (s StockView) GetStockVID() StockVID {
	return s.props.StockVID
}

// GetStockTicker method
func (s StockView) GetStockTicker() StockTicker {
	return s.props.StockTicker
}

// GetProfitDetails method
func (s StockView) GetProfitDetails() ProfitDetails {
	return s.props.ProfitDetails
}

// GetSafetyDetails method
func (s StockView) GetSafetyDetails() SafetyDetails {
	return s.props.SafetyDetails
}

// NewStockView define constructor
func NewStockView(props StockViewProps) StockView {
	return StockView{
		valueObject: domaincore.NewValueObject(props),
		props:       props,
		Result:      appcore.NewResultOk(),
	}
}

// // StockViewProps struct
// type StockViewProps struct {
// 	StockVID    StockVID
// 	StockTicker StockTicker
// }

// // StockView struct
// type StockView struct {
// 	valueObject domaincore.IValueObject
// 	props       StockViewProps
// }

// // GetStockVID method
// func (s StockView) GetStockVID() StockVID {
// 	return s.props.StockVID
// }

// // GetStockTicker method
// func (s StockView) GetStockTicker() StockTicker {
// 	return s.props.StockTicker
// }

// // NewStockView constructor
// func NewStockView(props StockViewProps) appcore.Result {
// 	return appcore.NewResultOk(StockView{
// 		valueObject: domaincore.NewValueObject(props),
// 		props:       props,
// 	})
// }
