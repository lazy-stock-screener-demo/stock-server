package catalogdomain

import (
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"
	idcore "stock-contexts/pkg/shared/domain/id"
)

// StockProps as props struct
type StockProps struct {
	StockTicker StockTicker
}

// Stock stuct
type Stock struct {
	agEventHandler domaincore.DomainEventHandler
	props          StockProps
	Entity         domaincore.Entity
	Result         appcore.IResult
}

// GetStockID Method
func (s Stock) GetStockID() StockID {
	return NewStockID(s.Entity.ID)
}

// GetStockTicker define
func (s Stock) GetStockTicker() StockTicker {
	return s.props.StockTicker
}

// NewStock to create
func NewStock(ID *idcore.UniqueEntityID, props StockProps) Stock {
	return Stock{
		agEventHandler: domaincore.NewDomainEventHandler(),
		Entity:         domaincore.NewEntity(ID),
		props:          props,
		Result:         appcore.NewResultOk(),
	}
}

// // StockProps as props struct
// type StockProps struct {
// 	StockTicker StockTicker
// }

// // Stock stuct
// type Stock struct {
// 	domaincore.AggregateRoot
// 	StockProps
// }

// // GetStockID Method
// func (s Stock) GetStockID() StockID {
// 	stockID := NewStockID(s.Entity.ID).GetData().Data
// 	return stockID.(StockID)
// }

// // GetStockTicker define
// func (s Stock) GetStockTicker() StockTicker {
// 	return s.StockTicker
// }

// // NewStock to create
// func NewStock(ID *idcore.UniqueEntityID, props StockProps) appcore.Result2 {
// 	return appcore.NewResultOk(Stock{
// 		domaincore.NewAggregateRoot(ID),
// 		props,
// 	},
// 	)
// }
