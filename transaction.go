package main

import "time"

type Transaction struct {
	SenderID          string
	ReceiverID        string
	TransactionAmount float64
	CreatedAt         time.Time
}
