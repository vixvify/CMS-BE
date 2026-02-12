package repository

import (
	"server/internal/models"

	"github.com/google/uuid"
)


type BlogRepository interface {
	FindAll() ([]models.Blog, error)
	Create(blog models.Blog) (models.Blog, error)
	Update(id uuid.UUID,updatedBlog models.Blog)  (models.Blog, error)
	Delete(id uuid.UUID)  (error)
}
