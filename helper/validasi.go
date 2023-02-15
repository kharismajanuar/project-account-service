package helper

import (
	"database/sql"
	"regexp"
	"time"
)

func ValidasiNama(data string) (valid bool, msg string) {
	//validasi nama
	//maksimal 50 karakter
	if len(data) > 50 {
		valid, msg = false, "Nama maksimal 50 karakter"
		return
	}
	//hanya huruf dan spasi
	if !regexp.MustCompile(`^[a-zA-Z ]*$`).MatchString(data) {
		valid, msg = false, "Nama hanya boleh alfabet atau spasi"
		return
	}
	valid, msg = true, ""
	return
}

func ValidasiTanggalLahir(dob time.Time) (valid bool, msg string) {
	if time.Since(dob).Hours()/24/365 < 17 {
		valid, msg = false, "Minimal usia 17 tahun"
		return
	}
	valid, msg = true, ""
	return
}

func ValidasiPassword(data string) (valid bool, msg string) {
	//minimal 8 angka
	if len(data) < 8 {
		valid, msg = false, "Password minimal 5 huruf"
		return
	}
	valid, msg = true, ""
	return
}

func ValidasiTelepon(data string, db *sql.DB) (valid bool, msg string) {
	//validasi telepon
	//minimal 10 karakter
	if len(data) < 10 || len(data) > 12 {
		valid, msg = false, "Telepon minimal 10 karakter maksimal 12"
		return
	}

	//hanya angka
	if !regexp.MustCompile(`^[0-9]*$`).MatchString(data) {
		valid, msg = false, "Telepon hanya terdiri dari angka"
		return
	}

	//validasi duplikat nomor
	query := "SELECT id FROM users WHERE phone = ?;"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		valid, msg = false, "error prepare select"
		return
	}

	errScan := statement.QueryRow(data).Scan()
	if errScan == nil {
		valid, msg = false, "Nomor telepon telah digunakan"
		return
	}

	valid, msg = true, ""
	return
}
