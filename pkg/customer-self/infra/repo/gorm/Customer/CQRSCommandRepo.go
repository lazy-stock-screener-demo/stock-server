package customerrepo

import (
	customermapper "stock-contexts/pkg/customer-self/application/mapper/Customer"
	customerselfdomain "stock-contexts/pkg/customer-self/domain/core"
	customerselfschema "stock-contexts/pkg/customer-self/infra/repo/schema"
	appcore "stock-contexts/pkg/shared/application"
	gormclient "stock-contexts/pkg/shared/infra/repo/gorm/config"

	"gorm.io/gorm"
)

var cqrsModel customerselfschema.Customer

// ICQRSCommand as interface
type ICQRSCommand interface {
	isExistByCustomerID(customerID customerselfdomain.CustomerID) (bool, appcore.ErrProps)
	isExistByCustomerName(customerName customerselfdomain.CustomerName) (bool, appcore.ErrProps)
	Save(customer customerselfdomain.Customer) appcore.ErrProps
	Delete(customer customerselfdomain.Customer)
}

// CQRSCommand Define reac operation of database ORM
type CQRSCommand struct {
	client *gorm.DB
}

func (c *CQRSCommand) isExistByCustomerName(customerName customerselfdomain.CustomerName) (bool, appcore.ErrProps) {
	result := c.client.Limit(1).Find(&cqrsModel, "customerName = ?", customerName.GetValue())
	if result.Error != nil {
		return false, appcore.ErrProps{Error: result.Error}
	}
	if result.RowsAffected >= 0 {
		return true, appcore.ErrProps{}
	}
	return false, appcore.ErrProps{}
}

func (c *CQRSCommand) isExistByCustomerID(customerID customerselfdomain.CustomerID) (bool, appcore.ErrProps) {
	result := c.client.Limit(1).Find(&cqrsModel, "customerVID = ?", customerID.GetID())
	if result.Error != nil {
		return false, appcore.ErrProps{Error: result.Error}
	}
	if result.RowsAffected >= 0 {
		return true, appcore.ErrProps{}
	}
	return false, appcore.ErrProps{}
}

// Delete method
func (c *CQRSCommand) Delete(cus customerselfdomain.Customer) {
	customerID := cus.GetCustomerID()
	c.client.Delete(&customerselfschema.Customer{}, customerID.GetID())
}

// Save define CQRS command
func (c *CQRSCommand) Save(cus customerselfdomain.Customer) appcore.ErrProps {
	var errProp interface{}
	customerID := cus.GetCustomerID()
	isExist, err := c.isExistByCustomerID(customerID)
	if err.Error != nil {
		errProp = err
	}
	rawCustomer := customermapper.NewEntityMap().ToPersistence(cus)
	// Update
	if isExist {
		if result := c.client.Model(&customerselfschema.Customer{}).Updates(rawCustomer); result.Error != nil {
			errProp = result.Error
		}
		return appcore.ErrProps{Error: errProp}
	}
	// Create
	if result := c.client.Model(&customerselfschema.Customer{}).Create(rawCustomer); result.Error != nil {
		c.client.Delete(&customerselfschema.Customer{}, customerID.GetID())
		errProp = result.Error
	}
	return appcore.ErrProps{Error: errProp}
}

// NewCQRSCommandRepo define ReadRepo Instance
func NewCQRSCommandRepo() *CQRSCommand {
	return &CQRSCommand{
		client: gormclient.NewConnectedGorm(),
	}
}
