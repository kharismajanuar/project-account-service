package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"project/models"
)

func MenuGetUser(db *sql.DB) int {

	opsi := 1
	for opsi != 9 {
		fmt.Println("\nMenu Lihat Profil\n1. Cari akun berdasarkan nama\n2. Cari akun berdasarkan nomor telpon\n9. Kembali Ke Menu Utama")
		fmt.Print("\nPilih menu: ")
		fmt.Scanln(&opsi)
		switch opsi {
		case 1:
			GetUserByName(db)
		case 2:
			GetUserByPhone(db)
		}
	}
	return -1
}

func GetUserByName(db *sql.DB) {
	//input pencarian berdasarkan nama
	search := "%"
	var name string
	fmt.Print("\nMasukan pencarian: ")
	fmt.Scanln(&name)
	search += name
	search += "%"

	rows, errSelect := db.Query("SELECT id, name, phone FROM users WHERE deleted_at IS NULL AND name LIKE ?;", search)
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

	// fmt.Println(allUser)
	fmt.Print("\n")
	fmt.Println("Hasil pencarian:")
	for _, users := range allUsers {
		fmt.Print("\n")
		fmt.Printf("ID         : %d\n", users.ID)
		fmt.Printf("Nama       : %s\n", users.Name)
		fmt.Printf("No Telepon : %s\n", users.Phone)
	}
}

func GetUserByPhone(db *sql.DB) {
	//input pencarian berdasarkan no telpon
	search := "%"
	var phone string
	fmt.Print("\nMasukan pencarian: ")
	fmt.Scanln(&phone)
	search += phone
	search += "%"

	rows, errSelect := db.Query("SELECT id, name, phone FROM users WHERE deleted_at IS NULL AND phone LIKE ?;", search)
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

	// fmt.Println(allUser)
	fmt.Print("\n")
	fmt.Println("Hasil pencarian:")
	for _, users := range allUsers {
		fmt.Print("\n")
		fmt.Printf("ID         : %d\n", users.ID)
		fmt.Printf("Nama       : %s\n", users.Name)
		fmt.Printf("No Telepon : %s\n", users.Phone)
	}
}
