package event

var listeners = make(map[string]func(*EventType))

// Register adds an event handler for this event
func Register(event string, listener func(*EventType)) {
	listeners[event] = listener
}

// Dispatch sends out an event with the payload
func Dispatch(event string, m *EventType) {
	listeners[event](m)
}
