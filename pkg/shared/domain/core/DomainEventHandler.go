package domaincore

import (
	"fmt"
	"runtime"
	domainevent "stock-contexts/pkg/shared/domain/event"
	idcore "stock-contexts/pkg/shared/domain/id"
)

type IAggregate interface {
	GetUniqueEntityID() idcore.UniqueEntityID
	GetRootEventHandler() DomainEventHandler
}

// DomainEventHandler of entities
type DomainEventHandler struct {
	// Entity      Entity
	eventsQueue []domainevent.IDomainEvent
}

// // GetUniqueEntityID can be derived from this.
// func (ag *AggregateRoot) GetUniqueEntityID() idcore.UniqueEntityID {
// 	return *ag.Entity.ID
// }

// ReadEventsQueue deine
func (ag DomainEventHandler) ReadEventsQueue() []domainevent.IDomainEvent {
	return ag.eventsQueue
}

// AppendEvents Append
func (ag *DomainEventHandler) AppendEvents(e domainevent.IDomainEvent) {
	ag.eventsQueue = append(ag.eventsQueue, e)
	ag.saveEvents(e)
}

// RemoveAllEvents Remove all events
func (ag DomainEventHandler) RemoveAllEvents() {
	for i := range ag.eventsQueue {
		ag.eventsQueue[i] = ag.eventsQueue[len(ag.eventsQueue)-1]
		ag.eventsQueue = ag.eventsQueue[:len(ag.eventsQueue)-1]
	}
}

func (ag *DomainEventHandler) saveEvents(e domainevent.IDomainEvent) {
	// TODO: Connect to event store
	_, file, no, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("called from %s#%d\n", file, no)
		fmt.Println("[Domain Event Created]")
	}
}

// NewDomainEventHandler for constructing AggregateRoot
func NewDomainEventHandler() DomainEventHandler {
	return DomainEventHandler{
		// Entity:      NewEntity(id),
		eventsQueue: []domainevent.IDomainEvent{},
	}
}

func DispatchEvents(id *idcore.UniqueEntityID) {
	m := NewEventMediator()
	m.DispatchForAggregate(id)
}
