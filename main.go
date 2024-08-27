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

	app.GET("/ping", func(c *gin.Context) {
		statusCode := http.StatusOK
		c.Status(statusCode)
		c.JSON(statusCode, gin.H{
			"status_code": statusCode,
			"message":     "pong!",
		})
	})

	// satu

	if err := server.ListenAndServe(); err != nil {
		logrus.Fatal(err)
	}
}
