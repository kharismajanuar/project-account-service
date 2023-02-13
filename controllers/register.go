package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"project/models"
)

func RegisterUser(db *sql.DB, newUser models.User) {
	query := "INSERT INTO users (name, phone, password, sex, date_of_birth) VALUES (?, ?, ?, ?, ?);"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("error prepare insert", errPrepare.Error())
	}

	result, errInsert := statement.Exec(newUser.Name, newUser.Phone, newUser.Password, newUser.Sex, newUser.DateOfBirth)
	if errInsert != nil {
		log.Fatal("error exec insert", errInsert.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Print("\n")
			fmt.Printf("Akun dengan nama '%s' telah berhasil ditambahkan!\n", newUser.Name)
		} else {
			fmt.Println("Gagal menambahkan akun baru!")
		}
	}
}
