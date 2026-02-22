package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewPostgresDB() (*sql.DB, error) {
	connStr := "host=localhost port=5433 user=postgres password=postgres dbname=kaztour sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Точная ошибка:", err.Error())
		return nil, err
	}

	return db, nil
}
