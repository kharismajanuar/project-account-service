package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"project/models"
)

func MenuTransferHistory(db *sql.DB, user models.User) int {

	opsi := 1
	for opsi != 9 {
		fmt.Print("\n")
		GetAllTransferHistories(db, user)
		fmt.Print("9. Kembali Ke Menu Utama\n")
		fmt.Print("\nPilih menu: ")
		fmt.Scanln(&opsi)
	}

	return -1
}

func GetAllTransferHistories(db *sql.DB, user models.User) {
	rows, errSelect := db.Query("SELECT trh.id, usrv.name as receiver_name, usrv.phone as receiver_phone, trh.date, trh.amount, trh.info FROM transfer_histories trh INNER JOIN users usr ON trh.user_id_sender = usr.id INNER JOIN users usrv ON trh.user_id_receiver = usrv.id WHERE usr.id = ? ORDER BY trh.id;", user.ID)
	if errSelect != nil {
		log.Fatal("error query select", errSelect.Error())
	}

	var allTransfer []models.TransferHistories
	for rows.Next() {
		var datarow models.TransferHistories
		errScan := rows.Scan(&datarow.ID, &datarow.ReceiverName, &datarow.ReceiverPhone, &datarow.Date, &datarow.Amount, &datarow.Info)
		if errScan != nil {
			log.Fatal("error scan select", errScan.Error())
		}
		allTransfer = append(allTransfer, datarow)
	}

	//fmt.Println(allTransfer)
	for _, transfers := range allTransfer {
		fmt.Printf("ID Transfer   : %d\n", transfers.ID)
		fmt.Printf("Nama Penerima : %s\n", transfers.ReceiverName)
		fmt.Printf("Nomor Penerima: %s\n", transfers.ReceiverPhone)
		fmt.Printf("Nominal       : Rp%.2f\n", transfers.Amount)
		fmt.Printf("Tanggal       : %s\n", transfers.Date.Format("02 January 2006"))
		fmt.Printf("Pukul         : %s\n", transfers.Date.Format("15:04:05"))
		fmt.Printf("Info          : %s\n", transfers.Info)
		fmt.Print("\n")
	}

	if len(allTransfer) == 0 {
		fmt.Println("Anda belum pernah melakukan transaksi dalam menu transfer")
		fmt.Print("\n")
	}
}
