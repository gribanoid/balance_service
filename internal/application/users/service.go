package users

import (
	"github.com/gribanoid/balance_service/internal/core/users"
)

type UserService struct {
	usersRepo users.IUsersRepository
}

func NewUserService(usersRepo users.IUsersRepository) (*UserService, error) {
	return &UserService{usersRepo: usersRepo}, nil
}
