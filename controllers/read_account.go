package controllers

import (
	"database/sql"
	"fmt"
	"project/models"
)

func ReadAccount(db *sql.DB, user models.User) int {
	fmt.Println("Nomor telepon\t:", user.Phone)
	fmt.Println("Nama\t:", user.Name)
	fmt.Println("Tanggal Lahir\t:", user.DateOfBirth)
	fmt.Println("Jenis Kelamin\t:", user.Sex)
	fmt.Println("\npilih :\n1.Kembali ke menu utama\n2.Exit")
	var menu int
	fmt.Scanln(&menu)
	if menu == 1 {
		return -1
	} else {
		return 9
	}
}
