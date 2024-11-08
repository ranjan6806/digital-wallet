package main

import (
	"errors"
)

type AccountsRepo interface {
	CreateWallet(userID string, amount float64) error
	TransferMoney(senderUserID, receiverUserID string, amount float64, offer Offer) error
	Statement(userID string) ([]*Transaction, error)
	Overview() map[string]*Account
}

type InMemoryAccountsRepo struct {
	UserAccounts map[string]*Account
}

func NewInMemoryAccountsRepo() *InMemoryAccountsRepo {
	return &InMemoryAccountsRepo{
		UserAccounts: make(map[string]*Account),
	}
}

func (r *InMemoryAccountsRepo) CreateWallet(userID string, amount float64) error {
	if _, ok := r.UserAccounts[userID]; ok {
		return errors.New("account already exists")
	}

	newAccount := NewAccount(userID, amount)
	r.UserAccounts[userID] = newAccount
	return nil
}

func (r *InMemoryAccountsRepo) TransferMoney(senderID string, receiverUserID string, amount float64, offer Offer) error {
	if _, ok := r.UserAccounts[senderID]; !ok {
		return errors.New("sender account does not exist")
	}

	if _, ok := r.UserAccounts[receiverUserID]; !ok {
		return errors.New("receiver account does not exist")
	}

	senderAccount := r.UserAccounts[senderID]
	receiverAccount := r.UserAccounts[receiverUserID]

	receiverAccount.Deposit(senderID, receiverUserID, amount)
	err := senderAccount.Deduct(senderID, receiverUserID, amount)
	if err != nil {
		receiverAccount.Deduct(senderID, receiverUserID, amount)
		return err
	}

	offerType := offer.GetOfferType()
	switch offerType {
	case OfferType1:
		if receiverAccount.Wallet.Amount == senderAccount.Wallet.Amount {
			receiverAccount.Wallet.Amount += 10
			senderAccount.Wallet.Amount += 10
		}
	}

	return nil
}

func (r *InMemoryAccountsRepo) Statement(userID string) ([]*Transaction, error) {
	if _, ok := r.UserAccounts[userID]; !ok {
		return nil, errors.New("account does not exist")
	}

	userAccount := r.UserAccounts[userID]
	return userAccount.Transactions, nil
}

func (r *InMemoryAccountsRepo) Overview() map[string]*Account {
	return r.UserAccounts
}
