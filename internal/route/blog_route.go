package route

import (
	"server/internal/handler"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup, h *handler.BlogHandler) {
	r.Use(middleware.RateLimitMiddleware())
	r.GET("", h.GetBlog)
	r.POST("", h.CreateBlog)
	r.PUT("",h.UpdateBlog)
	r.DELETE("",h.DeleteBlog)
}