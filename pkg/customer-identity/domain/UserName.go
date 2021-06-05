package identitydomain

import (
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"
)

// UserNameProps struct
type UserNameProps struct {
	Value string
}

// UserName struct
type UserName struct {
	valueObject domaincore.IValueObject
	props       UserNameProps
	Result      appcore.IResult
}

// GetValue method
func (s UserName) GetValue() string {
	return s.props.Value
}

func NewUserName(props UserNameProps) UserName {
	return UserName{
		valueObject: domaincore.NewValueObject(props),
		props:       props,
		Result:      appcore.NewResultOk(),
	}
}
