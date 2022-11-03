package event

type EventDispatcherInterface interface {
	Register(event string, listener func(*EventType))
	Dispatch(event string, m *EventType)
}
