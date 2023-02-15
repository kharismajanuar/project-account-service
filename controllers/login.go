package controllers

import (
	"database/sql"
	"fmt"
	"project/helper"
	"project/models"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Login(db *sql.DB) (models.User, int, bool) {
	//input phone, password
	fmt.Print("\n")
	fmt.Print("Input nomor telepon : ")
	var phone string
	_, err := fmt.Scanln(&phone)
	if err != nil {
		fmt.Println("kesalahan input nomor telepon")
		return models.User{}, -1, false
	}

	//validasi nomor telepon
	//validasi telepon
	//minimal 10 karakter
	if len(phone) < 10 || len(phone) > 12 {
		fmt.Println("telepon minimal 10 karakter maksimal 12")
		return models.User{}, -1, false
	}

	//hanya angka
	if !regexp.MustCompile(`^[0-9]*$`).MatchString(phone) {
		fmt.Println("telepon hanya terdiri dari angka")
		return models.User{}, -1, false
	}

	fmt.Print("Input password : ")
	var password string
	_, err = fmt.Scanln(&password)
	if err != nil {
		fmt.Println("kesalahan input nomor password")
		return models.User{}, -1, false
	}

	//validasi password
	valid, msg := helper.ValidasiPassword(password)
	if !valid {
		fmt.Println(msg)
		return models.User{}, -1, false
	}

	//ambil data user dari database dengan phone = phone
	var user models.User = models.User{Name: "acep"}

	err = db.QueryRow("SELECT ID, Phone, Name, Password, Date_Of_Birth, Sex From users WHERE phone = ? AND deleted_at IS NULL", phone).
		Scan(&user.ID, &user.Phone, &user.Name, &user.Password, &user.DateOfBirth, &user.Sex)
	if err != nil {
		fmt.Println("akun tidak ditemukan")
		return models.User{}, -1, false
	}

	//bandingkan password input dengan password dari database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println("password salah")
		return models.User{}, -1, false
	}

	//login berhasil
	fmt.Println(`
	╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╭╮
	╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╭╯╰╮
	╭━━┳━━┳━━┳━━┳╮╭┳━╋╮╭╯╭━━┳━━┳━┳╮╭┳┳━━┳━━╮╭━━┳━━┳━━╮
	┃╭╮┃╭━┫╭━┫╭╮┃┃┃┃╭╮┫┃╱┃━━┫┃━┫╭┫╰╯┣┫╭━┫┃━┫┃╭╮┃╭╮┃╭╮┃
	┃╭╮┃╰━┫╰━┫╰╯┃╰╯┃┃┃┃╰╮┣━━┃┃━┫┃╰╮╭┫┃╰━┫┃━┫┃╭╮┃╰╯┃╰╯┃
	╰╯╰┻━━┻━━┻━━┻━━┻╯╰┻━╯╰━━┻━━┻╯╱╰╯╰┻━━┻━━╯╰╯╰┫╭━┫╭━╯
	╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱┃┃╱┃┃
	╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╰╯╱╰╯`)
	c := time.NewTicker(25 * time.Millisecond)
	go func() {
		var counter int
		for {
			select {
			case _ = <-c.C:
				counter++
				fmt.Printf("=")
				if counter > 50 {
					fmt.Println()
					return
				}
			}
		}
	}()
	time.Sleep(2 * time.Second)
	c.Stop()
	return user, 0, true
}
