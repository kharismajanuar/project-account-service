package controllers

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"project/helper"
	"project/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func MenuRegister(db *sql.DB, user models.User) int {
	opsi := 1
	for opsi != -1 {
		fmt.Print("\n")
		fmt.Println("1. Register Akun Baru\n9. Kembali Ke Menu Utama")
		fmt.Print("\nPilih menu: ")
		fmt.Scanln(&opsi)
		switch opsi {
		case 1:
			RegisterUser(db, user)
		case 9:
			return -1
		default:
			fmt.Println("Input yang Anda masukan tidak tersedia")
		}
	}
	return -1
}

func RegisterUser(db *sql.DB, newUser models.User) int {
	scanner := bufio.NewScanner(os.Stdin)

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
	validName, msgName := helper.ValidasiNama(newUser.Name)
	if !validName {
		fmt.Println(msgName)
		return -1
	}

	//validasi nomor telpon
	validPhone, msgPhone := helper.ValidasiTelepon(newUser.Phone, db)
	if !validPhone {
		fmt.Println(msgPhone)
		return -1
	}

	//validasi password
	validPass, msgPass := helper.ValidasiPassword(newUser.Password)
	if !validPass {
		fmt.Println(msgPass)
		return -1
	}

	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		log.Fatal("error hashed password", err.Error())
		return -1
	}

	//validasi jenis kelamin
	if newUser.Sex != "Pria" && newUser.Sex != "Wanita" {
		fmt.Println("\nGagal menambahkan akun!")
		fmt.Println("Jenis kelamin hanya boleh diisi oleh Pria atau Wanita")
		return -1
	}

	//validasi tanggal lahir
	validDob, msgDob := helper.ValidasiTanggalLahir(date)
	if !validDob {
		fmt.Println(msgDob)
		return -1
	}

	queryInsert := "INSERT INTO users (name, phone, password, sex, date_of_birth) VALUES (?, ?, ?, ?, ?);"
	statementInsert, errPrepare := db.Prepare(queryInsert)
	if errPrepare != nil {
		log.Fatal("error prepare insert", errPrepare.Error())
		return -1
	}

	result, errInsert := statementInsert.Exec(newUser.Name, newUser.Phone, hashedPassword, newUser.Sex, newUser.DateOfBirth)
	if errInsert != nil {
		log.Fatal("error exec insert", errInsert.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Print("\n")
			fmt.Printf("Akun dengan nama '%s' telah berhasil ditambahkan!\n", newUser.Name)
			fmt.Print("\n")
		} else {
			fmt.Print("\n")
			fmt.Println("Gagal menambahkan akun baru!")
			fmt.Print("\n")
			return -1
		}
	}

	InsertBalances(db, 0)

	return -1
}

func InsertBalances(db *sql.DB, balance float64) {
	queryInsert := "INSERT INTO balances (balance) VALUES (?);"
	statementInsert, errPrepare := db.Prepare(queryInsert)
	if errPrepare != nil {
		log.Fatal("error prepare insert", errPrepare.Error())
	}

	result, errInsert := statementInsert.Exec(balance)
	if errInsert != nil {
		log.Fatal("error exec insert", errInsert.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Printf("Saldo Anda sekarang Rp%v\n", balance)
		} else {
			fmt.Println("Gagal menambahkan saldo")
		}
	}
}
