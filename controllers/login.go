package controllers

import (
	"database/sql"
	"project/models"
)

func Login(db *sql.DB, phone string, password string) (models.User, int) {
	return models.User{}, 0
}
