package events

type MessageEvent struct {
	Event
	Message string `json:"Message,omitempty"`
}

func (m MessageEvent) Process() error {
	return nil
}

func NewMessageEvent(message string) MessageEvent {
	return MessageEvent{
		Event: Event{Type: "send"},
		Message: message,
	}
}