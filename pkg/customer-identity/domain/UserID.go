package identitydomain

import (
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"
	idcore "stock-contexts/pkg/shared/domain/id"

	"github.com/google/uuid"
)

// UserID struct
type UserID struct {
	entity domaincore.Entity
	Result appcore.IResult
}

// GetID method
func (s *UserID) GetID() string {
	return s.entity.ID.Identifier.ToString()
}

// GetUUID method
func (s *UserID) GetUUID() (uuid.UUID, error) {
	id, err := uuid.Parse(s.entity.ID.Identifier.Value)
	return id, err
}

// GetIDEntity method
func (s *UserID) GetIDEntity() *idcore.UniqueEntityID {
	return s.entity.ID
}

// NewUserID constructor
func NewUserID(ID *idcore.UniqueEntityID) UserID {
	return UserID{
		entity: domaincore.NewEntity(ID),
		Result: appcore.NewResultOk(),
	}
}
