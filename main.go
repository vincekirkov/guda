package main

import (
	"github.com/vincekirkov/guda/car"
	"github.com/vincekirkov/guda/engine"
	"github.com/vincekirkov/guda/pedals"
	"github.com/vincekirkov/guda/transmission"
)

func main() {
	pedals.Brake()
	transmission.Park()
	engine.Start()
	engine.Status()
	pedals.Brake()
	transmission.Drive()
	pedals.Gas()
	car.Go()
	car.Stop()
	transmission.Park()
	engine.Stop()
}
