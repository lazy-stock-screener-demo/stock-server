package catalogdomain

import (
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"
	idcore "stock-contexts/pkg/shared/domain/id"
)

// StockID struct
type StockID struct {
	entity domaincore.Entity
	Result appcore.IResult
}

// GetID method
func (s *StockID) GetID() string {
	return s.entity.ID.Identifier.ToString()
}

// GetUID method
func (s *StockID) GetIDEntity() *idcore.UniqueEntityID {
	return s.entity.ID
}

// NewStockID constructor
func NewStockID(ID *idcore.UniqueEntityID) StockID {
	return StockID{
		entity: domaincore.NewEntity(ID),
		Result: appcore.NewResultOk(),
	}
}

// // StockID struct
// type StockID struct {
// 	domaincore.Entity
// }

// // GetID method
// func (s *StockID) GetID() idcore.UniqueEntityID {
// 	return *s.ID
// }

// // NewStockID constructor
// func NewStockID(ID *idcore.UniqueEntityID) appcore.Result {
// 	return appcore.NewResultOk(StockID{
// 		domaincore.NewEntity(ID),
// 	},
// 	)
// }
