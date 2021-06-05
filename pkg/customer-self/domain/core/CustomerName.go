package customerselfdomain

import (
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"
)

// CustomerNameProps struct
type CustomerNameProps struct {
	Value string
}

// CustomerName struct
type CustomerName struct {
	valueObject domaincore.IValueObject
	props       CustomerNameProps
	Result      appcore.IResult
}

// GetValue method
func (s CustomerName) GetValue() string {
	return s.props.Value
}

// NewCustomerName constructor
func NewCustomerName(props CustomerNameProps) CustomerName {
	guard := appcore.NewGurad()
	if validation := guard.AgainstAtMost(22, props.Value); validation.Passed == false {
		return CustomerName{
			valueObject: domaincore.NewValueObject(props),
			props:       props,
			Result: appcore.NewResultFailed(
				appcore.ErrProps{
					Error: validation.Message,
				},
			),
		}
	}
	if validation := guard.AgainstAtLeast(2, props.Value); validation.Passed == false {
		return CustomerName{
			valueObject: domaincore.NewValueObject(props),
			props:       props,
			Result: appcore.NewResultFailed(
				appcore.ErrProps{
					Error: validation.Message,
				},
			),
		}
	}
	// OK case
	return CustomerName{
		valueObject: domaincore.NewValueObject(props),
		props:       props,
		Result:      appcore.NewResultOk(),
	}
}
