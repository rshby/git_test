package printer

import "github.com/sirupsen/logrus"

type epson struct {
}

// NewEpson creates new intstance epson. it implements from interface Printer
func NewEpson() Printer {
	return &epson{}
}

// PrintFile prints file
func (e *epson) PrintFile() {
	logrus.Info("Printing by Epson Printer")
}
