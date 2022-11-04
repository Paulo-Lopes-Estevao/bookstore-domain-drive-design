package event

import (
	"bookstore/domain/entities/person/event/handler"
	"bookstore/domain/event"
	"time"
)

func init() {
	event.Register("Send email when customer is created to confirm account", handler.SendEmailWhenCustomerIsCreatedToConfirmAccount())
}

type CustomerCreatedEvent struct {
	event.EventType
}

func NewCustumerCreatedEvent(EventData any) {
	custumerEvent := CustomerCreatedEvent{}
	custumerEvent.DataTimeOccurred = time.Now()
	custumerEvent.EventData = EventData
}

func (customerCreatedEvent *CustomerCreatedEvent) DispatchSendEmailWhenCustomerIsCreatedToConfirmAccount() {
	typeData := new(event.EventType)
	typeData.EventData = customerCreatedEvent.EventData
	event.Dispatch("Send email when customer is created to confirm account", typeData)
}
