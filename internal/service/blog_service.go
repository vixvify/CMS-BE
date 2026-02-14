package service

import (
	"errors"
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

func (s *BlogService) GetBlogByID(id uuid.UUID) (models.Blog, error) {
	return s.repo.FindByID(id)
}

func (s *BlogService) CreateBlog(req dto.CreateBlogRequest,userID uuid.UUID) (models.Blog, error)  {
	blog := models.Blog{
		ID: uuid.New(),
		Title:   req.Title,
		Content: req.Content,
		Author:  req.Author,
		CreatedAt: time.Now(),
		UserID: userID,
	}
	return s.repo.Create(blog)
}

func (s *BlogService) UpdateBlog(id uuid.UUID, req dto.UpdateBlogRequest,userID uuid.UUID) (models.Blog, error)  {
	blogbyid, err := s.repo.FindByID(id)
	if err != nil {
		return models.Blog{}, err
	}

	if userID != blogbyid.UserID {
		return models.Blog{} ,errors.New("forbidden")
	}
	
	blog := models.Blog{
		Title: req.Title,
		Content: req.Content,
		Author: req.Author,
	}
	return s.repo.Update(id,blog)
}

func (s *BlogService) DeleteBlog(id uuid.UUID,userID uuid.UUID) (error)  {
	blogbyid, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if userID != blogbyid.UserID {
		return errors.New("forbidden")
	}
	
	return s.repo.Delete(id)
}
