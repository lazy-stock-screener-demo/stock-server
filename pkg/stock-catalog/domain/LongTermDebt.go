package catalogdomain

import (
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"
)

// LongTermDebtProps struct
type LongTermDebtProps struct {
	Value []map[string]float32
}

// LongTermDebt struct
type LongTermDebt struct {
	valueObject domaincore.IValueObject
	props       LongTermDebtProps
	Result      appcore.IResult
}

// GetValue method
func (s LongTermDebt) GetValue() []map[string]float32 {
	return s.props.Value
}

// NewLongTermDebt define LongTermDebt value object
func NewLongTermDebt(props LongTermDebtProps) LongTermDebt {
	//OK case
	return LongTermDebt{
		valueObject: domaincore.NewValueObject(props),
		props:       props,
		Result:      appcore.NewResultOk(),
	}
}
