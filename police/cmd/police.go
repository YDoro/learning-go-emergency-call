package police

import (
	event "emergency-call/Event"
	"fmt"
)

type Police struct {
}

func (*Police) Listen(e event.Event) {
	if e == "police" {
		fmt.Println("the police are on their way")
	} else {
		fmt.Println("police didnt answer the " + e + " call")
	}
}

func NewPolice() *Police {
	return &Police{}
}
