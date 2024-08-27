package main

import (
	"git_test/logger"
	"github.com/sirupsen/logrus"
)

func init() {
	logger.SetupLogger()
}

func main() {
	logrus.Info("start app")
}
