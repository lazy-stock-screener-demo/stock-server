package identitydomain

import (
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"
)

// UserViewProps struct
type UserViewProps struct {
	UserVID
	UserName
}

// UserView struct
type UserView struct {
	valueObject domaincore.IValueObject
	props       UserViewProps
	Result      appcore.IResult
}

// GetUserName method
func (c UserView) GetUserName() UserName {
	return c.props.UserName
}

// NewUserView define constructor
func NewUserView(props UserViewProps) *UserView {
	return &UserView{
		valueObject: domaincore.NewValueObject(props),
		props:       props,
		Result:      appcore.NewResultOk(),
	}
}
