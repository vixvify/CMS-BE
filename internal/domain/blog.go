package domain

import (
	"time"

	"github.com/google/uuid"
)


type Blog struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Title     string
	Content   string
	Author    string
	CreatedAt time.Time `json:"created_at"`
	UserID    uint
}
