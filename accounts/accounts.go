package accounts

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
