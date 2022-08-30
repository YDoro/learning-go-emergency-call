package main

import (
	citizen "emergency-call/citizen/cmd"
	fireDepartment "emergency-call/fire_department/cmd"
	police "emergency-call/police/cmd"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	kane := citizen.Citizen{
		Name: "Kayne",
	}

	kane.AddListener(police.NewPolice())
	kane.AddListener(fireDepartment.NewFireDepartment())

	kane.Emit("fire")
	time.Sleep(3 * time.Second)
	kane.Emit("police")
	kane.Emit("fire")

	listen()
}

func listen() {
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
	fmt.Println("stopped")
}
