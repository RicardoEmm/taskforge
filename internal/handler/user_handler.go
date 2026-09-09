package handler

import (
	"net/http"

	"github.com/RicardoEmm/taskforge/internal/domain/users"
	"github.com/RicardoEmm/taskforge/internal/dto"
	"github.com/RicardoEmm/taskforge/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) FindAll(c *gin.Context) {
	users, err := h.userService.FindAll(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (h *UserHandler) FindById(c *gin.Context) {
	parsedID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.FindById(c.Request.Context(), parsedID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *UserHandler) Create(c *gin.Context) {
	var req dto.UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userService.Create(c.Request.Context(), dto.UserCreateInput{
		FullName: req.FullName,
		Email:    req.Email,
		Role:     users.UserRole(req.Role),
	}); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": "user was created"})
}
