package handlers

import (
	_ "github.com/barbodimani81/Velooa/docs" // This import is required for Swagger to load generated docs
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// RegisterRoutes sets up the routing for the Echo engine
func (h *Handler) RegisterRoutes(e *echo.Echo) {
	// --- Global Middleware ---
	e.Use(middleware.Logger())    // Logs every request to terminal
	e.Use(middleware.Recover())   // Recovers from panics to prevent server crashes
	e.Use(middleware.CORS())      // Allows frontend to communicate with backend

	// --- Swagger Documentation ---
	// Accessible at: http://localhost:8080/swagger/index.html
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// --- API Version 1 Group ---
	api := e.Group("/api/v1")

	// --- Public Routes ---
	
	// Auth Routes
	auth := api.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
	}

	// Product & Category Routes
	api.GET("/categories", h.ListCategories)
	api.GET("/products", h.ListProducts)
	api.GET("/looks/:id", h.GetLook) // For "Shop the Look" feature

	// Cart Routes (Can be accessed by guests/logged-in users)
	cart := api.Group("/cart")
	{
		cart.GET("", h.GetCart)
		cart.POST("/items", h.AddToCart)
	}

	// --- Protected Routes (Require Authentication) ---
	// Note: You will eventually wrap this in your JWT Middleware
	// e.g., protected := api.Group("", YourJWTMiddleware())
	
	orders := api.Group("/orders")
	{
		orders.POST("", h.PlaceOrder)
	}
}