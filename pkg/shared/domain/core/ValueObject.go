package domaincore

// IValueObject define a interface if any implemented Equals
type IValueObject interface {
	Equals(interface{}) bool
}

// ValueObject struct
type ValueObject struct {
	Props interface{}
}

// Equals method
func (v *ValueObject) Equals(vo interface{}) bool {
	if vo == nil {
		return false
	}
	if v.Props == nil {
		return false
	}
	return vo == v.Props
}

// NewValueObject constructor
func NewValueObject(p interface{}) *ValueObject {
	return &ValueObject{
		Props: p,
	}
}
