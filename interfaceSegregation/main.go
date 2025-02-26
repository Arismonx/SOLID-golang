package main

import "fmt"

//interface
type (
	PowerDevice interface {
		TurnOn()
		TurnOff()
	}
	Printer interface {
		PowerDevice
		Print(data string)
	}
	Scanner interface {
		PowerDevice
		Scan() string
	}
)

//struct
type (
	PrinterOnly     struct{}
	ScannerOnly     struct{}
	AllInOnePrinter struct{}
)

func (p PrinterOnly) TurnOn()           { fmt.Println("PrinterOnly: ON") }
func (p PrinterOnly) TurnOff()          { fmt.Println("PrinterOnly: OFF") }
func (p PrinterOnly) Print(data string) { fmt.Println("PrinterOnly Show data :", data) }

func (s ScannerOnly) TurnOn()      { fmt.Println("ScannerOnly: ON") }
func (s ScannerOnly) TurnOff()     { fmt.Println("ScannerOnly: OFF") }
func (s ScannerOnly) Scan() string { return "ScannerOnly Scannner!" }

func (a AllInOnePrinter) TurnOn()           { fmt.Println("AllInOnePrinter: ON") }
func (a AllInOnePrinter) TurnOff()          { fmt.Println("AllInOnePrinter: OFF") }
func (a AllInOnePrinter) Print(data string) { fmt.Println("AllInOnePrinter Show data:", data) }
func (a AllInOnePrinter) Scan() string      { return ("AllInOnePrinter Scanner!") }

func main() {
	printer := PrinterOnly{}
	scanner := ScannerOnly{}
	all := AllInOnePrinter{}

	printer.TurnOn()
	printer.TurnOff()
	printer.Print("Hello World")

	fmt.Println()

	scanner.TurnOn()
	scanner.TurnOff()
	fmt.Println(scanner.Scan())

	fmt.Println()

	all.TurnOn()
	all.TurnOff()
	all.Print("KUY")
	fmt.Println(all.Scan())

}
