package customerselfdomain

import (
	"regexp"
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"
)

// CustomerNameProps struct
type CustomerEmailProps struct {
	Value string
}

// CustomerName struct
type CustomerEmail struct {
	valueObject domaincore.IValueObject
	props       CustomerEmailProps
	Result      appcore.IResult
}

// GetValue method
func (s CustomerEmail) GetValue() string {
	return s.props.Value
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)
	// fmt.Printf("%v", re.Match([]byte(email)))
	return re.Match([]byte(email))
}

func NewCustomerEmail(props CustomerEmailProps) CustomerEmail {
	if ok := isValidEmail(props.Value); ok == false {
		return CustomerEmail{
			valueObject: domaincore.NewValueObject(props),
			props:       props,
			Result: appcore.NewResultFailed(
				appcore.ErrProps{
					Error: `Email is not valided.`,
				},
			),
		}
	}
	// OK case
	return CustomerEmail{
		valueObject: domaincore.NewValueObject(props),
		props:       props,
		Result:      appcore.NewResultOk(),
	}
}
