package main

import "fmt"

type Account struct {
	balance float64
}

func NewAccount(b float64) *Account {
	a := Account{balance: b}
	return &a
}

func (a *Account) Deposit(amount float64) {
	a.balance = a.balance + amount
}

func (a *Account) Withdraw(amount float64) {
	a.balance = a.balance - amount
}

func (a *Account) Display() {
	fmt.Printf("    Balance: $%v\n", a.balance)
}

func main() {
	a := NewAccount(100.00)
	fmt.Println("Before Transactions:")
	a.Display()

	a.Deposit(74.35)
	a.Withdraw(20.00)

	fmt.Println("After Transactions:")
	a.Display()
}
