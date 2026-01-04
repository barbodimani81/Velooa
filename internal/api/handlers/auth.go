package handlers

import (
	"net/http"

	"github.com/barbodimani81/Velooa/internal/models"
	"github.com/labstack/echo/v4"
)

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterInput struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Phone    string `json:"phone" validate:"required"`
}

// Register godoc
// @Summary Register a new user
// @Description Creates a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param input body RegisterInput true "User Registration Info"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string "Bad Request"
// @Router /auth/register [post]
func (h *Handler) Register(c echo.Context) error {
	var input RegisterInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Call h.AuthService.Register(ctx, input.Username, input.Email, ...) here
	
	// Mock response
	return c.JSON(http.StatusCreated, models.User{
		Username: input.Username,
		Email:    input.Email,
	})
}

// Login godoc
// @Summary Login user
// @Description Authenticates user and returns JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param input body LoginInput true "Login Credentials"
// @Success 200 {object} map[string]string "Token Response"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Router /auth/login [post]
func (h *Handler) Login(c echo.Context) error {
	var input LoginInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Call h.AuthService.Login(...) here
	
	return c.JSON(http.StatusOK, map[string]string{"token": "sample-jwt-token"})
}