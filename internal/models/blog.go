package models

import (
	"time"

	"github.com/google/uuid"
)


type Blog struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Title     string 	`json:"title"`
	Content   string	`json:"content"`
	Author    string	`json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UserID    uuid.UUID	`json:"user_id"`
}
