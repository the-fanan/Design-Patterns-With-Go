package abstract_factory

import (
	"testing"
)

func TestMotorbikeFactory(t *testing.T) {
	motorbikeFactory, err := BuildFactory(MotorbikeFactoryType)

	if err != nil {
		t.Fatal(err)
	}

	motorbikeVehicle, err := motorbikeFactory.NewVehicle(SportMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}

	sportBike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion failed")
	}

	//validate that it was a sport bike returend
	if sportBike.GetMotorbikeType() != SportMotorbikeType {
		t.Fatal("Wrong motorbike type generated")
	}
}