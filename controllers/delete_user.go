package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"project/models"
)

func DeleteUser(db *sql.DB, updateUser models.User) int {
	//input menu
	fmt.Print("\n")
	fmt.Println("Hapus Akun")
	fmt.Println("Masukan nomor telepon:")
	fmt.Scanln(&updateUser.Phone)

	query := "UPDATE users SET deleted_at = now() WHERE phone = ?;"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("error prepare update", errPrepare.Error())
	}

	result, errUpdate := statement.Exec(updateUser.Phone)
	if errUpdate != nil {
		log.Fatal("error exec update", errUpdate.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Print("\n")
			fmt.Printf("Akun dengan nomor telpon %s berhasil dihapus!\n", updateUser.Phone)
			fmt.Print("\n")
		} else {
			fmt.Print("\n")
			fmt.Println("Gagal menghapus akun!")
			fmt.Print("\n")
		}
	}
	return 9
}
