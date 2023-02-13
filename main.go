package main

import (
	"fmt"
	"project/config"
	"project/controllers"
	"project/models"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//koneksi ke Database
	db, err := config.DBConn()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	//pilih menu 1
	var opsi int
	fmt.Println("pilih menu\n1.Login\n2.Register")
	_, err = fmt.Scanln(&opsi)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var user models.User

	for opsi != 0 {
		switch opsi {
		case 1:
			user, opsi = controllers.Login(db, "", "")
			var opsiLogin int
			fmt.Println("pilih menu\n1.Baca akun\n2.Update Akun\n3.Delete Akun\n4.Top Up\n5.Transfer\n6.Histori Top Up\n7.Histori Transfer\n8.Lihat Profil Lain\n9.Logout")
			fmt.Scanln(&opsiLogin)
			for opsiLogin != 9 {
				switch opsiLogin {
				case 1:
					//baca akun
				case 2:
					//update akun
				case 3:
					//delete akun
				case 4:
					//top up
				case 5:
					//transfer
				case 6:
					//histori top up
				case 7:
					//histori transfer
				case 8:
					//lihat profil lain
				case 9:
					//logout
					break
				}
			}
		case 2:
			opsi = controllers.Register(db)
		}
	}

	fmt.Println("Terima Kasih ", user)

}
