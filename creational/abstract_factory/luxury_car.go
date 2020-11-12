package abstract_factory

type LuxuryCar struct {}

func (l *LuxuryCar) NumDoors() int {
	return 4
}

func (l *LuxuryCar) NumWheels() int {
	return 4
}

func (l *LuxuryCar) NumSeats() int {
	return 5
}