package handlers

import (
	"github.com/barbodimani81/Velooa/internal/service"
)

// Handler holds dependencies for all HTTP handlers
type Handler struct {
	AuthService    service.AuthService
	UserService    service.UserService
	ProductService service.ProductService
	CartService    service.CartService
	OrderService   service.OrderService
}

// NewHandler is the constructor
func NewHandler(auth service.AuthService, user service.UserService, prod service.ProductService, cart service.CartService, order service.OrderService) *Handler {
	return &Handler{
		AuthService:    auth,
		UserService:    user,
		ProductService: prod,
		CartService:    cart,
		OrderService:   order,
	}
}