package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"log"
)

var DB *sql.DB

func Connect(host, port, user, pass, dbname string) error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname)
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	log.Println("✓ Connected to PostgreSQL")
	return nil
}

func CreateTable() error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	)`
	_, err := DB.Exec(query)
	return err
}
func CreateContoh() error {

	query := `CREATE TABLE IF NOT EXISTS contoh (
			id SERIAL PRIMARY KEY,
			pekerjaan VARCHAR(100) NOT NULL,
			umur INTEGER NOT NULL,
			alamat VARCHAR(999) NOT NULL,
			created_at TIMESTAMP DEFAULT NOW()
	)`

	_, err :=DB.Exec(query)

	return err
}
// func CreateStatus() error {
// 	query := 
// }

func Migrate() error {
	if err:= CreateTable(); err != nil {
		return err
	}
	if err := CreateContoh(); err != nil {
		return err
	}
	log.Println("All table succes create")
	return nil
}

