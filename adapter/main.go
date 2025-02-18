package main

import (
	"git_test/adapter/client"
	"git_test/adapter/computer"
	"github.com/sirupsen/logrus"
)

func main() {
	// create client
	c := client.NewClient()

	// create computer Mac
	mac := computer.NewMac()

	// insert lightningt into Mac (compatible)
	c.InsertLightningConnectorIntoComputer(mac)

	logrus.Info("-----------------------------")

	// create computer windows
	windows := computer.NewWindows()

	// insert lightning into Windows (not compatible)
	// c.InsertLightningConnectorIntoComputer(windows)

	// we need an adapter to wrap windows connector
	windowsAdapter := computer.NewWindowsAdapter(windows)

	// insert lightning into windows by using Adapter
	c.InsertLightningConnectorIntoComputer(windowsAdapter)
}
