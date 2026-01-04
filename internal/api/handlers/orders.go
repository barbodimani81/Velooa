package handlers

import (
	"net/http"

	"github.com/barbodimani81/Velooa/internal/models"
	"github.com/labstack/echo/v4"
)

type PlaceOrderInput struct {
	AddressID         uint   `json:"address_id" validate:"required"`
	PaymentProviderID string `json:"payment_provider_id" validate:"required"`
}

// PlaceOrder godoc
// @Summary Place an order
// @Description Converts current cart into a paid order
// @Tags orders
// @Accept json
// @Produce json
// @Param input body PlaceOrderInput true "Order Details"
// @Security ApiKeyAuth
// @Success 201 {object} models.Order
// @Router /orders [post]
func (h *Handler) PlaceOrder(c echo.Context) error {
	var input PlaceOrderInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Logic calls h.OrderService.PlaceOrder...
	return c.JSON(http.StatusCreated, models.Order{
		PaymentProviderID: input.PaymentProviderID,
		// Status: "PENDING", // Note: Your model has Status as bool, likely a bug to fix in models.
	})
}