package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open() (*gorm.DB, error) {
	return gorm.Open(
		postgres.Open(os.Getenv("DB_CONNECTION_STRING")),
		&gorm.Config{})
}
