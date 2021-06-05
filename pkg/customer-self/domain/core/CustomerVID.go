package customerselfdomain

import (
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"
)

// CustomerVIDProps struct
type CustomerVIDProps struct {
	value string
}

// CustomerVID struct
type CustomerVID struct {
	valueObject domaincore.IValueObject
	props       CustomerVIDProps
	Result      appcore.IResult
}

// GetValue method
func (s CustomerVID) GetValue() string {
	return s.props.value
}

// NewCustomerVID constructor
func NewCustomerVID(props CustomerVIDProps) CustomerVID {
	return CustomerVID{
		valueObject: domaincore.NewValueObject(props),
		props:       props,
		Result:      appcore.NewResultOk(),
	}
}
