package mapper

import (
	"server/internal/dto"
	"server/internal/models"
)

func ToUserResponse(t models.User) dto.UserResponse {
	return dto.UserResponse{
		Username: t.Username,
		Email:    t.Email,
	}
}
