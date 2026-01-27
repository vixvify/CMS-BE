package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"server/internal/handler"
	"server/internal/infra"
	"server/internal/service"
)

func main() {
dsn := "host=localhost user=postgres password=1234 dbname=test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	blogRepo := infra.NewBlogRepoGorm(db)
	blogService := service.NewBlogService(blogRepo)
	blogHandler := handler.NewBlogHandler(blogService)

	r := gin.Default()
	r.GET("/blog", blogHandler.GetBlog)
	r.POST("/blog", blogHandler.CreateBlog)

	r.Run(":8080")

}
