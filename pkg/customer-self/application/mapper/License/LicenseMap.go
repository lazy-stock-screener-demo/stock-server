package licensemapper

import (
	customerselfdomain "stock-contexts/pkg/customer-self/domain/core"
	customerselfschema "stock-contexts/pkg/customer-self/infra/repo/schema"
	idcore "stock-contexts/pkg/shared/domain/id"
)

// EntityMap struct
type EntityMap struct{}

// ToDomain Implemented
func (e EntityMap) ToDomain(raw *customerselfschema.License) customerselfdomain.License {
	licenseName := customerselfdomain.NewLicenseName(
		customerselfdomain.LicenseNameProps{
			Value: raw.LicenseName,
		})

	license := customerselfdomain.NewLicense(
		idcore.NewIDEntity(raw.ID.String()),
		customerselfdomain.LicenseProps{
			LicenseName: licenseName,
		},
	)
	return license
}

// NewEntityMap method
func NewEntityMap() EntityMap {
	return EntityMap{}
}
