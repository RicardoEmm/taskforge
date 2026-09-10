package handler

import (
	"net/http"

	"github.com/RicardoEmm/taskforge/internal/dto"
	"github.com/RicardoEmm/taskforge/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	tokenService *service.TokenService
	userService  *service.UserService
}

func NewAuthHandler(tokenService *service.TokenService, userService *service.UserService) *AuthHandler {
	return &AuthHandler{tokenService: tokenService, userService: userService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.Register(c.Request.Context(), req.Name, req.Email, req.Password)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.ValidateCredentials(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	access, _ := h.tokenService.GenerateAccessToken(user.ID, user.Email, string(user.Role))
	refresh, _ := h.tokenService.GenerateRefreshToken(user.ID, user.Email, string(user.Role))
	c.JSON(http.StatusOK, gin.H{"access_token": access, "refresh_token": refresh})
}
