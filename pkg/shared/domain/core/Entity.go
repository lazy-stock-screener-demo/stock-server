package domaincore

import (
	idcore "stock-contexts/pkg/shared/domain/id"

	"github.com/google/go-cmp/cmp"
)

// IEntity interface
type IEntity interface {
	Equals() bool
}

// Entity struct
type Entity struct {
	ID *idcore.UniqueEntityID
}

// Equals method
func (e *Entity) Equals(entity interface{}) bool {
	if entity == nil {
		return false
	}
	if cmp.Equal(e, entity) {
		return true
	}
	if !isEntity(entity) {
		return false
	}
	return e.ID.Identifier.Equals(*entity.(Entity).ID.Identifier)
}

func isEntity(entity interface{}) bool {
	_, ok := entity.(Entity)
	return ok
}

// NewEntity Constructor
func NewEntity(ID *idcore.UniqueEntityID) Entity {
	uid := idcore.NewUniqueEntityID(ID)
	return Entity{
		ID: uid,
	}
}
