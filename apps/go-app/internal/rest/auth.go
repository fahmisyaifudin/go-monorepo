package rest

import (
	"context"
	"fmt"
	"go-app/internal/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthService interface {
	LoginHandler(ctx context.Context, email string) (model.User, error)
	// Add other methods your AuthService provides
}

type AuthHandler struct {
	Service AuthService
}

func SetupAuthRoutes(e *echo.Echo, svc AuthService){
	handler := &AuthHandler{
		Service: svc,
	}

	e.POST("/login", handler.LoginHandler)
	e.POST("/forgot-password", handler.ForgotPasswordHandler)
	e.PATCH("/reset-password", handler.ResetPasswordHandler)
}

func (a *AuthHandler) LoginHandler(c echo.Context) error {
	var req model.LoginRequest
	fmt.Println(req.Email)
	if err := c.Bind(&req); err != nil {
		return err
	}
	
	ctx := c.Request().Context()
	user, err := a.Service.LoginHandler(ctx, req.Email)
	fmt.Println(user)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Not found"})
	}

	response := model.LoginResponse{
		BaseAPIResponse: model.BaseAPIResponse{Status: http.StatusOK, Success: true},
		Data: model.LoginData{
			AccessToken:  "sample_access_token",
			RefreshToken: "sample_refresh_token",
			Role:         "admin",
		},
	}
	return c.JSON(http.StatusOK, response)
}

func (a *AuthHandler) ForgotPasswordHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Endpoint for forgot password"})
}

func (a *AuthHandler) ResetPasswordHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Endpoint for reset password"})
}
