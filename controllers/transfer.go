package controllers

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"project/models"
)

func MenuTransfer(db *sql.DB, user models.User) int {
	opsi := 1
	for opsi != -1 {
		fmt.Println("\n1. Transfer\n9. Kembali Ke Menu Utama\n0. Exit")
		fmt.Print("\nPilih menu: ")
		fmt.Scanln(&opsi)
		switch opsi {
		case 1:
			Transfer(db, user)
		case 9:
			return -1
		case 0:
			return 9
		default:
			fmt.Println("Input yang Anda masukan tidak tersedia")
		}
	}
	return -1
}

func Transfer(db *sql.DB, user models.User) int {
	//input nomor telepon
	var phone string
	fmt.Println("Masukan nomor telepon tujuan:")
	fmt.Scanln(&phone)

	//input nominal transfer
	fmt.Println("Masukan nominal transfer:")
	var nominal float64
	_, errScanNominal := fmt.Scanln(&nominal)

	//input berita transfer
	fmt.Println("Masukan informasi:")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	info := in.Text()

	//validasi nomor telpon
	//cek apakah nomor telepon ada atau tidak di database
	if GetIdByPhone(db, user, phone) == 0 {
		fmt.Println("\nTransfer tidak dapat diproses!")
		fmt.Println("Nomor telepon yang Anda masukan salah")
		return 1
	}

	//masukan id ke variable
	receiverID := GetIdByPhone(db, user, phone)

	//tidak bisa transfer ke nomor sendiri
	if receiverID == user.ID {
		fmt.Println("\nTransfer tidak dapat diproses!")
		fmt.Println("Anda tidak dapat melakukan transfer ke nomor Anda sendiri")
		return 1
	}

	//validasi nominal tidak boleh huruf
	if errScanNominal != nil {
		fmt.Println("\nTransfer tidak dapat diproses!")
		fmt.Println("Nominal yang Anda masukan salah")
		return 1
	}

	//validasi nominal tidak boleh negatif
	if nominal < 1 {
		fmt.Println("\nTransfer tidak dapat diproses!")
		fmt.Println("Nominal yang Anda masukan salah")
		return 1
	}

	//validasi saldo pengirim
	if nominal > CheckBalance(db, user.ID) {
		fmt.Println("\nTransfer tidak dapat diproses!")
		fmt.Println("Saldo Anda tidak mencukupi")
		return 1
	}

	//operasi saldo pengirim dan penerima
	senderBalance := CheckBalance(db, user.ID) - nominal
	receiverBalance := CheckBalance(db, receiverID) + nominal

	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Gagal transaksi transfer")
		return 1
	}

	//kurangi saldo pengirim
	_, err = tx.Exec("UPDATE balances SET updated_at = now(), balance = ? WHERE user_id = ?", senderBalance, user.ID)
	if err != nil {
		fmt.Println("gagal update saldo pengirim")
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("tx err: %v, rb err : %v", err, rbErr)
			return 1
		}
		return 1
	}

	//tambahkan saldo penerima
	_, err = tx.Exec("UPDATE balances SET updated_at = now(), balance = ? WHERE user_id = ?", receiverBalance, receiverID)
	if err != nil {
		fmt.Println("gagal update saldo penerima")
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("tx err: %v, rb err : %v", err, rbErr)
			return 1
		}
		return 1
	}

	//tambahkan history transfer
	_, err = tx.Exec("INSERT INTO transfer_histories(date,amount,user_id_sender, user_id_receiver,info) VALUES (now(),?,?,?,?)", nominal, user.ID, receiverID, info)
	if err != nil {
		fmt.Println("gagal menambah history transfer")
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("tx err: %v, rb err : %v", err, rbErr)
			return 1
		}
		return 1
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("gagal commit transfer")
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("tx err: %v, rb err : %v", err, rbErr)
			return 1
		}
		return 1
	}

	//mengambil nama user dari id
	senderName := GetNameByID(db, user, user.ID)
	receiverName := GetNameByID(db, user, receiverID)

	fmt.Print("\n")
	fmt.Println("============================")
	fmt.Printf("Berhasil transfer Rp%.2f\n", nominal)
	fmt.Printf("\nNama Pengirim: %s\n", senderName)
	fmt.Printf("Nama Penerima: %s\n", receiverName)
	fmt.Printf("\nSisa saldo Anda sekarang Rp%.2f\n", senderBalance)
	fmt.Println("============================")

	return 1
}

func CheckBalance(db *sql.DB, userId int) float64 {
	query := "SELECT balance FROM balances WHERE user_id = ?;"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("error prepare select", errPrepare.Error())
	}

	var balance models.Balance
	balance.ID = userId
	errScan := statement.QueryRow(balance.ID).Scan(&balance.Balance)
	if errScan != nil {
		log.Fatal("error scan select", errScan.Error())
	}

	return balance.Balance
}

func GetIdByPhone(db *sql.DB, selectUser models.User, phone string) int {
	query := "SELECT id FROM users WHERE deleted_at IS NULL AND phone = ?;"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("error prepare select", errPrepare.Error())
	}

	var user models.User
	selectUser.Phone = phone
	errScan := statement.QueryRow(selectUser.Phone).Scan(&user.ID)
	if errScan != nil {
		return 0
	}
	return user.ID
}

func GetNameByID(db *sql.DB, selectUser models.User, id int) string {
	query := "SELECT name FROM users WHERE id = ?;"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("error prepare select", errPrepare.Error())
	}

	var user models.User
	selectUser.ID = id
	errScan := statement.QueryRow(selectUser.ID).Scan(&user.Name)
	if errScan != nil {
		log.Fatal("error scan select", errScan.Error())
	}

	return user.Name
}
