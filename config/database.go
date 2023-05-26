package config

import (
	"fmt"

	"github.com/nattrio/go-clean-arch/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB(config *Config) *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Bangkok",
		config.DBHost, config.DBPort, config.DBUsername, config.DBPassword, config.DBName)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	utils.ErrorPanic(err)

	fmt.Println("Connected to database successfully")
	return db
}
