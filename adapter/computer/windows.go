package computer

import "github.com/sirupsen/logrus"

type Windows struct {
}

// NewWindows creates new instance of Windows
func NewWindows() *Windows {
	return &Windows{}
}

func (w *Windows) InsertIntoUSBPort() {
	logrus.Info("USB connector is plugged into Windows machine.")
}
