package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open(
		"postgres",
		"host=localhost port=5432 user=nerd password=123 dbname=testDB sslmode=disable",
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}
func CheckUser(db *sql.DB, user User) error {
	return nil
}
func InsertUser(db *sql.DB, user User) error {
	return nil
}
func GetUserPass(db *sql.DB, user User) (string, error) {
	return "", nil
}