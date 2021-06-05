package customerselfdomain

import (
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"
	idcore "stock-contexts/pkg/shared/domain/id"

	"github.com/google/uuid"
)

// LicenseID struct
type LicenseID struct {
	entity domaincore.Entity
	Result appcore.IResult
}

// GetID method
func (s *LicenseID) GetID() string {
	return s.entity.ID.Identifier.ToString()
}

// GetUUID method
func (s *LicenseID) GetUUID() (uuid.UUID, error) {
	uuid, err := uuid.Parse(s.entity.ID.Identifier.Value)
	return uuid, err
}

// GetIDEntity method
func (s *LicenseID) GetIDEntity() *idcore.UniqueEntityID {
	return s.entity.ID
}

// NewLicenseID constructor
func NewLicenseID(ID *idcore.UniqueEntityID) LicenseID {
	return LicenseID{
		entity: domaincore.NewEntity(ID),
		Result: appcore.NewResultOk(),
	}
}
