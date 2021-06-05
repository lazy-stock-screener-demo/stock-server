package customerselfdomain

import (
	"time"
)

type CustomerCreatedProps struct {
	timeOccured time.Time
	customer    Customer
}

type CustomerCreated struct {
	props CustomerCreatedProps
}

func (c CustomerCreated) GetAggregateID() string {
	id := c.props.customer.GetCustomerID()
	return id.GetID()
}

func (c CustomerCreated) GetCustomerName() CustomerName {
	return c.props.customer.props.CustomerName
}

func (c CustomerCreated) GetCustomerPWD() CustomerPWD {
	return c.props.customer.props.CustomerPWD
}

func NewCustomerCreated(props CustomerCreatedProps) CustomerCreated {
	return CustomerCreated{
		props: props,
	}
}
