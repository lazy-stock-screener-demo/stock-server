package customerselfdomain

import (
	// customerselfevent "stock-contexts/pkg/customer-self/domain/event"
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"
	idcore "stock-contexts/pkg/shared/domain/id"
	"time"
)

// CustomerProps as props struct
type CustomerProps struct {
	CustomerName
	CustomerPWD
	CustomerEmail
	LicenseID
}

// Customer stuct
type Customer struct {
	agEventHandler domaincore.DomainEventHandler
	props          CustomerProps
	Entity         domaincore.Entity
	Result         appcore.IResult
}

// GetUniqueEntityID can be derived from this.
func (s *Customer) GetUniqueEntityID() idcore.UniqueEntityID {
	return *s.Entity.ID
}

func (s Customer) GetRootEventHandler() domaincore.DomainEventHandler {
	return s.agEventHandler
}

// GetCustomerID Method
func (s *Customer) GetCustomerID() CustomerID {
	return NewCustomerID(s.Entity.ID)
}

func (s *Customer) GetCustomerName() CustomerName {
	return s.props.CustomerName
}

func (s *Customer) GetCustomerPWD() CustomerPWD {
	return s.props.CustomerPWD
}

func (s *Customer) GetEmail() CustomerEmail {
	return s.props.CustomerEmail
}

func (s *Customer) GetLicenseID() LicenseID {
	return s.props.LicenseID
}

// NewCustomer to create
func NewCustomer(ID *idcore.UniqueEntityID, props CustomerProps) Customer {
	// TODO: Handle creation failed case
	// Handle OK
	customer := Customer{
		agEventHandler: domaincore.NewDomainEventHandler(),
		Entity:         domaincore.NewEntity(ID),
		props:          props,
		Result:         appcore.NewResultOk(),
	}
	// Side-Effect: Add Event into Queue
	m := domaincore.NewEventMediator()
	customer.agEventHandler.AppendEvents(
		NewCustomerCreated(
			CustomerCreatedProps{
				timeOccured: time.Now(),
				customer:    customer,
			}))
	m.AppendAggregate(&customer)

	// m.ReadAggregateQueue()
	return customer
}
