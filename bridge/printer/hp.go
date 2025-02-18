package printer

import "github.com/sirupsen/logrus"

type hp struct {
}

// NewHp creates new instance of hp. it implements from interface Printer
func NewHp() Printer {
	return &hp{}
}

func (h *hp) PrintFile() {
	logrus.Info("Printing by a HP Printer")
}
