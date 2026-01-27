package service

import (
	"server/internal/domain"
	"server/internal/repository"
)

type BlogService struct {
	repo repository.BlogRepository
}

func NewBlogService(r repository.BlogRepository) *BlogService {
	return &BlogService{repo: r}
}

func (s *BlogService) GetBlog() ([]domain.Blog, error) {
	return s.repo.FindAll()
}

func (s *BlogService) CreatBlog(blog domain.Blog) error {
	return s.repo.Create(blog)
}
