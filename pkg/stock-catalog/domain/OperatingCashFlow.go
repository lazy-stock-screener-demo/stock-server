package catalogdomain

import (
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"
)

// OperatingCashFlowProps struct
type OperatingCashFlowProps struct {
	Value []map[string]float32
}

// OperatingCashFlow struct
type OperatingCashFlow struct {
	valueObject domaincore.IValueObject
	props       OperatingCashFlowProps
	Result      appcore.IResult
}

// GetValue method
func (s OperatingCashFlow) GetValue() []map[string]float32 {
	return s.props.Value
}

// NewOperatingCashFlow define OperatingCashFlow value object
func NewOperatingCashFlow(props OperatingCashFlowProps) OperatingCashFlow {
	//OK case
	return OperatingCashFlow{
		valueObject: domaincore.NewValueObject(props),
		props:       props,
		Result:      appcore.NewResultOk(),
	}
}
