package controllers

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
)

func TopUp(db *sql.DB, ID int) int {
	//input jumlah saldo
	fmt.Print("\n")
	fmt.Print("Input jumlah saldo top up :")
	var saldo float64
	_, err := fmt.Scanln(&saldo)
	if err != nil {
		fmt.Println("Gagal top up")
		return -1
	}
	//validasi saldo
	if saldo <= 0 {
		fmt.Println("Saldo harus di atas 0")
		return -1
	}
	//input berita top up
	fmt.Print("Input berita top up :")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	info := in.Text()

	//validasi info
	if len(info) > 250 {
		fmt.Println("Gagal top up, maksimal 250 karakter")
		return -1
	}

	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Gagal transaksi top up")
		return -1
	}

	//select saldo
	var currentBalance float64
	err = tx.QueryRow("SELECT balance FROM balances WHERE user_id - ? FOR UPDATE", ID).Scan(&currentBalance)
	if err != nil {
		fmt.Println("Gagal update saldo")
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("tx err: %v, rb err : %v", err, rbErr)
			return -1
		}
		return -1
	}

	//tambahkan saldo ke balance
	_, err = tx.Exec("UPDATE balances SET balance = ?,updated_at = now() WHERE user_ID = ?", currentBalance+saldo, ID)
	if err != nil {
		fmt.Println("Gagal update saldo")
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("tx err: %v, rb err : %v", err, rbErr)
			return -1
		}
		return -1
	}

	//tambahkan history top up
	_, err = tx.Exec("INSERT INTO top_up_histories(date,amount,user_id,info) VALUES (now(),?,?,?)", saldo, ID, info)
	if err != nil {
		fmt.Println("Gagal menambah history top up")
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("tx err: %v, rb err : %v", err, rbErr)
			return -1
		}
		return -1
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("Gagal commit top up")
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("tx err: %v, rb err : %v", err, rbErr)
			return -1
		}
		return -1
	}

	fmt.Printf("Top up berhasil sejumlah %.2f\n", saldo)

	fmt.Print("\n1.Menu Utama\n2.Exit\n\nPilih menu : ")
	var opsi int
	_, err = fmt.Scanln(&opsi)
	if err != nil {
		return 9
	}
	if opsi == 1 {
		return -1
	}
	return 9
}
