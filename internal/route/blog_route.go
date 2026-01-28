package route

import (
	"server/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup, h *handler.BlogHandler) {
	r.GET("", h.GetBlog)
	r.POST("", h.CreateBlog)
}