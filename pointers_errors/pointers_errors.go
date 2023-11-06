package pointers_errors

import "fmt"

type Wallet struct {
	balance int
}

/*
// w is defined as Wallet type
func (w Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance) // different memory address than in the test
	w.balance += amount
}

func (w Wallet) Balance() int {
	return w.balance
}
*/

// *Wallet -pointer
// w is defined as pointer to Wallet type; receiver type is *Wallet
func (w *Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance) // the same memory address as in the test
	// struct pointer - automatically dereferenced
	w.balance += amount
	// another valid code using explicit dereference, but not commonly used
	//(*w).balance
}

// by convention you should keep your method receiver types the same for consistency (the receiver could be defained as w Wallet instead w *Wallet)
func (w *Wallet) Balance() int {
	return w.balance
}

// Refactor https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors#refactor
// creating a new type Bitcoin. To make Bitcoin use syntax Bitcoin(10).
// type MyName OriginalType
type Bitcoin int

type BitcoinWallet struct {
	balance Bitcoin
}

// methods on BitcoinWallet
func (bw *BitcoinWallet) BitcoinDeposit(amount Bitcoin) {
	bw.balance += amount
}

func (bw *BitcoinWallet) BitcoinBalance() Bitcoin {
	return bw.balance
}

func (bw *BitcoinWallet) BitcoinWithdraw(amount Bitcoin) {
	bw.balance -= amount
}

/*
	MyStringer Interface
	This interface is defined in the fmt package
	and lets you define how your type is printed when used
	with the %s format string in prints.
*/
/*
type MyStringer interface {
	String() string
}
*/
// the syntax for creating a method on a type declaration
// is the same as it is on a struct.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
