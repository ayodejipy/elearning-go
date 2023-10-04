package database

import (
	"fmt"
	"log"

	"github.com/ayodejipy/elearning-go/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func InitDatabaseConnection (config config.DatabaseConfig) (*gorm.DB, error)  {
	connectionString := fmt.Sprintf("host:%s port=%s dbname=%s user=%s password=%s sslmode=disabl", config.Host, config.Port, config.DBName, config.User, config.Password)

	// connect to database using gorm
	database, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
		return nil, err
	}

	// return database and a null error
	return database, nil
}
