package abstract_factory

import (
	"testing"
)

func TestMotorbikeFactory(t *testing.T) {
	motorbikeFactory, err := BuildFactory(MotorbikeFactoryType)

	if err != nil {
		t.Fatal(err)
	}
	//test sport bike
	motorbikeVehicle, err := motorbikeFactory.NewVehicle(SportMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}

	sportBike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion failed")
	}

	if sportBike.(Vehicle).NumWheels() != 2 {
		t.Fatal("Wrong number of wheels generated for sport motorbike")
	}

	if sportBike.(Vehicle).NumSeats() != 2 {
		t.Fatal("Wrong number of seats generated for sport motorbike")
	}
	//validate that it was a sport bike returend
	if sportBike.GetMotorbikeType() != SportMotorbikeType {
		t.Fatal("Wrong motorbike type generated")
	}

	//test cruise bike
	motorbikeVehicle, err = motorbikeFactory.NewVehicle(CruiseMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}

	cruiseBike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion failed")
	}

	if cruiseBike.(Vehicle).NumWheels() != 2 {
		t.Fatal("Wrong number of wheels generated for cruise motorbike")
	}

	if cruiseBike.(Vehicle).NumSeats() != 2 {
		t.Fatal("Wrong number of seats generated for cruise motorbike")
	}

	//validate that it was a cruise bike returend
	if cruiseBike.GetMotorbikeType() != CruiseMotorbikeType {
		t.Fatal("Wrong motorbike type generated")
	}
}