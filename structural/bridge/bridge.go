package bridge

import (
	"errors"
	"io"
	"fmt"
)

type PrinterAPI interface {
	PrintMessage(string) error
}
//structs that implement PrinterAPI
type Printer1 struct{}

func (p *Printer1) PrintMessage(msg string) error {
	fmt.Printf("%s\n", msg)
	return nil
}

type Printer2 struct {
	Writer io.Writer
}

func (p *Printer2) PrintMessage(msg string) error {
	if p.Writer == nil {
		return errors.New("You need to pass an io.Writer to Printer2")
	}
	fmt.Fprintf(p.Writer, "%s", msg)
	return nil
}	

type TestWriter struct {
	Msg string
}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 {
		t.Msg = string(p)
		return n, nil
	}
	err = errors.New("Empty content received")
	return
}

//Structs that use printer API
type PrinterAbstraction interface {
	Print() error
}

type NormalPrinter struct {
	Msg string 
	Printer PrinterAPI
}

func (n *NormalPrinter) Print() error {
	n.Printer.PrintMessage(n.Msg)
	return nil
}

type FananPrinter struct {
	Msg string
	Printer PrinterAPI
}

func (f *FananPrinter) Print() error {
	f.Printer.PrintMessage(fmt.Sprintf("Message from Fanan: %s", f.Msg))
	return nil
}