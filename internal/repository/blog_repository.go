package repository

import "server/internal/domain"

type BlogRepository interface {
	FindAll() ([]domain.Blog, error)
	Create(blog domain.Blog) (domain.Blog, error)
}
