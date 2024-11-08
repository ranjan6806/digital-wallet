package main

import "fmt"

func main() {
	repo := NewInMemoryAccountsRepo()
	service := NewWalletService(repo)

	err := service.CreateWallet("u1", 100)
	if err != nil {
		fmt.Println(err)
	}

	err = service.CreateWallet("u2", 50)
	if err != nil {
		fmt.Println(err)
	}

	noOffer := NewNoOffer()

	err = service.TransferMoney("u1", "u2", 20, noOffer)
	if err != nil {
		fmt.Println(err)
	}

	service.Statement("u1")
	service.Statement("u2")
}
