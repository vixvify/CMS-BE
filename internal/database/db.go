package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "postgresql://neondb_owner:npg_IBw2KUA5RYpz@ep-winter-morning-ahd31e69-pooler.c-3.us-east-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ failed to connect database")
	}

	DB = db
	log.Println("✅ database connected")
}