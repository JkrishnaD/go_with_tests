package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	t.Run("deposite", func(t *testing.T) {
		wallet.Deposite(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		initialBalance := Bitcoin(20)

		wallet := Wallet{initialBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, initialBalance)
		if err != nil {
			t.Errorf("wanted an error but didn't get one")
		}
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}
}
