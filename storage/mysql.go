package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// New return mysql storage instance
func New(dataSourceName string) (*mysql, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	fmt.Println("success")
	return &mysql{db}, nil
}

// Add new item to database
func (m *mysql) Add(url string) error {
	return nil
}

type mysql struct{ db *sql.DB }
