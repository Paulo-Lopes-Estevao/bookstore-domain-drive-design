package event

import (
	"bookstore/domain/aggregate/customer/event/handler"
	"bookstore/domain/event"
	"time"
)

func init() {
	event.Register("Send Email When Customer Is Created Handler", handler.SendEmailWhenCustomerIsCreatedHandler())
}

type CustomerCreatedEvent struct {
	*event.EventType
}

func NewCustumerCreatedEvent(EventData any) {
	custumerEvent := CustomerCreatedEvent{}
	custumerEvent.DataTimeOccurred = time.Now()
	custumerEvent.EventData = EventData
}

func (customerCreatedEvent *CustomerCreatedEvent) SendEmailWhenCustomerIsCreatedDispatch() {
	typeData := new(event.EventType)
	typeData.EventData = customerCreatedEvent.EventData
	event.Dispatch("Send Email When Customer Is Created Handler", typeData)
}
