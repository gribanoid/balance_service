package user

import (
	"context"
	"github.com/gribanoid/balance_service/internal/core/users"
	"github.com/gribanoid/balance_service/internal/repositories/postgres"
	"github.com/gribanoid/balance_service/pkg/pgtasks"
	"time"
)

type UsersRepository struct {
	storage pgtasks.AbstactStorageTx
}

func NewUserRepository(pool *postgres.PGConnPool, timeout time.Duration) *UsersRepository {
	return &UsersRepository{
		storage: pgtasks.NewStorage(pool, timeout),
	}
}
func (r *UsersRepository) GetTask(ctx context.Context, timeout time.Duration) (*pgtasks.Task, error) {
	return r.storage.CreateTx(ctx, timeout)
}
func (r *UsersRepository) CreateUser(ctx context.Context, user *users.User, task *pgtasks.Task) error {
	query := `insert into balance_service.users
		(
		id,
		user_id,
		amount
		)
		values
		($1, $2, $3)
		`

	innertask := task == nil
	if innertask {
		var err error
		task, err = r.GetTask(ctx, time.Second*5)
		if err != nil {
			return err
		}
		defer task.Rollback(ctx)
	}

	_, err := task.Tx.Exec(ctx,
		query,
		user.ID,
		user.UserID,
		user.Amount,
	)

	if err != nil {
		return err
	}

	if innertask {
		if err := task.Commit(ctx); err != nil {
			return err
		}
	}

	return err
}

func (r *UsersRepository) Update(ctx context.Context, user *users.User, task *pgtasks.Task) error {
	//TODO implement me
	panic("implement me")
}

func (r *UsersRepository) GetUserByID(ctx context.Context, userID string, task *pgtasks.Task) (users.IUser, error) {
	//TODO implement me
	panic("implement me")
}
