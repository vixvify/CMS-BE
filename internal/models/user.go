package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Email     string 	`json:"email"`
	Password  string 	`json:"password"`
	Username  string 	`json:"username"`
	CreatedAt time.Time `json:"created_at"`
	Blogs     []Blog    `gorm:"foreignKey:UserID"`
}
