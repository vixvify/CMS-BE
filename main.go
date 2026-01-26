package main

import (
	"server/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	r.Run(":8080")
}
