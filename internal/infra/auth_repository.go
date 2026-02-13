package infra

import (
	"server/internal/models"
	"server/internal/repository"
	"server/internal/util"

	"gorm.io/gorm"
)

type AuthRepoGorm struct {
	db *gorm.DB
}

func NewAuthRepoGorm(db *gorm.DB) repository.AuthRepository {
	return &AuthRepoGorm{db: db}
}

func (r *AuthRepoGorm) Signup(data models.User) (models.User, error) {
	var user models.User
	err := r.db.Create(&user).Error
	return user, err
}

func (r *AuthRepoGorm) Login(data models.User) (models.User, error) {
	var user models.User

	err := r.db.Where("email = ?",data.Email).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	
	if !util.CheckPassword(user.Password, data.Password) {
		return models.User{}, err
	}
	return user,nil
}

