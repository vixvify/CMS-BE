package infra

import (
	"server/internal/domain"
	"server/internal/repository"

	"gorm.io/gorm"
)

type BlogRepoGorm struct {
	db *gorm.DB
}

func NewBlogRepoGorm(db *gorm.DB) repository.BlogRepository {
	return &BlogRepoGorm{db: db}
}

func (r *BlogRepoGorm) FindAll() ([]domain.Blog, error) {
	var blogs []domain.Blog
	err := r.db.Find(&blogs).Error
	return blogs, err
}

func (r *BlogRepoGorm) Create(blog domain.Blog) error {
	return r.db.Create(&blog).Error
}
