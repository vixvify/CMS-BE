package main

import (
	"server/internal/database"
	"server/internal/handler"
	"server/internal/infra"
	"server/internal/route"
	"server/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	blogRepo := infra.NewBlogRepoGorm(database.DB)
	blogService := service.NewBlogService(blogRepo)
	blogHandler := handler.NewBlogHandler(blogService)

	r := gin.Default()
	api := r.Group("/api")
	route.RegisterUserRoutes(api.Group("/blog"), blogHandler)

	r.Run(":8080")

}
