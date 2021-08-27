package transmission

import (
	"fmt"

	"github.com/vincekirkov/guda/pedals"
)

var state string

func Park() {
	if pedals.Status() == "brake" {
		state = "P"
	} else {
		fmt.Println("Please press the brake pedal")
	}
}

func Drive() {
	if pedals.Status() == "brake" {
		state = "D"
	} else {
		fmt.Println("Please press the brake pedal")
	}
}

func Reverse() {
	if pedals.Status() == "brake" {
		state = "R"
	} else {
		fmt.Println("Please press the brake pedal")
	}
}

func Neutural() {
	if pedals.Status() == "brake" {
		state = "N"
	} else {
		fmt.Println("Please press the brake pedal")
	}
}
func Status() string {
	return state
}
