package idcore

import (
	"github.com/gofrs/uuid"
)

// UniqueEntityID is used for unique entity id
type UniqueEntityID struct {
	Identifier *Identifier
}

// NewUniqueEntityID as Constructor to build new uniqueEntityID
func NewUniqueEntityID(id *UniqueEntityID) *UniqueEntityID {
	return &UniqueEntityID{
		Identifier: &Identifier{
			Value: GetID(id),
		},
	}
}

func NewIDEntity(id string) *UniqueEntityID {
	return &UniqueEntityID{
		Identifier: &Identifier{
			Value: id,
		},
	}
}

// GetID use id if id exist or generate new uuid
func GetID(id *UniqueEntityID) string {
	if id != nil {
		return id.Identifier.Value
	}
	uuid, _ := uuid.NewV4()
	return uuid.String()
}

// // GetID use id if id exist or generate new uuid
// func GetID(id interface{}) string {
// 	fmt.Println(fmt.Sprintf("UniqueEntity ID: %v", id))
// 	switch id.(type) {
// 	case string:
// 		return id.(string)
// 	case int:
// 		return fmt.Sprint(id.(int))
// 	case UniqueEntityID:
// 		return id.(UniqueEntityID).Identifier.value
// 	case nil:
// 		uuid, _ := uuid.NewV4()
// 		return uuid.String()
// 	default:
// 		uuid, _ := uuid.NewV4()
// 		return uuid.String()
// 	}
// }
