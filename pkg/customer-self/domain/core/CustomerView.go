package customerselfdomain

import (
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"
)

// CustomerViewProps struct
type CustomerViewProps struct {
	CustomerVID
	CustomerName
	CustomerPWD
	CustomerEmail
}

// CustomerView struct
type CustomerView struct {
	valueObject domaincore.IValueObject
	props       CustomerViewProps
	Result      appcore.IResult
}

// GetStockVID method
func (c CustomerView) GetCustomerName() CustomerName {
	return c.props.CustomerName
}

// GetStockVID method
func (c CustomerView) GetEmail() CustomerEmail {
	return c.props.CustomerEmail
}

// NewCustomerView define constructor
func NewCustomerView(props CustomerViewProps) *CustomerView {
	return &CustomerView{
		valueObject: domaincore.NewValueObject(props),
		props:       props,
		Result:      appcore.NewResultOk(),
	}
}
