package adapter

import (
	"testing"
	"fmt"
)

func TestAdapter(t *testing.T) {
	msg := "Hello World!"

	adapter := PrinterAdapter{
		OldPrinter: &FananLegacyPrinter{},
		Msg: msg,
	}

	returnedMsg := adapter.PrintStored()

	if returnedMsg != fmt.Sprintf("Legacy Printer: Adapter: %s\n", msg) {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}

	adapter2 := PrinterAdapter{
		Msg: msg,
	}

	returnedMsg = adapter2.PrintStored()

	if returnedMsg != msg {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}
}