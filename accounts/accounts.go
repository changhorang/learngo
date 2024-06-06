package accounts

import "errors"

// Account struck : private
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't Withdraw you don't have enough balance")

// NewAccount create Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account : Method add & * (dont copy)
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// Balance of your receiver
func (a Account) Balance() int {
	return a.balance
}
