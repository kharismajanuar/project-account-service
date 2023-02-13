package controllers

import (
	"database/sql"
	"fmt"
	"project/models"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func Login(db *sql.DB) (models.User, int, bool) {
	//input phone, password
	fmt.Println("Input nomor telepon : ")
	var phone string
	_, err := fmt.Scanln(&phone)
	if err != nil {
		fmt.Println("kesalahan input nomor telepon")
		return models.User{}, -1, false
	}

	//validasi nomor telepon
	if regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(phone) {
		fmt.Println("nomor telepon mengandung huruf")
		return models.User{}, -1, false
	}
	if len(phone) < 10 {
		fmt.Println("nomor telepon kurang dari 10")
		return models.User{}, -1, false
	}

	fmt.Println("Input password : ")
	var password string
	_, err = fmt.Scanln(&password)
	if err != nil {
		fmt.Println("kesalahan input nomor password")
		return models.User{}, -1, false
	}

	//validasi password
	if len(password) < 5 {
		fmt.Println("password tidak boleh kurang dari 5")
		return models.User{}, -1, false
	}

	//ambil data user dari database dengan phone = phone
	var user models.User = models.User{Name: "acep"}

	err = db.QueryRow("SELECT ID, Phone, Name, Password, Date_Of_Birth, Sex From users WHERE phone = ?", phone).
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
	fmt.Println("\n***Login sukses***")
	return user, 0, true
}
