package models

import "time"

type TransferHistories struct {
	ID             int
	UserIDSender   int
	UserIDReceiver int
	Date           time.Time
	Amount         float64
	Info           string
	ReceiverName   string
	ReceiverPhone  string
}
