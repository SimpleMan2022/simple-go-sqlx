package repository

import (
	"github.com/jmoiron/sqlx"
	"go-gin-sqlx/domain"
)

type RepositoryUsers interface {
	Login(email string, password string) (*domain.UserResponse, error)
}

type usersRepository struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) RepositoryUsers {
	return &usersRepository{db}
}

func (r *usersRepository) Login(email string, password string) (*domain.UserResponse, error) {
	var user domain.UserResponse
	err := r.db.Get(&user, "SELECT Name, Email FROM users WHERE Email = ? AND Password = ?", email, password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
