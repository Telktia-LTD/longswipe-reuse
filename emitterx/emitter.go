package emitterx

import (
	"sync"
)

// EventPayload represents an event with a dynamic data type.
type EventPayload struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

// EventHandler is a function that handles an event.
type EventHandler func(EventPayload)

// EventEmitter is a struct that manages event listeners and emits events.
type EventEmitter struct {
	listeners map[string][]EventHandler
	mu        sync.Mutex
}

// NewEventEmitter creates a new EventEmitter.
func NewEventEmitter() *EventEmitter {
	return &EventEmitter{
		listeners: make(map[string][]EventHandler),
	}
}

// On adds a new listener for an event.
func (e *EventEmitter) On(eventName string, handler EventHandler) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.listeners[eventName] = append(e.listeners[eventName], handler)
}

// Emit emits an event to all registered listeners.
func (e *EventEmitter) Emit(event EventPayload) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if handlers, found := e.listeners[event.Event]; found {
		for _, handler := range handlers {
			go handler(event)
		}
	}
}
