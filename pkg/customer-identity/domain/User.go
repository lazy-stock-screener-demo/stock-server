package identitydomain

import (
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"
	idcore "stock-contexts/pkg/shared/domain/id"
)

// UserProps as props struct
type UserProps struct {
	UserName
	UserID
	AccessToken  string
	RefreshToken string
}

// User stuct
type User struct {
	agEventHandler domaincore.DomainEventHandler
	props          UserProps
	Entity         domaincore.Entity
	Result         appcore.IResult
}

// GetUserID Method
func (u *User) GetUserID() UserID {
	return NewUserID(u.Entity.ID)
}

// GetAccessToken define a method to get accessToken
func (u *User) GetAccessToken() string {
	return u.props.AccessToken
}

// GetRefreshToken define a method to get refreshtoken
func (u *User) GetRefreshToken() string {
	return u.props.RefreshToken
}

// SetAccessToken Method
func (u *User) SetAccessToken(at string, rt string) {
	u.props.AccessToken = at
	u.props.RefreshToken = rt
}

// GetAccessToken define a method to get accessToken
func (u *User) GetUserName() UserName {
	return u.props.UserName
}

// NewUser to create
func NewUser(ID *idcore.UniqueEntityID, props UserProps) User {
	// TODO: Handle creation failed case
	// Handle OK
	user := User{
		agEventHandler: domaincore.NewDomainEventHandler(),
		Entity:         domaincore.NewEntity(ID),
		props:          props,
		Result:         appcore.NewResultOk(),
	}
	// m := domaincore.NewEventMediator()
	// User.ag.AppendEvents()
	// m.AppendAggregate()
	return user
}
