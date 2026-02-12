package service

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repository"

	"github.com/google/uuid"
)

type BlogService struct {
	repo repository.BlogRepository
}

func NewBlogService(r repository.BlogRepository) *BlogService {
	return &BlogService{repo: r}
}

func (s *BlogService) GetBlog() ([]models.Blog, error) {
	return s.repo.FindAll()
}

func (s *BlogService) CreateBlog(blog models.Blog) (models.Blog, error)  {
	return s.repo.Create(blog)
}

func (s *BlogService) UpdateBlog(id uuid.UUID, updatedblog dto.UpdateBlogRequest) (models.Blog, error)  {
	return s.repo.Update(id,updatedblog)
}

func (s *BlogService) DeleteBlog(id uuid.UUID) (error)  {
	return s.repo.Delete(id)
}
