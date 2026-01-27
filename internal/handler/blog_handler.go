package handler

import (
	"net/http"
	"server/internal/domain"
	"server/internal/service"

	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	service *service.BlogService
}

func NewBlogHandler(s *service.BlogService) *BlogHandler {
	return &BlogHandler{service: s}
}

func (h *BlogHandler) GetBlog(c *gin.Context) {
	users, err := h.service.GetBlog()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *BlogHandler) CreateBlog(c *gin.Context) {
	var blog domain.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreatBlog(blog); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, blog)
}
