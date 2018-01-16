// Package mysql provides implementation for Storage interface
package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/matrosov-nikita/url-shortener/storage"
)

// New creates new instance of MySQL storage.
func New(addr string) (storage.Storage, error) {
	db, err := sql.Open("mysql", addr)
	if err != nil {
		return nil, err
	}

	if err = createTable(db); err != nil {
		return nil, err
	}

	return &sqlStorage{db}, nil
}

func createTable(db *sql.DB) error {
	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS urls 
		(id int NOT NULL AUTO_INCREMENT, 
		short_url varchar(20), 
		origin_url varchar(400) character set 'utf8',
		PRIMARY KEY (id));`)

	if err != nil {
		return err
	}

	if _, err = stmt.Exec(); err != nil {
		return err
	}

	return nil
}

func (s *sqlStorage) Count() (int64, error) {
	var count int64
	if err := s.db.QueryRow("SELECT COUNT(*) FROM urls").Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (s *sqlStorage) AddURL(short, origin string) error {
	stmt, err := s.db.Prepare("INSERT INTO urls (short_url, origin_url) VALUES (?, ?)")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(short, origin); err != nil {
		return err
	}

	return nil
}

func (s *sqlStorage) GetURL(short string) (string, error) {
	stmt, err := s.db.Prepare("SELECT origin_url FROM urls WHERE binary short_url=?")
	if err != nil {
		return "", err
	}

	var origin string
	if err := stmt.QueryRow(short).Scan(&origin); err != nil {
		return "", err
	}

	return origin, nil
}

type sqlStorage struct{ db *sql.DB }
