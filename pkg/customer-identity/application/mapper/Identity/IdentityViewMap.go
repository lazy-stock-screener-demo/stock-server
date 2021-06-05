package identitymapper

import (
	identitydto "stock-contexts/pkg/customer-identity/application/dto"
	identitydomain "stock-contexts/pkg/customer-identity/domain"
	identityschema "stock-contexts/pkg/customer-identity/infra/repo/schema"
)

type ViewMap struct{}

func (v ViewMap) ToDomain(raw *identityschema.User) *identitydomain.UserView {
	return &identitydomain.UserView{}
}

func (v ViewMap) ToDTO(customerView *identitydomain.UserView) identitydto.IdentityDTO {
	return identitydto.IdentityDTO{}
}

func NewViewMap() ViewMap {
	return ViewMap{}
}
