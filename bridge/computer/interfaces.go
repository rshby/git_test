package computer

import "git_test/bridge/printer"

type Computer interface {
	Print()
	SetPrinter(p printer.Printer)
}
