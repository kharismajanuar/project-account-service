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
	var opsi int = -1

	var user models.User
	isLoggedIn := false
	for opsi != 0 {
		fmt.Println("pilih menu\n1.Login\n2.Register")
		_, err = fmt.Scanln(&opsi)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		switch opsi {
		case 1:
			user, opsi, isLoggedIn = controllers.Login(db)
		case 2:
			controllers.MenuRegister(db, user)
		}
	}

	var opsiLogin int = -1
out:
	for isLoggedIn && opsiLogin != 9 {
		fmt.Println("pilih menu\n1.Baca akun\n2.Update Akun\n3.Delete Akun\n4.Top Up\n5.Transfer\n6.Histori Top Up\n7.Histori Transfer\n8.Lihat Profil Lain\n9.Logout")
		fmt.Scanln(&opsiLogin)
		switch opsiLogin {
		case 1:
			//baca akun
			opsiLogin = controllers.ReadAccount(db, user)
		case 2:
			//update akun
			opsi = controllers.UpdateAccount(db, user.Phone)
		case 3:
			//delete akun
			controllers.DeleteUser(db, user)
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
			controllers.MenuGetAllUser(db)
		case 9:
			//logout
			break out
		}
	}

	fmt.Println("Terima Kasih ", user.Name)
}
