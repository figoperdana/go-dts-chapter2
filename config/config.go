package config

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type AppConfig struct {
	DB *DBConfig
}

type DBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

func ConnectDB() (*gorm.DB, error) {
	dbConfig := &DBConfig{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "admin",
		Name:     "testgolang",
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable search_path=dbgolang",
		dbConfig.Host, dbConfig.Port, dbConfig.Username,
		dbConfig.Password, dbConfig.Name)

	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}

