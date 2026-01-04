package handlers

import (
	"net/http"

	"github.com/barbodimani81/Velooa/internal/models"
	"github.com/labstack/echo/v4"
)

type AddToCartInput struct {
	VariantID uint `json:"variant_id" validate:"required"`
	Quantity  int  `json:"quantity" validate:"required,min=1"`
}

// GetCart godoc
// @Summary Get current cart
// @Description Retrieves the cart for the logged-in user or guest (via ID)
// @Tags cart
// @Produce json
// @Param guest_cart_id query int false "Guest Cart ID (if not logged in)"
// @Success 200 {object} models.Cart
// @Router /cart [get]
func (h *Handler) GetCart(c echo.Context) error {
	return c.JSON(http.StatusOK, models.Cart{})
}

// AddToCart godoc
// @Summary Add item to cart
// @Description Adds a product variant to the cart
// @Tags cart
// @Accept json
// @Produce json
// @Param guest_cart_id query int false "Guest Cart ID"
// @Param input body AddToCartInput true "Item details"
// @Success 200 {object} models.Cart
// @Router /cart/items [post]
func (h *Handler) AddToCart(c echo.Context) error {
	var input AddToCartInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	return c.JSON(http.StatusOK, models.Cart{})
}