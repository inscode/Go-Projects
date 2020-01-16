package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB  *sql.DB
	err error
)

func init() {
	DB, err = sql.Open("mysql", "root:11019@tcp(localhost:3306)/bookstore")
	if err != nil {
		panic(err.Error())
	}
}
