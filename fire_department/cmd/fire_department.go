package fire_department

import (
	event "emergency-call/Event"
	"fmt"
)

type FireDepartment struct {
}

func NewFireDepartment() *FireDepartment {
	return &FireDepartment{}
}

func (*FireDepartment) Listen(e event.Event) {
	if e == "fire" {
		fmt.Println("the fire department are on their way")
	} else {
		fmt.Println("the fire department didnt answer the " + e + " call")
	}
}
