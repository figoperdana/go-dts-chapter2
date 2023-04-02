package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(config *DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable search_path=dbgolang",
		config.Host, config.Port, config.Username, config.Password, config.Name)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

