package eventhandler

import (
	"fmt"
	"reflect"
	login "stock-contexts/pkg/customer-identity/application/useCase/login"
	customerselfdomain "stock-contexts/pkg/customer-self/domain/core"
	domaincore "stock-contexts/pkg/shared/domain/core"
	domainevent "stock-contexts/pkg/shared/domain/event"
)

// AfterCustomerCreated is a struct
type AfterCustomerCreated struct {
	listeningEvent domainevent.IDomainEvent
	useCase        login.IUseCaseCRUD
}

func (a *AfterCustomerCreated) register() {
	listeningEventName := reflect.TypeOf(a.listeningEvent)
	domaincore.NewEventMediator().Register(a.onCustomerCreated, listeningEventName.Name())
}

func (a *AfterCustomerCreated) onCustomerCreated(event domainevent.IDomainEvent) {
	fmt.Printf("[Customer ID: ]%v", event.GetAggregateID())
	// TODO: Trigger ID could only be domain aggregate ID or aggregate itself.
	// It is not ideal to mix with ReqDTO, define another DTO instead, for example eventDTO.
	a.useCase.Execute(login.ReqDTO{
		UserName: event.GetAggregateID(),
	})
	fmt.Println("[Success Trigger Login UseCase]")
}

// NewAfterCustomerCreated define
func NewAfterCustomerCreated() {
	event := AfterCustomerCreated{
		listeningEvent: customerselfdomain.CustomerCreated{},
		useCase:        login.NewUseCaseCRUD(),
	}
	event.register()
}
