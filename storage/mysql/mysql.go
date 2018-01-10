package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func New(address string) (*sqlStorage, error) {
	db, err := sql.Open("mysql", address)
	if err != nil {
		return nil, err
	}

	if err = create(db); err != nil {
		return nil, err
	}

	return &sqlStorage{db}, nil
}

func create(db *sql.DB) error {
	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS urls 
		(id int NOT NULL AUTO_INCREMENT, 
		short_url varchar(200), 
		origin_url varchar(400), 
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
	stmt, err := s.db.Prepare("INSERT urls SET short_url=?,origin_url=?")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(short, origin); err != nil {
		return err
	}

	return nil
}

func (s *sqlStorage) GetURL(short string) (string, error) {
	stmt, err := s.db.Prepare("select origin_url from urls where short_url=?")
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
