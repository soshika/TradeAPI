package main

import (
	"tradeAPI/app"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	app.StartApplication()
}
