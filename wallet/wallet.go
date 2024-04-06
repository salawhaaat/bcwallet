package wallet

import (
	"errors"
	"fmt"
	"sync"
)

// Bitcoin represents the amount of Bitcoin in the wallet.
type Bitcoin float64

// Wallet represents a Bitcoin wallet.
type Wallet struct {
	balance Bitcoin
	mutex   sync.Mutex
}

// function New creates a new wallet with the specified amount of Bitcoin.
func New(amount Bitcoin) *Wallet {
	return &Wallet{
		balance: amount,
		mutex:   sync.Mutex{},
	}
}

// function String returns the string representation of the Bitcoin amount.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}

// function Deposit adds the specified amount of Bitcoin to the wallet.
func (w *Wallet) Deposit(amount Bitcoin) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if amount < 0 {
		return errors.New("cannot deposit negative amount")
	}
	w.balance += amount
	return nil
}

// function Withdraw deducts the specified amount of Bitcoin from the wallet.
func (w *Wallet) Withdraw(amount Bitcoin) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if amount > w.balance {
		return errors.New("insufficient funds")
	}
	w.balance -= amount
	return nil
}

// function Balance returns the current balance of the wallet.
func (w *Wallet) Balance() Bitcoin {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	return w.balance
}
