package db

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func GetDB() *sql.DB {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	return

}
