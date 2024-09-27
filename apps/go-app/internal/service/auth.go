package service

import (
	"context"
	"go-app/internal/model"
)

type AuthRepository interface {
	CheckLoginInfo(ctx context.Context, email string)(model.User, error)
}

type AuthService struct {
	authRepo AuthRepository 
}

// NewService will create a new article service object
func NewAuthService(a AuthRepository) *AuthService {
	return &AuthService{
		authRepo: a,
	}
}

func (a *AuthService) LoginHandler(ctx context.Context, email string) (res model.User, err error) {
	res, err = a.authRepo.CheckLoginInfo(ctx, email)
	if err != nil {
		return
	}
	return
}