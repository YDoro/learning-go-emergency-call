package event

type Event string

type Listener interface {
	Listen(event Event)
}
