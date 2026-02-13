package infra

import (
	"server/internal/models"
	"server/internal/repository"

	"gorm.io/gorm"
)

type AuthRepoGorm struct {
	db *gorm.DB
}

func NewAuthRepoGorm(db *gorm.DB) repository.AuthRepository {
	return &AuthRepoGorm{db: db}
}

func (r *AuthRepoGorm) Signup(data models.User) (models.User, error) {
	err := r.db.Create(&data).Error
	return data, err
}

func (r *AuthRepoGorm) FindByEmail(email string) (models.User, error) {
	var user models.User

	err := r.db.Where("email = ?",email).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user,nil
}

