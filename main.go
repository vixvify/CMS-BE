package main

import (
	"server/config"
	"server/route"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()
	api := r.Group("/api")
	route.RegisterAPIRoutes(api)

	r.Run(":8080")
}
