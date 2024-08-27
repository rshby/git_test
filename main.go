package main

import (
	"git_test/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func init() {
	logger.SetupLogger()
}

func main() {
	logrus.Info("start app")

	app := gin.Default()
	server := http.Server{
		Addr:    ":4000",
		Handler: app,
	}

	if err := server.ListenAndServe(); err != nil {
		logrus.Fatal(err)
	}
}
