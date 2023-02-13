package controllers

import (
	"database/sql"
	"project/models"
)

func Login(db *sql.DB) (models.User, int, bool) {
	return models.User{}, 0, true
}
