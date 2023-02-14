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
			opsiNama := 1
			for opsiNama != 9 {
				fmt.Print("\n")
				GetUserByName(db)
				fmt.Print("9. Kembali ke Menu\n")
				fmt.Print("\nPilih Menu: ")
				fmt.Scanln(&opsiNama)
			}

		case 2:

		}
	}
	return -1
}

func GetUserByName(db *sql.DB) {
	//input pencarian berdasarkan nama
	var name string
	fmt.Println("Masukan pencarian:")
	fmt.Scanln(&name)

	rows, errSelect := db.Query("SELECT id, name, phone FROM users WHERE deleted_at IS NULL AND name LIKE '% ? %';", name)
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
