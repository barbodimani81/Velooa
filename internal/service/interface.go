package service

import (
	"context"

	"github.com/barbodimani81/Velooa.git/internal/models"
)

type AuthService interface {
	Register(ctx context.Context, fullName, email, password, phone string) (*models.User, error)
	Login(ctx context.Context, email, password string) (string, error) // Returns JWT
}

type ProductService interface {
	ListProducts(ctx context.Context, categoryID uint) ([]models.Product, error)
	GetProduct(ctx context.Context, productID uint) (*models.Product, error)
	ListCategories(ctx context.Context) ([]models.Category, error)
}

type CartService interface {
	GetCart(ctx context.Context, userID uint) (*models.Cart, error)
	AddToCart(ctx context.Context, userID uint, variantID uint, quantity int) error
	UpdateCartItemQuantity(ctx context.Context, userID, cartItemID uint, newQuantity int) error
	RemoveFromCart(ctx context.Context, userID, cartItemID uint) error
}

type OrderService interface {
	PlaceOrder(ctx context.Context, userID uint, addressID uint, paymentToken string) (*models.Order, error)
	GetOrder(ctx context.Context, orderID, userID uint) (*models.Order, error)
}
