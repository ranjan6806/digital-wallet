package main

import "fmt"

type WalletService interface {
	CreateWallet(userID string, amount float64) error
	TransferMoney(senderUserID, receiverUserID string, amount float64, offer Offer) error
	Statement(userID string) error
	Overview()
	//Offer2()
}

type WalletServiceImpl struct {
	accountsRepo AccountsRepo
}

func NewWalletService(accountsRepo AccountsRepo) WalletService {
	return &WalletServiceImpl{
		accountsRepo: accountsRepo,
	}
}

func (s *WalletServiceImpl) CreateWallet(userID string, amount float64) error {
	return s.accountsRepo.CreateWallet(userID, amount)
}

func (s *WalletServiceImpl) TransferMoney(senderUserID, receiverUserID string, amount float64, offer Offer) error {
	return s.accountsRepo.TransferMoney(senderUserID, receiverUserID, amount, offer)
}

func (s *WalletServiceImpl) Statement(userID string) error {
	transactions, err := s.accountsRepo.Statement(userID)
	if err != nil {
		return err
	}

	for _, transaction := range transactions {
		fmt.Printf("transactions - %+v\n", transaction)
	}

	return nil
}

func (s *WalletServiceImpl) Overview() {
	overviewObj := s.accountsRepo.Overview()
	for userID, account := range overviewObj {
		fmt.Printf("userID - %s, accounts - %+v\n", userID, account)
	}
}
