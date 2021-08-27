package engine

import (
	"fmt"

	"github.com/vincekirkov/guda/transmission"
)

var state = "stopped"

func Start() {
	if transmission.Status() == "P" || transmission.Status() == "N" {
		state = "running"
	} else {
		fmt.Println("Car needs to be on Neutural or Park before starting the engine")
	}
	if state == "running" {
		fmt.Println("Engine is started")
	}
}

func Stop() {
	if transmission.Status() == "P" || transmission.Status() == "N" {
		state = "stopped"
	} else {
		fmt.Println("Please put car on Neutural or Park in order to stop engine")
	}
	if state == "stopped" {
		fmt.Println("Engine is stopped")
	}
}

func Status() string {
	return state
}
