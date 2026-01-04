package handlers

import (
	"net/http"

	"github.com/barbodimani81/Velooa/internal/models"
	"github.com/labstack/echo/v4"
)

// ListProducts godoc
// @Summary List products
// @Description Get a paginated list of products, optionally filtered by category
// @Tags products
// @Produce json
// @Param category_id query int false "Category ID"
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {array} models.Product
// @Router /products [get]
func (h *Handler) ListProducts(c echo.Context) error {
	// Logic to call ProductService
	return c.JSON(http.StatusOK, []models.Product{})
}

// GetLook godoc
// @Summary Get a specific Look
// @Description Returns a Look and its associated products
// @Tags looks
// @Produce json
// @Param id path int true "Look ID"
// @Success 200 {object} models.Look
// @Router /looks/{id} [get]
func (h *Handler) GetLook(c echo.Context) error {
	// Logic to call ProductService.GetLook
	return c.JSON(http.StatusOK, models.Look{})
}

// ListCategories godoc
// @Summary List categories
// @Description Get full category tree
// @Tags products
// @Produce json
// @Success 200 {array} models.Category
// @Router /categories [get]
func (h *Handler) ListCategories(c echo.Context) error {
	return c.JSON(http.StatusOK, []models.Category{})
}