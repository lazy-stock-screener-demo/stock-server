package customerselfdomain

import (
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"

	"golang.org/x/crypto/bcrypt"
)

// CustomerPWDProps struct
type CustomerPWDProps struct {
	Value  string
	Hashed bool
}

// CustomerPWD struct
type CustomerPWD struct {
	valueObject domaincore.IValueObject
	props       CustomerPWDProps
	Result      appcore.IResult
}

// GetValue method
func (s CustomerPWD) GetValue() string {
	return s.props.Value
}

func (s CustomerPWD) ComparePassword(plaintextPWD string) bool {
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

// NewCustomerPWD constructor
func NewCustomerPWD(props CustomerPWDProps) CustomerPWD {
	guard := appcore.NewGurad()
	if validation := guard.AgainstAtLeast(8, props.Value); validation.Passed == false {
		return CustomerPWD{
			valueObject: domaincore.NewValueObject(props),
			props:       props,
			Result: appcore.NewResultFailed(
				appcore.ErrProps{
					Error: validation.Message,
				},
			),
		}
	}
	// OK case
	return CustomerPWD{
		valueObject: domaincore.NewValueObject(props),
		props:       props,
		Result:      appcore.NewResultOk(),
	}
}
