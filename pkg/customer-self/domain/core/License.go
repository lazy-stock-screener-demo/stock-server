package customerselfdomain

import (
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"
	idcore "stock-contexts/pkg/shared/domain/id"
)

// LicenseProps as props struct
type LicenseProps struct {
	LicenseName
}

// License stuct
type License struct {
	agEventHandler domaincore.DomainEventHandler
	Entity         domaincore.Entity
	props          LicenseProps
	Result         appcore.IResult
}

// GetLicenseID Method
func (s *License) GetLicenseID() LicenseID {
	return NewLicenseID(s.Entity.ID)
}

func (s *License) GetLicenseName() LicenseName {
	return s.props.LicenseName
}

// NewLicense to create
func NewLicense(ID *idcore.UniqueEntityID, props LicenseProps) License {
	// TODO: Handle creation failed case
	// Handle OK
	return License{
		agEventHandler: domaincore.NewDomainEventHandler(),
		Entity:         domaincore.NewEntity(ID),
		props:          props,
		Result:         appcore.NewResultOk(),
	}
}
