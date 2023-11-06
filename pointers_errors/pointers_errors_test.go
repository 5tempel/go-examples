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

/*
func TestBitcoinWallet(t *testing.T) {

	wallet := BitcoinWallet{}

	wallet.BitcoinDeposit(Bitcoin(10))

	got := wallet.BitcoinBalance()

	want := Bitcoin(100)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
*/
// next time start with "Let's implement Stringer on Bitcoin"
// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors#refactor

/*
func TestBitcoinWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := BitcoinWallet{}

		wallet.BitcoinDeposit(Bitcoin(10))

		got := wallet.BitcoinBalance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := BitcoinWallet{balance: Bitcoin(20)}

		wallet.BitcoinWithdraw(Bitcoin(10))

		got := wallet.BitcoinBalance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}

*/

// refactor
func TestBitcoinWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet BitcoinWallet, want Bitcoin) {
		t.Helper()
		got := wallet.BitcoinBalance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := BitcoinWallet{}
		wallet.BitcoinDeposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := BitcoinWallet{balance: Bitcoin(20)}
		wallet.BitcoinWithdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
}

/*
 next time start here:
 https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors#write-the-test-first-2

 start with:
 "What should happen if you try to Withdraw more
 than is left in the account?"
*/
