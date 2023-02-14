package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"project/models"
)

func MenuDelete(db *sql.DB, user models.User) int {

	var opsi string = "a"

	for opsi != "n" {
		fmt.Println("\nApakah Anda yakin untuk menghapus akun?")
		fmt.Print("Pilih (y/n): ")
		fmt.Scan(&opsi)
		if opsi == "y" {
			DeleteUser(db, user)
			break
		}
	}

	return -1
}

func DeleteUser(db *sql.DB, user models.User) int {
	//input menu
	var phone string
	fmt.Print("\n")
	fmt.Println("Masukan nomor telepon untuk konfirmasi:")
	fmt.Scanln(&phone)

	if GetIdByPhone(db, user, phone) != user.ID {
		fmt.Println("\nGagal menghapus akun!")
		fmt.Println("Nomor yang Anda masukan tidak sesuai")
		fmt.Print("\n")
		return -1
	}

	query := "UPDATE users SET deleted_at = now() WHERE id = ?;"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("error prepare update", errPrepare.Error())
	}

	result, errUpdate := statement.Exec(user.ID)
	if errUpdate != nil {
		log.Fatal("error exec update", errUpdate.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Print("\n")
			fmt.Printf("Akun Anda berhasil dihapus!\n")
			fmt.Print("\n")
		} else {
			fmt.Print("\n")
			fmt.Println("Gagal menghapus akun!")
			fmt.Print("\n")
		}
	}
	return 9
}
