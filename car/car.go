package car

import (
	"fmt"

	"github.com/vincekirkov/guda/engine"
	"github.com/vincekirkov/guda/pedals"
	"github.com/vincekirkov/guda/transmission"
)

var state = "still"

func Go() {
	if engine.Status() == "running" && pedals.Status() == "gas" && (transmission.Status() == "D" || transmission.Status() == "R") {
		state = "moving"
		fmt.Println("Car is moving")
	} else {
		state = "still"
		fmt.Println(engine.Status(), pedals.Status(), transmission.Status())
		fmt.Println("Engine needs to be started, tranmission needs to be on Drive or Reverse and gas pedal must be pressed")
	}
}

func Stop() {
	pedals.Brake()
	if state == "moving" {
		fmt.Println("Car is stopping")
	} else {
		fmt.Println("Car already stopped")
	}
	state = "still"
}

func Status() string {
	return state
}
