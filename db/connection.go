package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() error {
	var err error
	DB, err = sql.Open("mysql", "root:1nf0m3D!4@tcp(127.0.0.1:3306)/go")
	if err != nil {
		return err
	}
	return DB.Ping()
}
