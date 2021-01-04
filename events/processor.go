package events

type Processor interface {
	Process() error
}
