package route

import (
	"server/controller"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(router *gin.RouterGroup) {

	router.GET("/users", controller.GetUsers)
	router.POST("/users", controller.CreateUser)
}