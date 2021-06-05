package idcore

import (
	"strconv"
)

// Identifier is used for unique entity id
type Identifier struct {
	Value string
}

// NewIdentifier as Constructor
func NewIdentifier(Value string) *Identifier {
	return &Identifier{Value}
}

// Equals Method
func (i *Identifier) Equals(id Identifier) bool {
	return id.Value == i.Value
}

// ToString Method
func (i *Identifier) ToString() string {
	return strconv.Quote(i.Value)
}

func (i *Identifier) toValue() string {
	return i.Value
}
