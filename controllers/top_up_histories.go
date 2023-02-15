package controllers

import (
	"database/sql"
	"fmt"
	"project/models"
)

func TopUpHistories(db *sql.DB, ID int) int {
	//pilih jangka waktu dalam hari
	var day int
	fmt.Print("\n")
	fmt.Printf("Input jangka waktu (dalam hari) : ")
	fmt.Scanln(&day)

	//validasi jangka waktu
	if day <= 0 {
		fmt.Println("Jangka waktu tidak valid")
		return -1
	}

	// //select all data from top_up_histories
	rows, err := db.Query("SELECT date, amount, info FROM top_up_histories WHERE user_id = ? AND datediff(now(),date) <= ?", ID, day)
	if err != nil {
		fmt.Println("Gagal menampilkan riwayat top up")
		return -1
	}

	topup := []models.TopUpHistories{}

	for rows.Next() {
		var tmp models.TopUpHistories
		err = rows.Scan(&tmp.Date, &tmp.Amount, &tmp.Info)
		if err != nil {
			fmt.Println("Gagal menampilkan riwayat top up")
			return -1
		}
		topup = append(topup, tmp)
	}

	//tampilkan data
	if len(topup) != 0 {
		fmt.Print("\n")
		fmt.Println("No\tTanggal\t\t\t\tJumlah\t\tInfo")
	}

	for i, v := range topup {
		fmt.Print("\n")
		fmt.Printf("%d\t", i+1)
		fmt.Printf("%s\t", v.Date.Format("15:04:05 January 2, 2006"))
		fmt.Printf("%.2f\t", v.Amount)
		fmt.Printf("%s\t", v.Info)
		fmt.Print("\n")
	}

	if len(topup) == 0 {
		fmt.Println("Tidak ada data riwayat top up")
	}

	fmt.Print("\n1.Menu utama\n2.Exit\n\nPilih menu : ")
	var opsi int
	fmt.Scanln(&opsi)
	if opsi == 1 {
		return -1
	} else {
		return 9
	}
}
