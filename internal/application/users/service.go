package users

import (
	"context"
	"github.com/gribanoid/balance_service/internal/core/users"
	"time"
)

type UserService struct {
	usersRepo users.IUsersRepository
}

func NewUserService(usersRepo users.IUsersRepository) (*UserService, error) {
	return &UserService{usersRepo: usersRepo}, nil
}

func (s *UserService) CreateUser(ctx context.Context, userID string) error {
	task, err := s.usersRepo.GetTask(ctx, time.Second*5)
	if err != nil {
		return err
	}
	defer task.Rollback(ctx)

	if err = s.usersRepo.CreateUser(ctx, userID, task); err != nil {
		return err
	}
	return task.Commit(ctx)

}
func (s *UserService) GetBalance(ctx context.Context, userID string) (int, error) {
	task, err := s.usersRepo.GetTask(ctx, time.Second*5)
	if err != nil {
		return 0, err
	}
	defer task.Rollback(ctx)

	user, err := s.usersRepo.GetUserByID(ctx, userID, task)
	if err != nil {
		return 0, err
	}

	balance := user.GetBalance()
	return balance, nil
}

func (s *UserService) Withdrawal(ctx context.Context, userID string, amount int) error {
	task, err := s.usersRepo.GetTask(ctx, time.Second*5)
	if err != nil {
		return err
	}
	defer task.Rollback(ctx)

	user, err := s.usersRepo.GetUserByID(ctx, userID, task)
	if err != nil {
		return err
	}
	err = user.Withdrawal(amount)
	if err != nil {
		return err
	}
	if err = s.usersRepo.Update(ctx, user, task); err != nil {
		return err
	}
	return task.Commit(ctx)
}
func (s *UserService) Deposit(ctx context.Context, userID string, amount int) error {
	task, err := s.usersRepo.GetTask(ctx, time.Second*5)
	if err != nil {
		return err
	}
	defer task.Rollback(ctx)

	user, err := s.usersRepo.GetUserByID(ctx, userID, task)
	if err != nil {
		return err
	}
	user.Deposit(amount)

	if err = s.usersRepo.Update(ctx, user, task); err != nil {
		return err
	}

	return task.Commit(ctx)
}

//func (s *UserService) Send(ctx context.Context, from string, to string, amount int) error {
//
//	err = s.Withdrawal(ctx, from, amount)
//	err = s.Deposit(ctx, to, amount)
//
//
//	return task.Commit(ctx)
//}
