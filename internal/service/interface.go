package service

import (
	"context"

	"github.com/barbodimani81/Velooa/internal/models"
)

type AuthService interface {
	Register(ctx context.Context, fullName, email, password, phone string) (*models.User, error)
	Login(ctx context.Context, email, password string) (string, error) // Returns JWT
	ForgotPassword(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, token, newPassword string) error
}

type UserService interface {
	GetProfile(ctx context.Context, userID uint) (*models.User, error)
	UpdateProfile(ctx context.Context, userID uint, updates models.User) (*models.User, error)
	
	ListAddresses(ctx context.Context, userID uint) ([]models.Address, error)
	AddAddress(ctx context.Context, userID uint, addr models.Address) (*models.Address, error)
	UpdateAddress(ctx context.Context, userID, addressID uint, addr models.Address) error
	DeleteAddress(ctx context.Context, userID, addressID uint) error
}

type ProductService interface {
	ListProducts(ctx context.Context, categoryID uint, page, limit int) ([]models.Product, error)
	GetProduct(ctx context.Context, productID uint) (*models.Product, error)
	ListCategories(ctx context.Context) ([]models.Category, error)
	ListLooks(ctx context.Context) ([]models.Look, error)
	GetLook(ctx context.Context, lookID uint) (*models.Look, []models.Product, error)
	SearchProducts(ctx context.Context, query string) ([]models.Product, error)
}

type CartService interface {
	GetCart(ctx context.Context, userID *uint, cartID *uint) (*models.Cart, error)
	AddToCart(ctx context.Context, userID *uint, cartID *uint, variantID uint, quantity int) (*models.Cart, error)
	UpdateCartItemQuantity(ctx context.Context, userID *uint, cartItemID uint, newQuantity int) error
	RemoveFromCart(ctx context.Context, cartItemID uint) error
	MergeCarts(ctx context.Context, guestCartID uint, userID uint) error
}

type OrderService interface {
	PlaceOrder(ctx context.Context, userID uint, addressID uint, paymentProviderID string) (*models.Order, error)
	ListOrders(ctx context.Context, userID uint) ([]models.Order, error)
	GetOrder(ctx context.Context, orderID, userID uint) (*models.Order, error)
	CancelOrder(ctx context.Context, orderID, userID uint) error
}