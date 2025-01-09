package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// in these methods if we pass the struct as a argument then each method create it's own copy and update them
// so to avoid this confusion we need to pass the Wallet memory address
// by this all the methods will update the same balance in the struct

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposite(amount Bitcoin) {
	fmt.Printf("addresse of the balance in is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	w.balance -= amount
	if amount > w.balance {
		return errors.New("oh no")
	}
	return nil
}
