package utils

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDBSession() error {
	var err error

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=%s",
		Config.DBUser,
		Config.DBPassword,
		Config.DBName,
		Config.DBPort,
		Config.DBSSLMode)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return err
}
