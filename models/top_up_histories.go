package models

import "time"

type TopUpHistories struct {
	ID     int
	UserID int
	Date   time.Time
	Amount float64
	Info   string
}
