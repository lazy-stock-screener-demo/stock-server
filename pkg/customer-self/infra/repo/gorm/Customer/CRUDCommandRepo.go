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

var crudModel customerselfschema.Customer

// ICRUDCommand define interface
type ICRUDCommand interface {
	isExistByCustomerName(customerName customerselfdomain.CustomerName) (bool, appcore.ErrProps)
	isExistByCustomerID(customerID customerselfdomain.CustomerID) (bool, appcore.ErrProps)
	Save(customer customerselfdomain.Customer) (*customerdto.CustomerDTO, appcore.ErrProps)
	Delete(customer customerselfdomain.Customer)
}

// CRUDCommand Define reac operation of database ORM
type CRUDCommand struct {
	client *gorm.DB
}

func (c *CRUDCommand) isExistByCustomerName(customerName customerselfdomain.CustomerName) (bool, appcore.ErrProps) {
	result := c.client.Limit(1).Find(&crudModel, "customer_name = ?", customerName.GetValue())
	if result.Error != nil {
		return false, appcore.ErrProps{Error: result.Error}
	}
	if result.RowsAffected > 0 {
		return true, appcore.ErrProps{}
	}
	return false, appcore.ErrProps{}
}

func (c *CRUDCommand) isExistByCustomerID(customerID customerselfdomain.CustomerID) (bool, appcore.ErrProps) {
	result := c.client.Limit(1).Find(&crudModel, "id = ?", customerID.GetID())
	if result.Error != nil {
		return false, appcore.ErrProps{Error: result.Error}
	}
	if result.RowsAffected > 0 {
		return true, appcore.ErrProps{}
	}
	return false, appcore.ErrProps{}
}

// Delete method
func (c *CRUDCommand) Delete(cus customerselfdomain.Customer) {
	customerID := cus.GetCustomerID()
	c.client.Delete(&customerselfschema.Customer{}, customerID.GetID())
}

// Save define command method
func (c *CRUDCommand) Save(cus customerselfdomain.Customer) (*customerdto.CustomerDTO, appcore.ErrProps) {
	var errProp interface{}
	customerID := cus.GetCustomerID()
	isExist, err := c.isExistByCustomerID(customerID)
	if err.Error != nil {
		errProp = err
	}
	rawCustomer := customermapper.NewEntityMap().ToPersistence(cus)
	// 	// Update
	if isExist {
		if result := c.client.Model(&customerselfschema.Customer{}).Updates(rawCustomer); result.Error != nil {
			errProp = result.Error
		}
		customerView := customermapper.NewEntityMap().ToDTO(rawCustomer)
		return &customerView, appcore.ErrProps{Error: errProp}
	}
	// Create
	if result := c.client.Model(&customerselfschema.Customer{}).Create(rawCustomer); result.Error != nil {
		c.client.Delete(&customerselfschema.Customer{}, "id = ?", customerID.GetID())
		errProp = result.Error
	}
	customerView := customermapper.NewEntityMap().ToDTO(rawCustomer)
	return &customerView, appcore.ErrProps{Error: errProp}
}

// NewCRUDCommandRepo define ReadRepo Instance
func NewCRUDCommandRepo() *CRUDCommand {
	return &CRUDCommand{
		client: gormclient.NewConnectedGorm(),
	}
}
