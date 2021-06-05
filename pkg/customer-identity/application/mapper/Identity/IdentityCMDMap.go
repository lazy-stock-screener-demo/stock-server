package identitymapper

import (
	identitydto "stock-contexts/pkg/customer-identity/application/dto"
	identitydomain "stock-contexts/pkg/customer-identity/domain"
	identityschema "stock-contexts/pkg/customer-identity/infra/repo/schema"
	idcore "stock-contexts/pkg/shared/domain/id"
	// identityschema "stock-contexts/pkg/customer-identity/infra/repo/schema"
)

// EntityMap struct
type EntityMap struct{}

// ToPersistence method
func (e EntityMap) ToPersistence(cus identitydomain.User) *identityschema.User {

	return &identityschema.User{}
}

// ToDTO method
func (e EntityMap) ToDTO(cus identitydomain.User) identitydto.IdentityDTO {
	return identitydto.IdentityDTO{
		AccessToken:  cus.GetAccessToken(),
		RefreshToken: cus.GetRefreshToken(),
	}
}

// ToDomain Implemented
func (e EntityMap) ToDomain(raw *identityschema.User) identitydomain.User {
	userName := identitydomain.NewUserName(
		identitydomain.UserNameProps{
			Value: raw.UserName,
		})
	user := identitydomain.NewUser(
		idcore.NewIDEntity(raw.ID.String()),
		identitydomain.UserProps{
			UserName: userName,
		},
	)
	return user
}

// NewEntityMap method
func NewEntityMap() EntityMap {
	return EntityMap{}
}
