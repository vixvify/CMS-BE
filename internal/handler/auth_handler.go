package handler

import (
	"net/http"
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/response"
	"server/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) Signup(c *gin.Context) {
	var data dto.SignupRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	created, err := h.service.Signup(data)
	if err != nil {
		response.Internal(c, err.Error())
		return
	}
	response.OK(c, mapper.ToUserResponse(created))

}

func (h *AuthHandler) Login(c *gin.Context) {
	var data dto.LoginRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	_, token, err := h.service.Login(data)
	if err != nil {
		response.Internal(c, err.Error())
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "access_token",
		Value:    token,
		Path:     "/",
		Domain:   "blog-app-go.onrender.com",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})

	response.OK(c, nil)
}

func (h *AuthHandler) Me(c *gin.Context) {
	cookie, err := c.Request.Cookie("access_token")
	if err != nil {
		response.Unauthorized(c, "Unauthorized")
		return
	}

	user, err := h.service.Me(cookie.Value)
	if err != nil {
		response.Unauthorized(c, err.Error())
		return
	}

	response.OK(c, mapper.ToUserResponse(user))
}

func (h *AuthHandler) Logout(c *gin.Context) {

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		Domain:   "blog-app-go.onrender.com",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})

	response.OK(c, nil)
}
