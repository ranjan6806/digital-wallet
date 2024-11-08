package main

import "time"

type Account struct {
	UserID       string
	Wallet       *Wallet
	CreatedAt    time.Time
	Transactions []*Transaction
}

func NewAccount(userID string, amount float64) *Account {
	newWallet := NewWallet(amount)
	return &Account{
		UserID:       userID,
		Wallet:       newWallet,
		CreatedAt:    time.Now(),
		Transactions: make([]*Transaction, 0),
	}
}

func (a *Account) Deposit(senderUserID, receiverUserID string, amount float64) {
	a.Wallet.AddAmount(amount)
	a.Transactions = append(a.Transactions, &Transaction{
		SenderID:          senderUserID,
		ReceiverID:        receiverUserID,
		TransactionAmount: amount,
		CreatedAt:         time.Now(),
	})
}

func (a *Account) Deduct(senderUserID, receiverUserID string, amount float64) error {
	err := a.Wallet.DeductAmount(amount)
	if err != nil {
		return err
	}

	// issue in this
	a.Transactions = append(a.Transactions, &Transaction{
		SenderID:          senderUserID,
		ReceiverID:        receiverUserID,
		TransactionAmount: amount,
	})
	return nil
}
