package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func connectDB() (*sql.DB, error) {

	db, err := sql.Open("mysql", "string")

	return db, err
}
