package database

import (
	"log"
	"server/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "postgresql://neondb_owner:npg_IBw2KUA5RYpz@ep-winter-morning-ahd31e69-pooler.c-3.us-east-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require"

	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, 
		}),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatal("❌ failed to connect database")
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)
	if err != nil {
		log.Fatal("❌ auto migrate failed:", err)
	}

	DB = db
	log.Println("✅ database connected")
}