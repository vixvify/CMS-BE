package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID     `gorm:"type:uuid;primaryKey" json:"id"`
	Email     string
	Password  string
	Username  string
	CreatedAt time.Time `json:"created_at"`
	Blogs     []Blog    `gorm:"foreignKey:UserID"`
}
