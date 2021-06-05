package identitydomain

import (
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"
)

// UserVIDProps struct
type UserVIDProps struct {
	value string
}

// UserVID struct
type UserVID struct {
	valueObject domaincore.IValueObject
	props       UserVIDProps
	Result      appcore.IResult
}

// GetValue method
func (s UserVID) GetValue() string {
	return s.props.value
}

// NewUserVID constructor
func NewUserVID(props UserVIDProps) UserVID {
	return UserVID{
		valueObject: domaincore.NewValueObject(props),
		props:       props,
		Result:      appcore.NewResultOk(),
	}
}
