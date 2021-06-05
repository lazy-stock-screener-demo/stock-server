package customerselfdomain

import (
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"
	idcore "stock-contexts/pkg/shared/domain/id"

	"github.com/google/uuid"
)

// CustomerID struct
type CustomerID struct {
	entity domaincore.Entity
	Result appcore.IResult
}

// GetID method
func (s *CustomerID) GetID() string {
	return s.entity.ID.Identifier.ToString()
}

// GetUUID method
func (s *CustomerID) GetUUID() (uuid.UUID, error) {
	id, err := uuid.Parse(s.entity.ID.Identifier.Value)
	return id, err
}

// GetIDEntity method
func (s *CustomerID) GetIDEntity() *idcore.UniqueEntityID {
	return s.entity.ID
}

// NewCustomerID constructor
func NewCustomerID(ID *idcore.UniqueEntityID) CustomerID {
	return CustomerID{
		entity: domaincore.NewEntity(ID),
		Result: appcore.NewResultOk(),
	}
}
