package licenserepo

import (
	licensemapper "stock-contexts/pkg/customer-self/application/mapper/License"
	customerselfdomain "stock-contexts/pkg/customer-self/domain/core"
	customerselfschema "stock-contexts/pkg/customer-self/infra/repo/schema"
	appcore "stock-contexts/pkg/shared/application"
	gormclient "stock-contexts/pkg/shared/infra/repo/gorm/config"

	"gorm.io/gorm"
)

var readModel customerselfschema.License

// IRead as interface
type IRead interface {
	ReadLicenseByID(ID customerselfdomain.CustomerID) (customerselfdomain.License, appcore.ErrProps)
	ReadLicenseByName(name customerselfdomain.LicenseName) (customerselfdomain.License, appcore.ErrProps)
}

// Read Define reac operation of database ORM
type Read struct {
	client *gorm.DB
}

// ReadLicenseByID define the method to read of stock by ticker
func (r *Read) ReadLicenseByID(ID customerselfdomain.CustomerID) (customerselfdomain.License, appcore.ErrProps) {
	r.client.Find(&readModel, "id = ?", ID.GetID())
	license := licensemapper.NewEntityMap().ToDomain(&readModel)
	return license, appcore.ErrProps{}
}

// ReadLicenseByName define the method to read of stock by ticker
func (r *Read) ReadLicenseByName(name customerselfdomain.LicenseName) (customerselfdomain.License, appcore.ErrProps) {
	r.client.Find(&readModel, "license_name = ?", name.GetValue())
	license := licensemapper.NewEntityMap().ToDomain(&readModel)
	return license, appcore.ErrProps{}
}

// NewReadRepo define ReadRepo Instance
func NewReadRepo() *Read {
	return &Read{
		client: gormclient.NewConnectedGorm(),
	}
}
