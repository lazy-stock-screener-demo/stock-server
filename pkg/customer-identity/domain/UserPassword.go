package identitydomain

import (
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"

	"golang.org/x/crypto/bcrypt"
)

// UserPWDProps struct
type UserPWDProps struct {
	Value  string
	Hashed bool
}

// UserPWD struct
type UserPWD struct {
	valueObject domaincore.IValueObject
	props       UserPWDProps
	Result      appcore.IResult
}

// GetValue method
func (s UserPWD) GetValue() string {
	return s.props.Value
}

func (s UserPWD) ComparePassword(plaintextPWD string) bool {
	// if PWD is hashed
	if s.props.Hashed {
		if err := bcrypt.CompareHashAndPassword([]byte(s.props.Value), []byte(plaintextPWD)); err == nil {
			return true
		}
		return false
	}
	// if it remains plain
	return plaintextPWD == s.props.Value
}

func NewUserPWD(props UserPWDProps) UserPWD {
	return UserPWD{
		valueObject: domaincore.NewValueObject(props),
		props:       props,
		Result:      appcore.NewResultOk(),
	}
}
