package customerselfdomain

import (
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"
)

// LicenseNameProps struct
type LicenseNameProps struct {
	Value string
}

// LicenseName struct
type LicenseName struct {
	valueObject domaincore.IValueObject
	props       LicenseNameProps
	Result      appcore.IResult
}

// GetValue method
func (s LicenseName) GetValue() string {
	return s.props.Value
}

// NewLicenseName constructor
func NewLicenseName(props LicenseNameProps) LicenseName {
	guard := appcore.NewGurad()
	if validation := guard.AgainstAtMost(22, props.Value); validation.Passed == false {
		return LicenseName{
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
		return LicenseName{
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
	return LicenseName{
		valueObject: domaincore.NewValueObject(props),
		props:       props,
		Result:      appcore.NewResultOk(),
	}
}
