package event

import "time"

type EventType struct {
	DataTimeOccurred time.Time
	EventData        interface{}
}
