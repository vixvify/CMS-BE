package handler

import (
	"net/http"
	"server/internal/dto"
	"server/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BlogHandler struct {
	service *service.BlogService
}

func NewBlogHandler(s *service.BlogService) *BlogHandler {
	return &BlogHandler{service: s}
}

func (h *BlogHandler) GetBlog(c *gin.Context) {
	blogs, err := h.service.GetBlog()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": blogs,
		"status": "success",
		"statusCode": 200,
	})
}

func (h *BlogHandler) CreateBlog(c *gin.Context) {
	var blog dto.CreateBlogRequest

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userIDStr, exists := c.Get("userID")
    if !exists {
        c.JSON(401, gin.H{"error": "unauthorized"})
        return
    }

    userID, err := uuid.Parse(userIDStr.(string))
    if err != nil {
        c.JSON(400, gin.H{"error": "invalid user id"})
        return
    }

	created, err := h.service.CreateBlog(blog,userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, created)
}

func (h *BlogHandler) UpdateBlog(c *gin.Context) {
	var updatedblog dto.UpdateBlogRequest
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
    	c.JSON(400, gin.H{"error": "invalid id"})
    	return
	}

	if err := c.ShouldBindJSON(&updatedblog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userIDStr, exists := c.Get("userID")
    if !exists {
        c.JSON(401, gin.H{"error": "unauthorized"})
        return
    }

    userID, err := uuid.Parse(userIDStr.(string))
    if err != nil {
        c.JSON(400, gin.H{"error": "invalid user id"})
        return
    }

	updated, err := h.service.UpdateBlog(id,updatedblog,userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, updated)
}	

func (h *BlogHandler) DeleteBlog(c *gin.Context) {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
    	c.JSON(400, gin.H{"error": "invalid id"})
    	return
	}

	userIDStr, exists := c.Get("userID")
    if !exists {
        c.JSON(401, gin.H{"error": "unauthorized"})
        return
    }

    userID, err := uuid.Parse(userIDStr.(string))
    if err != nil {
        c.JSON(400, gin.H{"error": "invalid user id"})
        return
    }

	if err := h.service.DeleteBlog(id,userID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(204)
}	
