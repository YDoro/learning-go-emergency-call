package citzen

import (
	event "emergency-call/Event"
	"fmt"
)

type Citizen struct {
	Name               string
	EmergencyListeners []event.Listener
}

func (c *Citizen) AddListener(listener event.Listener) {
	c.EmergencyListeners = append(c.EmergencyListeners, listener)

}

func (c *Citizen) Emit(e event.Event) {
	fmt.Println("=====================================")

	fmt.Printf("%v screams %v!\n", c.Name, e)
	for _, listener := range c.EmergencyListeners {
		go listener.Listen(e)
	}
	fmt.Println("=====================================")
}
