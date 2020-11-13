package bridge

import (
	"testing"
	"strings"
)

func TestPrintAPI1(t *testing.T) {
	api1 := Printer1{}
	err := api1.PrintMessage("Hello")

	if err != nil {
		t.Errorf("Error trying to use Printer1 implementation of PrintAPI. Message: %s\n", err.Error())
	}
}

func TestPrintAPI2(t *testing.T) {
	api2 := Printer2{}
	err := api2.PrintMessage("Hello")

	if err != nil {
		expectedErrorMessage := "You need to pass an io.Writer to Printer2"

		if !strings.Contains(err.Error(), expectedErrorMessage) {
			t.Errorf("Error message was not correct.\n Actual: %s\nExpected: %s\n", err.Error(), expectedErrorMessage)
		}
	}

	testWriter := TestWriter{}
	api2 = Printer2{
		Writer: &testWriter,
	}

	expectedMessage := "Hello"
	err = api2.PrintMessage(expectedMessage)

	if err != nil {
		t.Errorf("Error trying to use Printer2 implementation of PrinterAPI. Message: %s\n", err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Fatalf("Printer2 did not write correctly on the io.Writer.\nActual: %s\nExpected: %s\n", testWriter.Msg, expectedMessage)
	}
}

func TestNormalPrinterPrint(t *testing.T) {
	expectedMessage := "Hello io.Writer"

	normal := NormalPrinter{
		Msg: expectedMessage,
		Printer: &Printer1{},
	}

	err := normal.Print()
	if err != nil {
		t.Errorf(err.Error())
	}

	testWriter := TestWriter{}
	normal = NormalPrinter{
		Msg: expectedMessage,
		Printer: &Printer2{
			Writer: &testWriter,
		},
	}

	err = normal.Print()
	if err != nil {
		t.Error(err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Errorf("The expected message on the io.Writer doesn't match actual.\nActual: %s\nExpected: %s\n", testWriter.Msg, expectedMessage)
	}
}

func TestFananPrinterPrint(t *testing.T) {
	passedMessage := "Hello io.Writer"
	expectedMessage := "Message from Fanan: Hello io.Writer"

	fananPrinter := FananPrinter{
		Msg: passedMessage,
		Printer: &Printer1{},
	}

	err := fananPrinter.Print()
	if err != nil {
		t.Errorf(err.Error())
	}

	testWriter := TestWriter{}
	fananPrinter = FananPrinter{
		Msg: passedMessage,
		Printer: &Printer2{
			Writer: &testWriter,
		},
	}

	err = fananPrinter.Print()
	if err != nil {
		t.Error(err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Errorf("The expected message on the io.Writer doesn't match actual.\nActual: %s\nExpected: %s\n", testWriter.Msg,expectedMessage)
	}
}