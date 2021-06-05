package domainevent

type IDomainEvent interface {
	GetAggregateID() string
}
