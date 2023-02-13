package models

import "time"

type User struct {
	ID          int
	Phone       string
	Name        string
	Password    string
	DateOfBirth time.Time
	Sex         string
}
