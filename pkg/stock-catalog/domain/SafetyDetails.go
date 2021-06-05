package catalogdomain

import (
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"
)

// SafetyDetailsProps struct
type SafetyDetailsProps struct {
	LongTermDebt
	CurrentRatio
}

// SafetyDetails define
type SafetyDetails struct {
	valueObject domaincore.IValueObject
	props       SafetyDetailsProps
	Result      appcore.IResult
}

func (p SafetyDetails) GetLongTermDebt() LongTermDebt {
	return p.props.LongTermDebt
}

func (p SafetyDetails) GetCurrentRatio() CurrentRatio {
	return p.props.CurrentRatio
}

func NewSafetyDetails(props SafetyDetailsProps) SafetyDetails {
	return SafetyDetails{
		valueObject: domaincore.NewValueObject(props),
		props:       props,
		Result:      appcore.NewResultOk(),
	}
}
