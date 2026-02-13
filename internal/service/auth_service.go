package service

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repository"
	"server/internal/util"
	"time"

	"github.com/google/uuid"
)

type AuthService struct {
	repo repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *AuthService {
	return &AuthService{repo: r}
}

func (s *AuthService) Signup(req dto.SignupRequest) (models.User, error) {
	hashed,err := util.HashPassword(req.Password)
	if err != nil {
		return models.User{}, err
	}
	user := models.User{
		ID: uuid.New(),
		Username: req.Username,
		Email: req.Email,
		Password: hashed,
		CreatedAt: time.Now(),
	}
	return s.repo.Signup(user)
}

func (s *AuthService) Login(req dto.LoginRequest) (models.User, error)  {
	user := models.User{
		Email: req.Email,
		Password: req.Password,
	}
	return s.repo.Login(user)
}