package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateBlogRequest struct {
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Author  string    `json:"author"`
	UserID  uuid.UUID `json:"user_id"`
}

type UpdateBlogRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type BlogResponse struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}
