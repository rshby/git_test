package main

import (
	"fmt"
	"git_test/bridge/computer"
	"git_test/bridge/printer"
	"git_test/drivers/logger"
)

func main() {
	logger.SetupLogger()
	
	// create HP Printer
	hpPrinter := printer.NewHp()

	// create Epson Printer
	epsonPrinter := printer.NewEpson()

	// create Mac Computer
	macComputer := computer.NewMac()

	// set printer
	macComputer.SetPrinter(hpPrinter)

	// do print
	macComputer.Print()

	fmt.Println("-------------------------------------------")

	// create Windows Computer
	windowsComputer := computer.NewWindows()

	// set printer to HP
	windowsComputer.SetPrinter(hpPrinter)

	// do print with Windows and HP
	windowsComputer.Print()

	fmt.Println("-------------------------------------------")

	// re-set Windows with Epson Printer
	windowsComputer.SetPrinter(epsonPrinter)

	// do print with Windows and Epson
	windowsComputer.Print()
}
