package main

import (
	"fmt"

	"github.com/crazybirdz/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("Yane Jae-Hoon")
	account.Deposit(1000)
	fmt.Println(account.Balance())
}
