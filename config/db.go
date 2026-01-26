package config

import (
	"log"
	"server/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "postgresql://neondb_owner:npg_IBw2KUA5RYpz@ep-winter-morning-ahd31e69-pooler.c-3.us-east-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB error: ", err)
	}

	DB = database
	DB.AutoMigrate(&model.User{}, &model.Blog{})
}
