package accounts

type Account struct {
	owner    string
	balance  int
	currUnit string
}

func NewAccount(owner string) *Account {
	return &Account{
		owner:    owner,
		balance:  0,
		currUnit: "KRW",
	}
}
