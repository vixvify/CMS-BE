package controller

import "github.com/gin-gonic/gin"

func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "get users",
	})
}

func CreateUser(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "user created",
	})
}
