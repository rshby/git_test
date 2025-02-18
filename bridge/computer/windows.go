package computer

import (
	"git_test/bridge/printer"
	"github.com/sirupsen/logrus"
)

type windows struct {
	printer printer.Printer
}

// NewWindows creates new instance windows. it implements from interface Computer
func NewWindows() Computer {
	return &windows{}
}

func (w *windows) Print() {
	logrus.Info("Print from windows")
	w.printer.PrintFile()
}

func (w *windows) SetPrinter(p printer.Printer) {
	w.printer = p
}
