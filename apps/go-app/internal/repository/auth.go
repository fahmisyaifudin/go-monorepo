package repository

import (
	"context"
	"database/sql"
	"go-app/internal/model"
)

type AuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		DB: db,
	}
}

func (m *AuthRepository) CheckLoginInfo(ctx context.Context, email string) (res model.User, err error)  {
	query := `SELECT * FROM user WHERE email=?`
	stmt, err := m.DB.PrepareContext(ctx, query)
	row := stmt.QueryRowContext(ctx, email)
	res = model.User{}

	err = row.Scan(
		&res.ID,
		&res.Email,
		&res.Password,
	)
	return
}