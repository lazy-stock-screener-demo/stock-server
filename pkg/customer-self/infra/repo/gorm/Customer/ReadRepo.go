package customerrepo

import (
	customerdto "stock-contexts/pkg/customer-self/application/dto"
	customermapper "stock-contexts/pkg/customer-self/application/mapper/Customer"
	customerselfdomain "stock-contexts/pkg/customer-self/domain/core"
	customerselfschema "stock-contexts/pkg/customer-self/infra/repo/schema"
	appcore "stock-contexts/pkg/shared/application"
	gormclient "stock-contexts/pkg/shared/infra/repo/gorm/config"

	"gorm.io/gorm"
)

var customerViewModel customerselfschema.CustomerView
var customerModel customerselfschema.Customer
var licenseModel customerselfschema.License

// IRead as interface
type IRead interface {
	// ReadCustomerViewByVID(VID string) (*customerselfdomain.CustomerView, appcore.ErrProps)
	ReadCustomerViewByName(cus customerselfdomain.CustomerName) (*customerdto.CustomerDTO, appcore.ErrProps)
	// ReadCustomerByName(cus customerselfdomain.CustomerName) (customerselfdomain.Customer, appcore.ErrProps)
	// ReadCustomerByVID(VID string) (domain customerselfdomain.Customer, err appcore.ErrProps)
}

// Read Define reac operation of database ORM
type Read struct {
	client *gorm.DB
}

// // ReadCustomerViewByVID define the method to read of customer by VID
// func (r *Read) ReadCustomerViewByVID(VID string) (*customerselfdomain.CustomerView, appcore.ErrProps) {
// 	r.client.Find(&customerModel, "customerVID = ?", VID)
// 	customerView := customermapper.NewViewMap().ToDomain(&customerModel)
// 	return customerView, appcore.ErrProps{}
// }

// ReadCustomerViewByName define the method to read of customer by name
func (r *Read) ReadCustomerViewByName(name customerselfdomain.CustomerName) (*customerdto.CustomerDTO, appcore.ErrProps) {
	r.client.Model(&customerModel).Select("customer.customer_name, customer.email, license.license_name").Joins("left join license on customer.license_id = license.id").Scan(&customerViewModel)
	customerView := customermapper.NewViewMap().ToDTO(&customerViewModel)
	return &customerView, appcore.ErrProps{}
}

// // ReadCustomerByName define the method to read of customer by name
// func (r *Read) ReadCustomerByName(name customerselfdomain.CustomerName) (customerselfdomain.Customer, appcore.ErrProps) {
// 	r.client.Find(&readModel, "customer_name = ?", name.GetValue())
// 	customer := customermapper.NewEntityMap().ToDomain(&readModel)
// 	return customer, appcore.ErrProps{}
// }

// NewReadRepo define ReadRepo Instance
func NewReadRepo() *Read {
	return &Read{
		client: gormclient.NewConnectedGorm(),
	}
}
