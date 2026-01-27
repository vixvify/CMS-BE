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
dsn := "postgresql://neondb_owner:npg_IBw2KUA5RYpz@ep-winter-morning-ahd31e69-pooler.c-3.us-east-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require"
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
