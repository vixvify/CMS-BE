package repository

import (
	"server/internal/dto"
	"server/internal/models"

	"github.com/google/uuid"
)


type BlogRepository interface {
	FindAll() ([]models.Blog, error)
	Create(blog models.Blog) (models.Blog, error)
	Update(id uuid.UUID,updatedBlog dto.UpdateBlogRequest)  (models.Blog, error)
	Delete(id uuid.UUID)  (error)
}
