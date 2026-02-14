package infra

import (
	"server/internal/models"
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

func (r *BlogRepoGorm) FindAll() ([]models.Blog, error) {
	var blogs []models.Blog
	err := r.db.Find(&blogs).Error
	return blogs, err
}

func (r *BlogRepoGorm) FindByID(id uuid.UUID) (models.Blog, error) {
	var blog models.Blog
	err := r.db.Where("id = ?",id).First(&blog).Error
	return blog, err
}

func (r *BlogRepoGorm) Create(blog models.Blog) (models.Blog, error) {
	err := r.db.Create(&blog).Error
	return blog, err
}

func (r *BlogRepoGorm) Update(id uuid.UUID,updatedBlog models.Blog) (models.Blog, error) {
	if err := r.db.
		Model(&models.Blog{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"title":   updatedBlog.Title,
			"content": updatedBlog.Content,
			"author":  updatedBlog.Author,
		}).
		Error; err != nil {
		return models.Blog{}, err
	}

	var blog models.Blog
	if err := r.db.First(&blog, "id = ?", id).Error; err != nil {
		return models.Blog{}, err
	}
	
	return blog, nil
}

func (r *BlogRepoGorm) Delete(id uuid.UUID) (error) {
	return r.db.Delete(&models.Blog{}, "id = ?", id).Error
}
