package handler

import (
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/response"
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
		response.Internal(c, err.Error())
		return
	}
	response.OK(c, mapper.ToBlogResponseList(blogs))
}

func (h *BlogHandler) GetBlogByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	blog, err := h.service.GetBlogByID(id)
	if err != nil {
		response.Internal(c, err.Error())
		return
	}
	response.OK(c, mapper.ToBlogResponse(blog))
}

func (h *BlogHandler) CreateBlog(c *gin.Context) {
	var blog dto.CreateBlogRequest

	if err := c.ShouldBindJSON(&blog); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	userIDStr, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "invalid user id")
		return
	}

	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		response.BadRequest(c, "invalid user id")
		return
	}

	created, err := h.service.CreateBlog(blog, userID)
	if err != nil {
		response.Internal(c, err.Error())
		return
	}
	response.Created(c, mapper.ToBlogResponse(created))
}

func (h *BlogHandler) UpdateBlog(c *gin.Context) {
	var updatedblog dto.UpdateBlogRequest
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := c.ShouldBindJSON(&updatedblog); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	userIDStr, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "invalid user id")
		return
	}

	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		response.BadRequest(c, "invalid user id")
		return
	}

	updated, err := h.service.UpdateBlog(id, updatedblog, userID)
	if err != nil {
		response.Internal(c, err.Error())
		return
	}
	response.OK(c, mapper.ToBlogResponse(updated))
}

func (h *BlogHandler) DeleteBlog(c *gin.Context) {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	userIDStr, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "invalid user id")
		return
	}

	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		response.BadRequest(c, "invalid user id format")
		return
	}

	if err := h.service.DeleteBlog(id, userID); err != nil {
		response.Internal(c, err.Error())
		return
	}
	response.OK(c, nil)
}
