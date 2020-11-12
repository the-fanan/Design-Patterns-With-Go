package abstract_factory

import (
	"testing"
)

func TestCarFactory(t *testing.T) {
	carFactory, err := BuildFactory(CarFactoryType)

	if err != nil {
		t.Fatal(err)
	}
	//test luxury car
	carVehicle, err := carFactory.NewVehicle(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}

	luxuryCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion failed")
	}

	//validate that it was a luxury car returend
	if luxuryCar.NumDoors() != 4 {
		t.Fatal("Wrong car type generated")
	}

	if luxuryCar.(Vehicle).NumWheels() != 4 {
		t.Fatal("Wrong number of wheels generated for luxury car")
	}

	if luxuryCar.(Vehicle).NumSeats() != 5 {
		t.Fatal("Wrong number of seats generated for luxury car")
	}
	

	//test family car
	carVehicle, err = carFactory.NewVehicle(FamilyCarType)
	if err != nil {
		t.Fatal(err)
	}

	familyCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion failed")
	}

	//validate that it was a family car returend
	if familyCar.NumDoors() != 6 {
		t.Fatal("Wrong car type generated")
	}

	if familyCar.(Vehicle).NumWheels() != 4 {
		t.Fatal("Wrong number of wheels generated for luxury car")
	}

	if familyCar.(Vehicle).NumSeats() != 5 {
		t.Fatal("Wrong number of seats generated for luxury car")
	}
}