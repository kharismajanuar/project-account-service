package controllers

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func updatePassword(phone string, db *sql.DB) int {
	//input password
	var data string
	fmt.Println("input password :")
	fmt.Scanln(&data)

	//validasi password
	if len(data) < 5 {
		fmt.Println("password minimal 5 huruf")
		return -1
	}

	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data), 10)
	if err != nil {
		fmt.Println("update password gagal")
		return -1
	}
	_, err = db.Exec("UPDATE users SET password = ? WHERE phone = ?", string(hashedPassword), phone)
	if err != nil {
		fmt.Println("update password gagal")
		return -1
	}
	return -1
}

func updateTanggalLahir(phone string) int {
	return -1
}

func updateTelepon(phone string) int {
	return -1
}

func updateNama(phone string) int {
	return -1
}

func UpdateAccount(db *sql.DB, phone string) int {
	//input pilih data yang akan diupdate

	var opsi int = -1
	//pilih sesuai menu
	for opsi != 5 {
		fmt.Println("pilih data yang akan diupdate\n1.Telepon\n2.Nama\n3.Tanggal Lahir\n4.Password\n5.Menu Utama")
		fmt.Scanln(&opsi)
		switch opsi {
		case 1:

		case 2:

		case 3:

		case 4:
			updatePassword(phone, db)
		case 5:
			break
		}
	}
	return -1
}
