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
		fmt.Println("tidak dapat menampilkan info akun")
		return -1
	}

	fmt.Println("Nomor telepon\t:", user.Phone)
	fmt.Println("Nama\t:", user.Name)
	fmt.Println("Tanggal Lahir\t:", user.DateOfBirth)
	fmt.Println("Jenis Kelamin\t:", user.Sex)

	//tampilkan saldo
	var saldo float64
	err = db.QueryRow("SELECT balance FROM balances WHERE user_id = ?", user.ID).Scan(&saldo)
	if err != nil {
		fmt.Println("tidak dapat menampilkan saldo")
		return -1
	}

	fmt.Printf("Saldo\t:%.2f\n", saldo)

	var menu int
	fmt.Println("\npilih :\n1.Kembali ke menu utama\n2.Exit")
	fmt.Scanln(&menu)
	if menu == 1 {
		return -1
	} else {
		return 9
	}
}
