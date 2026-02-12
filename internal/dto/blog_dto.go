package dto

import "github.com/google/uuid"

type CreateBlogRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
	UserID  uuid.UUID `json:"userId"`
}

type UpdateBlogRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
