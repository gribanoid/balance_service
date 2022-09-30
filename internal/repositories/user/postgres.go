package user

import (
	"context"
	"errors"
	"fmt"
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
func (r *UsersRepository) CreateUser(ctx context.Context, userID string, task *pgtasks.Task) error {
	if _, err := r.GetUserByID(ctx, userID, task); nil == err {
		return errors.New("user already exists") //TODO
	}
	fmt.Println("прошел!!!!!!!!!!!!!!!!")
	query := `insert into balance_service.users
		(
		user_id,
		balance
		)
		values
		($1, $2)
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
		userID,
		0,
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

func (r *UsersRepository) Update(ctx context.Context, user users.IUser, task *pgtasks.Task) error {
	u := user.Convert()

	query := `update balance_service.users
		set user_id = $1,
			balance = $2
		where id = $3
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
		u.UserID,
		u.Balance,
		u.ID,
	)

	if innertask {
		if err := task.Commit(ctx); err != nil {
			return err
		}
	}

	return err
}

func (r *UsersRepository) GetUserByID(ctx context.Context, userID string, task *pgtasks.Task) (users.IUser, error) {
	var u users.User

	innertask := task == nil
	if innertask {
		var err error
		task, err = r.GetTask(ctx, time.Second*5)
		if err != nil {
			return nil, err
		}
		defer task.Rollback(ctx)
	}

	if err := task.Tx.QueryRow(ctx, `select  
			id, 
			user_id, 
			amount
		from balance_service.users where user_id = $1`, userID).Scan(&u.ID,
		&u.UserID,
		&u.ID,
	); err != nil {
		return nil, err
	}

	if innertask {
		if err := task.Commit(ctx); err != nil {
			return nil, err
		}
	}

	return &u, nil
}
