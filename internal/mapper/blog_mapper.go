package mapper

import (
	"server/internal/dto"
	"server/internal/models"
)

func ToBlogResponse(t models.Blog) dto.BlogResponse {
	return dto.BlogResponse{
		ID:        t.ID,
		Title:     t.Title,
		Content:   t.Content,
		Author:    t.Author,
		CreatedAt: t.CreatedAt,
	}
}

func ToBlogResponseList(teams []models.Blog) []dto.BlogResponse {
	out := make([]dto.BlogResponse, 0, len(teams))
	for _, t := range teams {
		out = append(out, ToBlogResponse(t))
	}
	return out
}
