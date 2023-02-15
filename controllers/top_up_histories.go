package controllers

import (
	"database/sql"
	"fmt"
	"project/models"
)

func TopUpHistories(db *sql.DB, ID int) int {
	//pilih jangka waktu dalam hari
	var day int
	fmt.Println("input jangka waktu (dalam hari) :")
	fmt.Scanln(&day)

	// //select all data from top_up_histories
	rows, err := db.Query("SELECT date, amount, info FROM top_up_histories WHERE user_id = ? AND datediff(now(),date) <= ?", ID, day)
	if err != nil {
		fmt.Println("gagal menampilkan riwayat top up")
		return -1
	}

	topup := []models.TopUpHistories{}

	for rows.Next() {
		var tmp models.TopUpHistories
		err = rows.Scan(&tmp.Date, &tmp.Amount, &tmp.Info)
		if err != nil {
			fmt.Println("gagal menampilkan riwayat top up")
			return -1
		}
		topup = append(topup, tmp)
	}

	//tampilkan data
	for i, v := range topup {
		fmt.Printf("No :\t%d\n", i+1)
		fmt.Printf("Tanggal :\t%s\n", v.Date.String())
		fmt.Printf("Jumlah :\t%.2f\n", v.Amount)
		fmt.Printf("Info :\t%s\n", v.Info)
	}

	if len(topup) == 0 {
		fmt.Println("Tidak ada data riwayat top up")
	}

	fmt.Println("pilih menu\n1.menu utama\n2.exit")
	var opsi int
	fmt.Scanln(&opsi)
	if opsi == 1 {
		return -1
	} else {
		return 9
	}
}
