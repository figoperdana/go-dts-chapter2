package config

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "testgolang"
)

func Connect() *sql.DB {
	connStr := "host=" + host + " port=" + strconv.Itoa(port) + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable search_path=dbgolang"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	}

	return db
}

func EnsureTableExists() {
	db := Connect()
	defer db.Close()

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS dbgolang.books (
		id SERIAL PRIMARY KEY,
		title VARCHAR(50) NOT NULL,
		author VARCHAR(25) NOT NULL,
		description VARCHAR(100) NOT NULL
	)`)

	if err != nil {
		log.Fatal("Failed to create table if not exists:", err)
	}
}

