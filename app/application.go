package app

import (
	"tradeAPI/logger"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	urlPatterns()
	logger.Info("about to start the application V 1.0.1 !")
	router.Run(":9091")
}
