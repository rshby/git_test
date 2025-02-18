package computer

import "github.com/sirupsen/logrus"

type mac struct {
}

// NewMac creates new instance of mac, it implements interface Computer
func NewMac() Computer {
	return &mac{}
}

func (m *mac) InsertIntoLightningPort() {
	logrus.Info("Lightning connector is plugged into Mac machine")
}
