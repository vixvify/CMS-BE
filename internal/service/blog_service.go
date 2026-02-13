package service

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repository"
	"time"

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

func (s *BlogService) CreateBlog(req dto.CreateBlogRequest) (models.Blog, error)  {
	blog := models.Blog{
		ID: uuid.New(),
		Title:   req.Title,
		Content: req.Content,
		Author:  req.Author,
		CreatedAt: time.Now(),
		UserID: req.UserID,
	}
	return s.repo.Create(blog)
}

func (s *BlogService) UpdateBlog(id uuid.UUID, req dto.UpdateBlogRequest) (models.Blog, error)  {
	blog := models.Blog{
		Title: req.Title,
		Content: req.Content,
		Author: req.Author,
	}
	return s.repo.Update(id,blog)
}

func (s *BlogService) DeleteBlog(id uuid.UUID) (error)  {
	return s.repo.Delete(id)
}
