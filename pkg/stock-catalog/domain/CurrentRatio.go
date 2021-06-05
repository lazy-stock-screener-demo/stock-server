package catalogdomain

import (
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"
)

// CurrentRatioProps struct
type CurrentRatioProps struct {
	Value []map[string]float64
}

// CurrentRatio struct
type CurrentRatio struct {
	valueObject domaincore.IValueObject
	props       CurrentRatioProps
	Result      appcore.IResult
}

// GetValue method
func (s CurrentRatio) GetValue() []map[string]float64 {
	return s.props.Value
}

// NewCurrentRatio define CurrentRatio value object
func NewCurrentRatio(props CurrentRatioProps) CurrentRatio {
	//OK case
	return CurrentRatio{
		valueObject: domaincore.NewValueObject(props),
		props:       props,
		Result:      appcore.NewResultOk(),
	}
}
