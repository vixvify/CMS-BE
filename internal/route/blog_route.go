package route

import (
	"server/internal/handler"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterBlogRoutes(r *gin.RouterGroup, h *handler.BlogHandler, jwtSecret string) {
	r.Use(middleware.RateLimitMiddleware())
	r.GET("", h.GetBlog)
	r.POST("", middleware.JWTAuth(jwtSecret), h.CreateBlog)
	r.PUT("/:id", middleware.JWTAuth(jwtSecret),h.UpdateBlog)
	r.DELETE("/:id", middleware.JWTAuth(jwtSecret),h.DeleteBlog)
}