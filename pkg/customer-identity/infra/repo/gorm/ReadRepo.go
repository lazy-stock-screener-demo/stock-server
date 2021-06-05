package identityrepo

import (
	identitymapper "stock-contexts/pkg/customer-identity/application/mapper/Identity"
	identitydomain "stock-contexts/pkg/customer-identity/domain"
	identityschema "stock-contexts/pkg/customer-identity/infra/repo/schema"
	appcore "stock-contexts/pkg/shared/application"
	gormclient "stock-contexts/pkg/shared/infra/repo/gorm/config"

	"gorm.io/gorm"
)

var readModel identityschema.User

// IRead as interface
type IRead interface {
	ReadCustomerByName(cus identitydomain.UserName) (identitydomain.User, appcore.ErrProps)
}

// Read Define reac operation of database ORM
type Read struct {
	client *gorm.DB
}

// ReadCustomerByName define the method to read of customer by name
func (r *Read) ReadCustomerByName(name identitydomain.UserName) (identitydomain.User, appcore.ErrProps) {
	r.client.Find(&readModel, "customer_name = ?", name.GetValue())
	customer := identitymapper.NewEntityMap().ToDomain(&readModel)
	return customer, appcore.ErrProps{}
}

// NewReadRepo define ReadRepo Instance
func NewReadRepo() *Read {
	return &Read{
		client: gormclient.NewConnectedGorm(),
	}
}
