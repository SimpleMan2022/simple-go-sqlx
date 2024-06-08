package usecase

import (
	"errors"
	"go-gin-sqlx/domain"
	"go-gin-sqlx/repository"
)

type UsersUsecase interface {
	Login(email string, password string) (*domain.UserResponse, error)
}

type usersUsecase struct {
	usersRepository repository.RepositoryUsers
}

func NewUsersUsecase(users repository.RepositoryUsers) UsersUsecase {
	return &usersUsecase{usersRepository: users}
}

func (uc usersUsecase) Login(email string, password string) (*domain.UserResponse, error) {
	if email == "" && password == "" {
		return nil, errors.New("All field are required")
	}

	usersLogin, err := uc.usersRepository.Login(email, password)
	if err != nil {
		return nil, errors.New("Failed to login, check your credentials")
	}
	return usersLogin, nil
}
