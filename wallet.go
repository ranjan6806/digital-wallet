package main

import "fmt"

type Wallet struct {
	Amount float64
}

func NewWallet(amount float64) *Wallet {
	return &Wallet{amount}
}

func (w *Wallet) AddAmount(amount float64) {
	w.Amount += amount
}

func (w *Wallet) DeductAmount(amount float64) error {
	if amount > w.Amount {
		return fmt.Errorf("cannot deduct amount %.2f", amount)
	}

	w.Amount -= amount
	return nil
}
