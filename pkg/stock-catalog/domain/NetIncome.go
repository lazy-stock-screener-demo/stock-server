package catalogdomain

import (
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"
)

// NetIncomeProps struct
type NetIncomeProps struct {
	Value []map[string]float32
}

// NetIncome struct
type NetIncome struct {
	valueObject domaincore.IValueObject
	props       NetIncomeProps
	Result      appcore.IResult
}

// GetValue method
func (s NetIncome) GetValue() []map[string]float32 {
	return s.props.Value
}

// NewNetIncome define NetIncome value object
func NewNetIncome(props NetIncomeProps) NetIncome {
	//OK case
	return NetIncome{
		valueObject: domaincore.NewValueObject(props),
		props:       props,
		Result:      appcore.NewResultOk(),
	}
}
