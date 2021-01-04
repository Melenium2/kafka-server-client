package events

type Event struct {
	ID int `json:"ID,omitempty"`
	Type string `json:"Type,omitempty"`
}
