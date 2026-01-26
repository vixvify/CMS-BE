package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Email     string
	Password  string
	Username  string
	CreatedAt time.Time `gorm:"type:timestamp"`
	Blogs     []Blog    `gorm:"foreignKey:UserID"`
}
