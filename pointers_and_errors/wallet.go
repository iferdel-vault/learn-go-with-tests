package main

import "fmt"

type Wallet struct {
	balance int
}

func (w Wallet) Deposit(amount int) {
	w.balance += amount
	fmt.Println("address of wallet balance in Deposit method is", &w.balance)
}

func (w Wallet) Balance() int {
	fmt.Printf("address of wallet balance in Balance method is %p \n", &w.balance)
	return w.balance
}
