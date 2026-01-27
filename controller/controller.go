package controller

import (
	"server/service"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	result := service.GetAllUsers()
	c.JSON(200, gin.H{
		"data": result,
	})
}

func CreateUser(c *gin.Context) {
	result := service.CreateNewUser()
	c.JSON(201, gin.H{
		"data": result,
	})
}
