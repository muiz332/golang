package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func CreateConnection() *sql.DB {

	var db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/test")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(2 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Second)

	return db
}
