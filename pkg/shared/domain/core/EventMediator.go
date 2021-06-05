package domaincore

import (
	"fmt"
	"reflect"
	domainevent "stock-contexts/pkg/shared/domain/event"
	idcore "stock-contexts/pkg/shared/domain/id"
)

var eventMediator *EventMediator = nil

type handlerFunc func(domainevent.IDomainEvent)

type EventMediatorProps struct{}

// EventMediator define a central commander of domain event
type EventMediator struct {
	aggregateQueue []IAggregate
	Handlers       map[string][]handlerFunc
}

// AppendAggregate define a way to append aggregate
func (e *EventMediator) AppendAggregate(ag IAggregate) {
	if eleInQueue, _ := e.findAggregateByID(ag.GetUniqueEntityID().Identifier); eleInQueue == nil {
		e.aggregateQueue = append(e.aggregateQueue, ag)
	}
}

func (e *EventMediator) ReadAggregateQueue() []IAggregate {
	fmt.Printf("Current Queue %v\n", e.aggregateQueue)
	return e.aggregateQueue
}

// findAggregateByID define a method to read aggregate
func (e *EventMediator) findAggregateByID(id *idcore.Identifier) (IAggregate, error) {
	var ag IAggregate
	for _, ele := range e.aggregateQueue {
		fmt.Printf("Current Element ID in Queue %v\n", ele.GetUniqueEntityID().Identifier.Value)
		if ele.GetUniqueEntityID().Identifier.Equals(*id) == true {
			return ele, nil
		}
	}
	return ag, fmt.Errorf("Can't find any aggregate in row")
}

// RemoveAggregateFromQueue define how to remove aggregate from queue.
func (e *EventMediator) removeAggregateFromQueue(ag IAggregate) {
	for i, ele := range e.aggregateQueue {
		if ele.GetUniqueEntityID().Identifier.Equals(*ag.GetUniqueEntityID().Identifier) == true {
			e.aggregateQueue[i] = e.aggregateQueue[len(e.aggregateQueue)-1]
			e.aggregateQueue = append(e.aggregateQueue[:i], e.aggregateQueue[i+1:]...)
		}
	}
}

func (e *EventMediator) removeAllAggregatesFromQueue() {
	for i := range e.aggregateQueue {
		e.aggregateQueue[i] = e.aggregateQueue[len(e.aggregateQueue)-1]
		e.aggregateQueue = e.aggregateQueue[:len(e.aggregateQueue)-1]
	}
}

func (e *EventMediator) publishEventsInAggregate(ag IAggregate) {
	for _, event := range ag.GetRootEventHandler().eventsQueue {
		e.publish(event)
	}
}

// DispatchForAggregate publish events in event queue.
func (e *EventMediator) DispatchForAggregate(id *idcore.UniqueEntityID) {
	if ag, err := e.findAggregateByID(id.Identifier); err == nil {
		e.publishEventsInAggregate(ag)
		fmt.Printf("Dispatch all events in Aggregate %v\n", reflect.TypeOf(ag).Name())
		ag.GetRootEventHandler().RemoveAllEvents()
		e.removeAggregateFromQueue(ag)
	} else {
		fmt.Printf("Aggregate %v not found.\n", id.Identifier.Value)
	}
}

// Publish dispatch handler function when receving an event
func (e *EventMediator) publish(event domainevent.IDomainEvent) {
	eventName := reflect.TypeOf(event).Name()
	if _, ok := e.Handlers[eventName]; ok {
		for _, handler := range e.Handlers[eventName] {
			handler(event)
		}
	}
}

// Register define a way to
func (e *EventMediator) Register(handler handlerFunc, eventName string) {
	e.Handlers[eventName] = append(e.Handlers[eventName], handler)
}

// NewEventMediator construct a new event mediator
func NewEventMediator() *EventMediator {
	if eventMediator == nil {
		eventMediator = &EventMediator{
			aggregateQueue: []IAggregate{},
			Handlers:       map[string][]handlerFunc{},
		}
		return eventMediator
	}
	return eventMediator
}
