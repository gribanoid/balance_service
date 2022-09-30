package users

import "errors"

var (
	NotEnoughMoney = errors.New("not enough money")
)

type IUser interface {
	GetBalance() int
	Withdrawal(amount int) error
	Deposit(amount int)
	Convert() *User
}

var _ = (IUser)(&User{})

type User struct {
	ID      int    `json:"id" db:"id"`
	UserID  string `json:"user_id" db:"user_id"`
	Balance int    `json:"balance" db:"balance"`
}

func NewUser(userID string) *User {
	return &User{UserID: userID}
}

func (u *User) GetBalance() int {
	return u.Balance
}
func (u *User) Withdrawal(amount int) error {
	if amount > u.GetBalance() {
		return NotEnoughMoney
	}
	u.Balance -= amount
	return nil
}
func (u *User) Deposit(amount int) {
	u.Balance += amount
	return
}

func (u *User) Convert() *User {
	return u
}
