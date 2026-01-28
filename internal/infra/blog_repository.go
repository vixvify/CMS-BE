package infra

import (
	"server/internal/domain"
	"server/internal/repository"

	"github.com/google/uuid"
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

func (r *BlogRepoGorm) Create(blog domain.Blog) (domain.Blog, error) {
	blog.ID = uuid.New() 
	err := r.db.Create(&blog).Error
	return blog, err
}
