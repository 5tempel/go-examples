package pointers_errors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()

	fmt.Printf("address of balance in test is %v \n", &wallet.balance)

	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestBitcoinWallet(t *testing.T) {

	wallet := BitcoinWallet{}

	wallet.BitcoinDeposit(Bitcoin(10))

	got := wallet.BitcoinBalance()

	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

// next time start with "Let's implement Stringer on Bitcoin"
// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors#refactor
