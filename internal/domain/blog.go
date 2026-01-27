package domain

import "time"

type Blog struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string
	Content   string
	Author    string
	CreatedAt time.Time `gorm:"type:timestamp"`
	UserID    uint
}
