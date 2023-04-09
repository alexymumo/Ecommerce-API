package db

import (
	"database/sql"
)

func SetUp() {

	db, err := sql.Open("mysql", "@tcp(127.0.0.1:3306)/testdb")

	defer db.Close()
	if err != nil {
		panic("failed to connect to the database")
	}

}
