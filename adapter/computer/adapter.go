package computer

import "github.com/sirupsen/logrus"

type windowsAdapter struct {
	machine *Windows
}

// NewWindowsAdapter creates new instance of windowsAdapter. it implements from interface Computer
func NewWindowsAdapter(machine *Windows) Computer {
	return &windowsAdapter{}
}

// InsertIntoLightningPort plugs windwos into lightning
func (w *windowsAdapter) InsertIntoLightningPort() {
	logrus.Info("Adapter converts Lightning signal to USB")
	w.machine.InsertIntoUSBPort()
}
