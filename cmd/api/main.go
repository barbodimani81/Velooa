package main

import (
	"github.com/barbodimani81/Velooa/internal/api/handlers"
	"github.com/labstack/echo/v4"
	// Import your service implementations once they exist
)

// @title Velooa API
// @version 1.0
// @description E-commerce API for Velooa
// @host localhost:8080
// @BasePath /api/v1
func main() {
    e := echo.New()

    // 1. Initialize your Services (Currently nil until implemented)
    // var authSvc service.AuthService
    // var userSvc service.UserService
    // ... etc

    // 2. Initialize the Handler struct with those services
    // Note: We use the package name 'handlers' to call NewHandler
    h := handlers.NewHandler(nil, nil, nil, nil, nil) 

    // 3. Call the method on the instance 'h', not the package
    h.RegisterRoutes(e)

    e.Logger.Fatal(e.Start(":8080"))
}