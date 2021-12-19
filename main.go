package main

import (
	"fmt"
	"log"

	"github.com/crazybirdz/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("Yane Jae-Hoon")
	account.Deposit(1000)
	fmt.Println(account.Balance())
	err := account.Withdraw(1010)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())
}
