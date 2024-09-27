package rest

import (
	"database/sql"
	"go-app/internal/repository"
	"go-app/internal/service"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, db *sql.DB) {
	// Root and health check routes
	e.GET("/", RootHandler)
	e.GET("/health", HealthCheckHandler)
	
	// Routes group for authenticated routes
	authRepo := repository.NewAuthRepository(db)
	authSvc := service.NewAuthService(authRepo)
	e.Group("/auth")
	{
		SetupAuthRoutes(e, authSvc)
	}
}
