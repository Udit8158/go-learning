package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	Balance Bitcoin
}

var InsufficientFundsError = errors.New("You must have more than what you want to withdraw!")

func (w *Wallet) Deposit(deposit_amount Bitcoin) {
	// fmt.Printf("In method %p and w = %p\n", &w.Balance, &w)
	w.Balance += deposit_amount
}

func (w *Wallet) Withdraw(bal_to_withdraw Bitcoin) error {
	if w.Balance < bal_to_withdraw {
		return InsufficientFundsError
	}
	w.Balance -= bal_to_withdraw
	return nil // if a return type is an interface then we can have nil return too
}

// we can have method on top of a custom type too
// so ig that'w where Sprintf used - for c
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
