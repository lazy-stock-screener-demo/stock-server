package domaincore

import (
	"fmt"
	domainevent "stock-contexts/pkg/shared/domain/event"
	idcore "stock-contexts/pkg/shared/domain/id"
	"testing"
)

func TestRegisterEvent(t *testing.T) {
	m := NewEventMediator()
	m.Register(func(event domainevent.IDomainEvent) {
	}, "TestEvent")
	fmt.Println(m.Handlers)
	if val, ok := m.Handlers["TestEvent"]; !ok {
		t.Errorf("Expected Handler, received %v", val)
	}
}

type TestEvent struct{}

func (c TestEvent) GetAggregateID() string {
	return ""
}

func NewTestEvent() TestEvent {
	return TestEvent{}
}

func TestDispatchEvent(t *testing.T) {
	// eventhandler.NewAfterCustomerCreated()
	m := NewEventMediator()
	m.Register(func(event domainevent.IDomainEvent) {
		fmt.Println("Handler Dispatched Succeed.")
	}, "TestEvent")
	// fmt.Println(reflect.TypeOf(NewTestEvent()).Name())
	m.publish(NewTestEvent())
	if val, ok := m.Handlers["TestEvent"]; !ok {
		t.Errorf("Expected Handler, received %v", val)
	}
}

type TypeAgRoot struct {
	agEventHandler DomainEventHandler
	Entity         Entity
}

func (s TypeAgRoot) GetUniqueEntityID() idcore.UniqueEntityID {
	return *s.Entity.ID
}

func (s TypeAgRoot) GetRootEventHandler() DomainEventHandler {
	return s.agEventHandler
}

func TestAppendAggregate(t *testing.T) {
	id := &idcore.UniqueEntityID{
		Identifier: &idcore.Identifier{
			Value: "123",
		},
	}
	typeRoot := TypeAgRoot{
		agEventHandler: NewDomainEventHandler(),
		Entity:         NewEntity(id),
	}
	m := NewEventMediator()
	typeRoot.agEventHandler.AppendEvents(NewTestEvent())
	m.AppendAggregate(typeRoot)

	found, err := m.findAggregateByID(typeRoot.GetUniqueEntityID().Identifier)
	fmt.Printf("Found ID is %v\n", found.GetUniqueEntityID().Identifier.Value)
	if err != nil {
		t.Errorf("TestAppendAggregate got wrong with message: %v", err)
	}
}
