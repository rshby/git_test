package client

import (
	"git_test/adapter/computer"
	"github.com/sirupsen/logrus"
)

type Client struct {
}

// NewClient creates new instance of Client
func NewClient() *Client {
	return &Client{}
}

func (c *Client) InsertLightningConnectorIntoComputer(com computer.Computer) {
	logrus.Info("Client inserts Lightning connector into computer")
	com.InsertIntoLightningPort()
}
