package computer

import (
	"git_test/bridge/printer"
	"github.com/sirupsen/logrus"
)

type mac struct {
	printer printer.Printer
}

// NewMac creates new instance of mac. it implements from interface Computer
func NewMac() Computer {
	return &mac{}
}

func (m *mac) Print() {
	logrus.Info("Print request from mac")
	m.printer.PrintFile()
}

func (m *mac) SetPrinter(p printer.Printer) {
	m.printer = p
}
