package pointers

import (
	"errors"
	"fmt"
)

// ErrorInsufficientFunds new type Error
var ErrorInsufficientFunds = errors.New("opsssss, insufficient funds")

// Bitcoin type float64
type Bitcoin float64

// Description with balance
type Description interface {
	String() string
}

// Wallet with balance
type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.8f BTC", b)
}

// Balance return balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Deposit in your wallet
func (w *Wallet) Deposit(value Bitcoin) {
	w.balance += value
}

// Withdrawal in your wallet
func (w *Wallet) Withdrawal(value Bitcoin) error {

	if value > w.balance {
		return ErrorInsufficientFunds
	}

	w.balance -= value
	return nil
}
