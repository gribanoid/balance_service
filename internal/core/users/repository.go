package users

import (
	"context"
	"github.com/gribanoid/balance_service/pkg/pgtasks"
	"time"
)

type IUsersRepository interface {
	CreateUser(ctx context.Context, user *User, task *pgtasks.Task) error
	Update(ctx context.Context, user *User, task *pgtasks.Task) error
	GetUserByID(ctx context.Context, userID string, task *pgtasks.Task) (IUser, error)
	GetTask(ctx context.Context, timeout time.Duration) (*pgtasks.Task, error)
}
