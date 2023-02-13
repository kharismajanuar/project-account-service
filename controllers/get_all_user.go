package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"project/models"
)

func MenuGetAllUser(db *sql.DB) {

	opsi := 1
	for opsi != 9 {
		fmt.Print("\n")
		GetAllUsers(db)
		fmt.Print("9. Kembali Ke Menu Utama\n")
		fmt.Print("\nPilih menu: ")
		fmt.Scanln(&opsi)
	}
}

func GetAllUsers(db *sql.DB) {
	rows, errSelect := db.Query("SELECT id, name, phone FROM users;")
	if errSelect != nil {
		log.Fatal("error query select", errSelect.Error())
	}

	var allUsers []models.User
	for rows.Next() {
		var datarow models.User
		errScan := rows.Scan(&datarow.ID, &datarow.Name, &datarow.Phone)
		if errScan != nil {
			log.Fatal("error scan select", errScan.Error())
		}
		allUsers = append(allUsers, datarow)
	}

	// fmt.Println(allGuru)
	for _, users := range allUsers {
		fmt.Printf("ID         : %d\n", users.ID)
		fmt.Printf("Nama       : %s\n", users.Name)
		fmt.Printf("No Telepon : %s\n", users.Phone)
		fmt.Print("\n")
	}
}
