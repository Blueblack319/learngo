package accounts

import (
	"errors"
	"fmt"
)

// errors
var errWithdraw = errors.New("Insufficient balance")

// Account struct
type Account struct {
	owner    string
	balance  int
	currUnit string
}

// NewAccount to create a new account
func NewAccount(owner string) *Account {
	return &Account{
		owner:    owner,
		balance:  0,
		currUnit: "KRW",
	}
}

// Deposit to an account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance to show remainer
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from an account
func (a *Account) Withdraw(amount int) error {
	if amount > a.balance {
		return errWithdraw
	}
	a.balance -= amount
	return nil
}

func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.owner, "'s account. \nHas: ", a.balance)
}
