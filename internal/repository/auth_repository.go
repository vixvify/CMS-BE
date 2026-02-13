package repository

import (
	"server/internal/models"
)


type AuthRepository interface {
	Signup(data models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
}