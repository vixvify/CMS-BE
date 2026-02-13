package main

import (
	"os"
	"server/internal/database"
	"server/internal/handler"
	"server/internal/infra"
	"server/internal/route"
	"server/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	blogRepo := infra.NewBlogRepoGorm(database.DB)
	blogService := service.NewBlogService(blogRepo)
	blogHandler := handler.NewBlogHandler(blogService)

	authRepo := infra.NewAuthRepoGorm(database.DB)
	authService := service.NewAuthService(authRepo,os.Getenv("JWT_SECRET"),)
	authHandler := handler.NewAuthHandler(authService)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	api := r.Group("/api")
	route.RegisterBlogRoutes(api.Group("/blog"), blogHandler)
	route.RegisterAuthRoutes(api.Group("/auth"), authHandler)

	r.Run(":8080")

}
