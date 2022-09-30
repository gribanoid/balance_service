package users

import (
	"context"
	"github.com/gribanoid/balance_service/pkg/pgtasks"
	"time"
)

type IUsersRepository interface {
	GetUserByID(ctx context.Context, userID string, task *pgtasks.Task) (IUser, error)
	CreateUser(ctx context.Context, userID string, task *pgtasks.Task) error
	Update(ctx context.Context, user IUser, task *pgtasks.Task) error
	GetTask(ctx context.Context, timeout time.Duration) (*pgtasks.Task, error)
}
