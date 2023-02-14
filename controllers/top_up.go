package controllers

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"project/models"
)

func TopUp(db *sql.DB, user models.User) int {
	//input jumlah saldo
	fmt.Println("input jumlah saldo top up :")
	var saldo float64
	_, err := fmt.Scanln(&saldo)
	if err != nil {
		fmt.Println("gagal top up")
		return -1
	}
	//input berita top up
	fmt.Println("input berita top up :")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	info := in.Text()

	tx, err := db.Begin()
	if err != nil {
		fmt.Println("gagal transaksi top up")
		return -1
	}

	//tambahkan saldo ke balance
	_, err = tx.Exec("UPDATE balances SET balance = balance + ? WHERE user_ID = ?", saldo, user.ID)
	if err != nil {
		fmt.Println("gagal update saldo")
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("tx err: %v, rb err : %v", err, rbErr)
			return -1
		}
		return -1
	}

	//tambahkan history top up
	_, err = tx.Exec("INSERT INTO top_up_histories(date,amount,user_id,info) VALUES (now(),?,?,?)", saldo, user.ID, info)
	if err != nil {
		fmt.Println("gagal menambah history top up")
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("tx err: %v, rb err : %v", err, rbErr)
			return -1
		}
		return -1
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("gagal commit top up")
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("tx err: %v, rb err : %v", err, rbErr)
			return -1
		}
		return -1
	}

	fmt.Println("top up berhasil sejumlah ", saldo)

	fmt.Println("pilih menu :\n1.Menu Utama\n2.Exit")
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
