package rest

import "github.com/labstack/echo/v4"

func SetupRoutes(e *echo.Echo) {
	// Root and health check routes
	e.GET("/", RootHandler)
	e.GET("/health", HealthCheckHandler)
	
	// Routes group for authenticated routes
	e.Group("/auth")
	{
		SetupAuthRoutes(e)
	}
}
