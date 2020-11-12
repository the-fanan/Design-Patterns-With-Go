package abstract_factory

import (
	"testing"
)

func TestCarFactory(t *testing.T) {
	carFactory, err := BuildFactory(CarFactoryType)

	if err != nil {
		t.Fatal(err)
	}

	carVehicle, err := carFactory.NewVehicle(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}

	luxuryCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion failed")
	}

	//validate that it was a sport bike returend
	if luxuryCar.NumDoors() != 4 {
		t.Fatal("Wrong car type generated")
	}
}