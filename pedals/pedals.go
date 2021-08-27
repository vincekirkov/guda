package pedals

var state string

func Gas() {
	state = "gas"
}

func Brake() {
	state = "brake"
}

func Status() string {
	return state
}
