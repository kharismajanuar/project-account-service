package controllers

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"project/models"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(db *sql.DB, newUser models.User) int {
	scanner := bufio.NewScanner(os.Stdin)
	newUser = models.User{}

	//input menu
	fmt.Print("\n")
	fmt.Println("Register Akun Baru")
	fmt.Println("Nama Lengkap:")
	scanner.Scan()
	name := scanner.Text()
	newUser.Name = name
	fmt.Println("No Telepon:")
	fmt.Scanln(&newUser.Phone)
	fmt.Println("Password:")
	fmt.Scanln(&newUser.Password)
	fmt.Println("Jenis Kelamin (Pria/Wanita):")
	fmt.Scanln(&newUser.Sex)
	fmt.Println("Tanggal Lahir (d/m/y):")
	var layoutFormat, value string
	var date time.Time
	layoutFormat = "2/1/2006"
	fmt.Scanln(&value)
	date, _ = time.Parse(layoutFormat, value)
	newUser.DateOfBirth = date

	//validasi nama
	//maksimal 50 karakter
	if len(newUser.Name) > 50 {
		fmt.Println("Karakter nama maksimal 50 karakter")
		return -1
	}

	//hanya huruf dan spasi
	if !regexp.MustCompile(`^[a-zA-Z ]*$`).MatchString(newUser.Name) {
		fmt.Println("Nama hanya boleh diisi oleh huruf alfabet atau spasi")
		return -1
	}

	//validasi nomor telepon
	//minimal 10 karakter
	if len(newUser.Phone) < 10 || len(newUser.Phone) > 12 {
		fmt.Println("Nomor telepon minimal 10 karakter dan maksimal 12")
		return -1
	}

	//hanya boleh memasukan angka
	if !regexp.MustCompile(`^[0-9]*$`).MatchString(newUser.Phone) {
		fmt.Println("Nomor telepon hanya boleh terdiri dari angka")
		return -1
	}

	//validasi password
	//minimal 8 karakter
	if len(newUser.Password) <= 8 {
		fmt.Println("Password minimal 8 karakter")
		return -1
	}

	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		fmt.Println("err hashed password")
		return -1
	}

	query := "INSERT INTO users (name, phone, password, sex, date_of_birth) VALUES (?, ?, ?, ?, ?);"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("error prepare insert", errPrepare.Error())
	}

	result, errInsert := statement.Exec(newUser.Name, newUser.Phone, hashedPassword, newUser.Sex, newUser.DateOfBirth)
	if errInsert != nil {
		log.Fatal("error exec insert", errInsert.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Print("\n")
			fmt.Printf("Akun dengan nama '%s' telah berhasil ditambahkan!\n", newUser.Name)
		} else {
			fmt.Println("Gagal menambahkan akun baru!")
		}
	}

	return -1
}

func MenuRegister(db *sql.DB, user models.User) {

	var opsi int = 1

	for opsi != 9 {
		fmt.Print("\n")
		fmt.Println("1. Register Akun Baru\n9. Kembali Ke Menu Utama")
		fmt.Print("\nPilih menu: ")
		fmt.Scanln(&opsi)
		switch opsi {
		case 1:
			opsi = RegisterUser(db, user)
		default:
			fmt.Println("Input tidak sesuai")
		}
	}
}
