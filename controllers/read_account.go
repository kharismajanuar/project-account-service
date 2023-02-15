package controllers

import (
	"database/sql"
	"fmt"
	"project/models"
)

func ReadAccount(db *sql.DB, ID int) int {
	//tampilkan data pribadi
	var user models.User
	err := db.QueryRow("SELECT name, phone, date_of_birth, sex FROM users WHERE id = ?", ID).
		Scan(&user.Name, &user.Phone, &user.DateOfBirth, &user.Sex)
	if err != nil {
		fmt.Println("Tidak dapat menampilkan info akun")
		return -1
	}
	fmt.Print("\n")
	fmt.Println("Nomor telepon\t:", user.Phone)
	fmt.Println("Nama\t\t:", user.Name)
	fmt.Println("Tanggal Lahir\t:", user.DateOfBirth.Format("January 2, 2006"))
	fmt.Println("Jenis Kelamin\t:", user.Sex)

	//tampilkan saldo
	var saldo float64
	err = db.QueryRow("SELECT balance FROM balances WHERE user_id = ?", ID).Scan(&saldo)
	if err != nil {
		fmt.Println("Tidak dapat menampilkan saldo")
		return -1
	}

	fmt.Printf("Saldo\t:%.2f\n", saldo)

	var menu int
	fmt.Print("\n\n1.Menu utama\n2.Exit\n\nPilih menu : ")
	fmt.Scanln(&menu)
	if menu == 1 {
		return -1
	} else {
		return 9
	}
}
