package users

import "errors"

var (
	NotEnoughMoney = errors.New("not enough money")
)

type IUser interface {
	GetBalance() int
	Withdrawal(amount int) error
	Deposit(amount int) error
	Send(userID int, amount int) error
}

var _ = (IUser)(&User{})

type User struct {
	ID     int
	UserID string
	Amount int
}

func NewUser(userID string) *User {
	return &User{UserID: userID}
}

func (b *User) GetBalance() int {
	return b.Amount
}
func (b *User) Withdrawal(amount int) error {
	if amount > b.GetBalance() {
		return NotEnoughMoney
	}
	b.Amount -= amount
	return nil
}
func (b *User) Deposit(amount int) error {
	b.Amount += amount
	return nil
}
func (b *User) Send(userID int, amount int) error {
	if amount > b.GetBalance() {
		return NotEnoughMoney
	}
	b.Amount -= amount
	return nil
}
