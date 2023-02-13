package config

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func createTable(db *sql.DB, q ...string) error {
	for _, v := range q {
		_, err := db.Exec(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func DBConn() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DB_CONN"))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = createTable(db, TABLE_USERS, TABLE_BALANCES, TABLE_TOP_UP_HISTORIES, TABLE_TRANSFER_HISTORIES)
	if err != nil {
		return nil, err
	}

	return db, nil
}
